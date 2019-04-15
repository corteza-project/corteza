package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/crusttech/crust/internal/rand"
	"github.com/crusttech/crust/system/internal/repository"
	"github.com/crusttech/crust/system/types"
)

type (
	auth struct {
		db  db
		ctx context.Context

		credentials   repository.CredentialsRepository
		users         repository.UserRepository
		settings      authSettings
		notifications AuthNotificationService

		providerValidator func(string) error
		now               func() *time.Time
	}

	AuthService interface {
		With(ctx context.Context) AuthService

		External(profile goth.User) (*types.User, error)

		InternalSignUp(input *types.User, password string) (*types.User, error)
		InternalLogin(email string, password string) (*types.User, error)
		SetPassword(userID uint64, newPassword string) error
		ChangePassword(userID uint64, oldPassword, newPassword string) error

		ValidateEmailConfirmationToken(token string) (user *types.User, err error)
		ExchangePasswordResetToken(token string) (user *types.User, exchangedToken string, err error)
		ValidatePasswordResetToken(token string) (user *types.User, err error)
		SendEmailAddressConfirmationToken(email string) (err error)
		SendPasswordResetToken(email string) (err error)
	}
)

const (
	credentialsTypePassword                    = "password"
	credentialsTypeEmailAuthToken              = "email-authentication-token"
	credentialsTypeResetPasswordToken          = "password-reset-token"
	credentialsTypeResetPasswordTokenExchanged = "password-reset-token-exchanged"

	credentialsTokenLength = 32
)

var (
	reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func defaultProviderValidator(provider string) error {
	_, err := goth.GetProvider(provider)
	return err
}

func Auth(ctx context.Context) AuthService {
	return (&auth{}).With(ctx)
}

func (svc *auth) With(ctx context.Context) AuthService {
	db := repository.DB(ctx)
	return &auth{
		db:  db,
		ctx: ctx,

		credentials: repository.Credentials(ctx, db),
		users:       repository.User(ctx, db),

		settings:      DefaultAuthSettings,
		notifications: DefaultAuthNotification,

		providerValidator: defaultProviderValidator,
		now: func() *time.Time {
			var now = time.Now()
			return &now
		},
	}
}

// External func performs login/signup procedures
//
// We fully trust external auth sources (see system/internal/auth/external) to provide a valid & validates
// profile (goth.User) that we use for user discovery and/or creation
//
// Flow
// 1.   check for existing credentials using profile provider & provider's user ID
// 1.1. find existing local -or- "shadow" user
// 1.2. if user exists and is valid, update credentials (last-used-at) and complete the procedure
//
// 2.   check for existing users using email from the profile
// 2.1. validate existing user -or-
// 2.2. create user on-the-fly if it does not exist
// 2.3. create credentials for that social login
//
func (svc *auth) External(profile goth.User) (u *types.User, err error) {
	if !svc.settings.externalEnabled {
		return nil, errors.New("external authentication disabled")
	}

	if err = svc.providerValidator(profile.Provider); err != nil {
		return nil, err
	}

	if profile.Email == "" {
		return nil, errors.New("can not use profile data without an email")
	}

	return u, svc.db.Transaction(func() error {
		var c *types.Credentials
		if cc, err := svc.credentials.FindByCredentials(profile.Provider, profile.UserID); err == nil {
			// Credentials found, load user
			for _, c := range cc {
				if !c.Valid() {
					continue
				}

				if u, err = svc.users.FindByID(c.OwnerID); err != nil {
					if err == repository.ErrUserNotFound {
						// Orphaned credentials (no owner)
						// try to auto-fix this by removing credentials and recreating user
						if err := svc.credentials.DeleteByID(c.ID); err != nil {
							return errors.Wrap(err, "could not cleanup orphaned credentials")
						} else {
							goto findByEmail
						}
					}
					return nil
				} else if u.Valid() {
					// Valid user, Bingo!
					c.LastUsedAt = svc.now()
					if c, err = svc.credentials.Update(c); err != nil {
						return err
					}

					log.Printf(
						"updating credential entry (%v, %v) for exisintg user (%v, %v)",
						c.ID,
						profile.Provider,
						u.ID,
						u.Email,
					)

					return nil
				} else {
					// Scenario: linked to an invalid user
					u = nil
					continue
				}
			}

			// If we could not find anything useful,
			// we can search for user via email
		} else {
			// A serious error occurred, bail out...
			return err
		}

	findByEmail:
		// Find user via his email
		if u, err = svc.users.FindByEmail(profile.Email); err == repository.ErrUserNotFound {
			// @todo check if it is ok to auto-create a user here

			// In case we do not have this email, create a new user
			u = &types.User{
				Email:    profile.Email,
				Name:     profile.Name,
				Username: profile.NickName,
				Handle:   profile.NickName,
			}

			if u, err = svc.users.Create(u); err != nil {
				return errors.Wrap(err, "could not create user after successful external authentication")
			}

			log.Printf("created new user after successful social auth (%v, %v)", u.ID, u.Email)
		} else if err != nil {
			return err
		} else if !u.Valid() {
			return errors.Errorf(
				"user not valid",
				u.ID,
			)
		} else {
			log.Printf(
				"autheticated user (%v, %v) via %s, existing user",
				u.ID,
				u.Email,
				profile.Provider,
			)
		}

		c = &types.Credentials{
			Kind:        profile.Provider,
			OwnerID:     u.ID,
			Credentials: profile.UserID,
			LastUsedAt:  svc.now(),
		}

		if c, err = svc.credentials.Create(c); err != nil {
			return err
		}

		log.Printf(
			"creating new credential entry (%v, %v) for exisintg user (%v, %v)",
			c.ID,
			profile.Provider,
			u.ID,
			u.Email,
		)

		// Owner loaded, carry on.
		return nil
	})
}

// InternalSignUp protocol
//
// Forgiving but strict: valid existing users get notified
//
// We're accepting the whole user object here and copy all we need to the new user
func (svc auth) InternalSignUp(input *types.User, password string) (u *types.User, err error) {
	if !svc.settings.internalEnabled {
		return nil, errors.New("internal authentication disabled")
	}

	if !svc.settings.internalSignUpEnabled {
		return nil, errors.New("internal signup disabled")
	}

	if input == nil {
		return nil, errors.New("invalid signup input")
	}

	if err = svc.validateInternalSignUp(input.Email); err != nil {
		return
	}

	err = svc.db.Transaction(func() error {
		existing, err := svc.users.FindByEmail(input.Email)

		if err == nil && existing.Valid() {
			if len(password) > 0 {
				cc, err := svc.credentials.FindByKind(existing.ID, credentialsTypePassword)
				if err != nil {
					return errors.Wrap(err, "could not find credentials")
				}

				err = svc.checkPassword(password, cc)
				if err != nil {
					return err
				}

				if !existing.EmailConfirmed {
					err = svc.sendEmailAddressConfirmationToken(existing)
					if err != nil {
						return err
					}
				}

				u = existing
				return nil
			}

			return errors.Wrap(err, "user with this email already exists")

			// if !svc.settings.internalSignUpSendEmailOnExisting {
			// 	return errors.Wrap(err, "user with this email already exists")
			// }

			// User already exists, but we're nice and we'll send this user an
			// email that will help him to login
			// if !u.Valid() {
			// 	return errors.New("could not validate the user")
			// }
			//
			// return nil
		}

		if err != repository.ErrUserNotFound {
			return errors.New("could not check existing emails")
		}

		// Whitelisted user data to copy
		u, err = svc.users.Create(&types.User{
			Email:    input.Email,
			Name:     input.Name,
			Username: input.Username,
			Handle:   input.Handle,

			// Do we need confirmed email?
			EmailConfirmed: !svc.settings.internalSignUpEmailConfirmationRequired,
		})

		if err != nil {
			return errors.Wrap(err, "could not create user")
		}

		if len(password) > 0 {
			var hash []byte
			hash, err = svc.hashPassword(password)
			if err != nil {
				return err
			}

			err = svc.changePassword(u.ID, hash)
			if err != nil {
				return err
			}
		}

		if !u.EmailConfirmed {
			err = svc.sendEmailAddressConfirmationToken(u)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if !u.EmailConfirmed {
		return nil, errors.New("user email pending confirmation")
	}

	return u, nil
}

func (svc auth) validateInternalSignUp(email string) (err error) {
	if !reEmail.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// InternalLogin verifies username/password combination in the internal credentials table
//
// Expects plain text password as an input
func (svc *auth) InternalLogin(email string, password string) (u *types.User, err error) {

	if !svc.settings.internalEnabled {
		return nil, errors.New("internal authentication disabled")
	}

	if err = svc.validateInternalLogin(email, password); err != nil {
		return
	}

	err = svc.db.Transaction(func() error {
		var (
			cc types.CredentialsSet
		)

		u, err = svc.users.FindByEmail(email)
		if err == repository.ErrUserNotFound {
			return errors.New("invalid username/password combination")
		}

		if err != nil {
			return errors.Wrap(err, "could not find user")
		}

		cc, err := svc.credentials.FindByKind(u.ID, credentialsTypePassword)
		if err != nil {
			return errors.Wrap(err, "could not find credentials")
		}

		err = svc.checkPassword(password, cc)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	if !u.Valid() {
		u = nil
		err = errors.New("user not valid")
		return
	}

	if !u.EmailConfirmed {
		err = svc.sendEmailAddressConfirmationToken(u)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("user email pending confirmation")
	}

	return u, err
}

// validateInternalLogin does basic format & length check
func (svc auth) validateInternalLogin(email string, password string) error {
	if !reEmail.MatchString(email) {
		return errors.Errorf("invalid email format, %s", email)
	}

	if len(password) == 0 {
		return errors.New("empty password")
	}

	return nil
}

func (svc auth) checkPassword(password string, cc types.CredentialsSet) (err error) {
	// We need only valid credentials (skip deleted, expired)
	cc, _ = cc.Filter(func(c *types.Credentials) (b bool, e error) {
		return c.Valid(), nil
	})

	for _, c := range cc {
		if len(c.Credentials) == 0 {
			continue
		}

		err = bcrypt.CompareHashAndPassword([]byte(c.Credentials), []byte(password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// Mismatch, continue with the checking
			continue
		} else if err != nil {
			// Some other error
			return errors.Wrap(err, "could not compare passwords")
		} else {
			// Password matched one of credentials
			return nil
		}
	}

	return errors.New("invalid username/password combination")
}

// SetPassword sets new password for a user
func (svc auth) SetPassword(userID uint64, newPassword string) (err error) {
	if !svc.settings.internalEnabled {
		return errors.New("internal authentication disabled")
	}

	if len(newPassword) == 0 {
		return errors.New("new password missing")
	}

	var hash []byte
	hash, err = svc.hashPassword(newPassword)
	if err != nil {
		return err
	}

	return svc.db.Transaction(func() error {
		if err != svc.changePassword(userID, hash) {
			return err
		}

		log.Printf("password set for user ID %d", userID)
		return nil
	})
}

// ChangePassword validates old password and changes it with new
func (svc auth) ChangePassword(userID uint64, oldPassword, newPassword string) (err error) {
	if !svc.settings.internalEnabled {
		return errors.New("internal authentication disabled")
	}

	if len(oldPassword) == 0 {
		return errors.New("old password missing")
	}

	if len(newPassword) == 0 {
		return errors.New("new password missing")
	}

	var hash []byte
	hash, err = svc.hashPassword(newPassword)
	if err != nil {
		return err
	}

	return svc.db.Transaction(func() error {
		var (
			cc types.CredentialsSet
		)

		cc, err = svc.credentials.FindByKind(userID, credentialsTypePassword)
		if err != nil {
			return errors.Wrap(err, "could not find credentials")
		}

		err = svc.checkPassword(oldPassword, cc)
		if err != nil {
			return errors.Wrap(err, "could not change password")
		}

		if err != svc.changePassword(userID, hash) {
			return err
		}

		log.Printf("password changed for user ID %d", userID)
		return nil
	})
}

func (svc auth) hashPassword(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "could not hash password")
	}

	return
}

// ChangePassword (soft) deletes old password entry and creates a new one
//
// Expects hashed password as an input
func (svc auth) changePassword(userID uint64, hash []byte) (err error) {
	if err = svc.credentials.DeleteByKind(userID, credentialsTypePassword); err != nil {
		return errors.Wrap(err, "could not remove old passwords")
	}

	_, err = svc.credentials.Create(&types.Credentials{
		OwnerID:     userID,
		Kind:        credentialsTypePassword,
		Credentials: string(hash),
	})

	return errors.Wrap(err, "could not create new password")
}

func (svc auth) ValidateEmailConfirmationToken(token string) (user *types.User, err error) {
	if !svc.settings.internalEnabled {
		return nil, errors.New("internal authentication disabled")
	}

	return svc.loadUserFromToken(token, credentialsTypeEmailAuthToken)
}

func (svc auth) ValidatePasswordResetToken(token string) (user *types.User, err error) {
	if !svc.settings.internalEnabled {
		return nil, errors.New("internal authentication disabled")
	}

	if !svc.settings.internalPasswordResetEnabled {
		return nil, errors.New("password reset disabled")
	}

	return svc.loadUserFromToken(token, credentialsTypeResetPasswordTokenExchanged)
}

// ExchangePasswordResetToken exchanges reset password token for a new one and returns it with user info
func (svc auth) ExchangePasswordResetToken(token string) (user *types.User, exchangedToken string, err error) {
	if !svc.settings.internalEnabled {
		err = errors.New("internal authentication disabled")
		return
	}

	if !svc.settings.internalPasswordResetEnabled {
		err = errors.New("password reset disabled")
		return
	}

	user, err = svc.loadUserFromToken(token, credentialsTypeResetPasswordToken)
	if err != nil {
		user = nil
		return
	}

	exchangedToken, err = svc.createUserToken(user, credentialsTypeResetPasswordTokenExchanged)
	if err != nil {
		user = nil
		exchangedToken = ""
		return
	}

	return
}

func (svc auth) SendEmailAddressConfirmationToken(email string) error {
	if !svc.settings.internalEnabled {
		return errors.New("internal authentication disabled")
	}

	u, err := svc.users.FindByEmail(email)
	if err != nil {
		return errors.Wrap(err, "could  not load user")
	}

	return svc.sendEmailAddressConfirmationToken(u)
}

func (svc auth) sendEmailAddressConfirmationToken(u *types.User) (err error) {
	var (
		notificationLang = "en"
		token, url       string
	)

	token, err = svc.createUserToken(u, credentialsTypeEmailAuthToken)
	if err != nil {
		return
	}

	log.Printf("email address validation token generated: %q", token)

	url = svc.settings.frontendUrlEmailConfirmation + token

	err = svc.notifications.EmailConfirmation(notificationLang, u.Email, url)
	if err != nil {
		return errors.Wrap(err, "could not send email authentication notification")
	}

	log.Printf("send email addresss validation notification mail to %s", u.Email)

	return nil
}

func (svc auth) SendPasswordResetToken(email string) error {

	if !svc.settings.internalEnabled {
		return errors.New("internal authentication disabled")
	}

	if !svc.settings.internalPasswordResetEnabled {
		return errors.New("password reset disabled")
	}

	u, err := svc.users.FindByEmail(email)
	if err != nil {
		return errors.Wrap(err, "could  not load user")
	}

	return svc.sendPasswordResetToken(u)
}

func (svc auth) sendPasswordResetToken(u *types.User) (err error) {
	var (
		notificationLang = "en"
		token, url       string
	)

	token, err = svc.createUserToken(u, credentialsTypeResetPasswordToken)
	if err != nil {
		return
	}

	log.Printf("password reset token generated: %q", token)

	url = svc.settings.frontendUrlPasswordReset + token

	err = svc.notifications.PasswordReset(notificationLang, u.Email, url)
	if err != nil {
		return errors.Wrap(err, "could not send password reset notification")
	}

	log.Printf("send password reset notification mail to %s", u.Email)

	return nil
}

func (svc auth) loadUserFromToken(token, kind string) (u *types.User, err error) {
	credentialsID, credentials, err := svc.validateToken(token)
	if err != nil {
		return
	}

	return u, svc.db.Transaction(func() error {
		c, err := svc.credentials.FindByID(credentialsID)
		if err == repository.ErrCredentialsNotFound {
			return errors.New("no such token")
		}

		if err != nil {
			return errors.Wrap(err, "could not load credentials")
		}

		if err = svc.credentials.DeleteByID(c.ID); err != nil {
			return errors.Wrap(err, "could not remove credentials")
		}

		if !c.Valid() {
			return errors.New("expired or invalid token")
		}

		if c.Credentials != credentials {
			return errors.New("invalid token")
		}

		u, err = svc.users.FindByID(c.OwnerID)
		if err != nil {
			return errors.Wrap(err, "could not load user")
		}

		if !u.Valid() {
			u = nil
			return errors.New("user not valid")
		}

		return nil
	})
}

func (svc auth) validateToken(token string) (ID uint64, credentials string, err error) {
	// Token = <32 random chars><credentials-id>
	if len(token) <= credentialsTokenLength {
		err = errors.New("invalid token length")
		return
	}

	ID, err = strconv.ParseUint(token[credentialsTokenLength:], 10, 64)
	if err != nil {
		err = errors.Wrap(err, "invalid token format")
		return
	}

	if ID == 0 {
		err = errors.New("invalid token ID")
		return
	}

	credentials = token[:credentialsTokenLength]
	return
}

func (svc auth) createUserToken(user *types.User, kind string) (token string, err error) {
	c, err := svc.credentials.Create(&types.Credentials{
		OwnerID:     user.ID,
		Kind:        kind,
		Credentials: string(rand.Bytes(credentialsTokenLength)),
	})

	if err != nil {
		return
	}

	token = fmt.Sprintf("%s%d", c.Credentials, c.ID)
	return
}

var _ AuthService = &auth{}
