package service

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/service/user_actions.yaml

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/system/types"
	"strings"
	"time"
)

type (
	userActionProps struct {
		user     *types.User
		new      *types.User
		update   *types.User
		existing *types.User
		filter   *types.UserFilter
	}

	userAction struct {
		timestamp time.Time
		resource  string
		action    string
		log       string
		severity  actionlog.Severity

		// prefix for error when action fails
		errorMessage string

		props *userActionProps
	}

	userLogMetaKey   struct{}
	userPropsMetaKey struct{}
)

var (
	// just a placeholder to cover template cases w/o fmt package use
	_ = fmt.Println
)

// *********************************************************************************************************************
// *********************************************************************************************************************
// Props methods
// setUser updates userActionProps's user
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *userActionProps) setUser(user *types.User) *userActionProps {
	p.user = user
	return p
}

// setNew updates userActionProps's new
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *userActionProps) setNew(new *types.User) *userActionProps {
	p.new = new
	return p
}

// setUpdate updates userActionProps's update
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *userActionProps) setUpdate(update *types.User) *userActionProps {
	p.update = update
	return p
}

// setExisting updates userActionProps's existing
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *userActionProps) setExisting(existing *types.User) *userActionProps {
	p.existing = existing
	return p
}

// setFilter updates userActionProps's filter
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *userActionProps) setFilter(filter *types.UserFilter) *userActionProps {
	p.filter = filter
	return p
}

// Serialize converts userActionProps to actionlog.Meta
//
// This function is auto-generated.
//
func (p userActionProps) Serialize() actionlog.Meta {
	var (
		m = make(actionlog.Meta)
	)

	if p.user != nil {
		m.Set("user.handle", p.user.Handle, true)
		m.Set("user.email", p.user.Email, true)
		m.Set("user.name", p.user.Name, true)
		m.Set("user.username", p.user.Username, true)
		m.Set("user.ID", p.user.ID, true)
	}
	if p.new != nil {
		m.Set("new.handle", p.new.Handle, true)
		m.Set("new.email", p.new.Email, true)
		m.Set("new.name", p.new.Name, true)
		m.Set("new.username", p.new.Username, true)
		m.Set("new.ID", p.new.ID, true)
	}
	if p.update != nil {
		m.Set("update.handle", p.update.Handle, true)
		m.Set("update.email", p.update.Email, true)
		m.Set("update.name", p.update.Name, true)
		m.Set("update.username", p.update.Username, true)
		m.Set("update.ID", p.update.ID, true)
	}
	if p.existing != nil {
		m.Set("existing.handle", p.existing.Handle, true)
		m.Set("existing.email", p.existing.Email, true)
		m.Set("existing.name", p.existing.Name, true)
		m.Set("existing.username", p.existing.Username, true)
		m.Set("existing.ID", p.existing.ID, true)
	}
	if p.filter != nil {
		m.Set("filter.query", p.filter.Query, true)
		m.Set("filter.userID", p.filter.UserID, true)
		m.Set("filter.roleID", p.filter.RoleID, true)
		m.Set("filter.handle", p.filter.Handle, true)
		m.Set("filter.email", p.filter.Email, true)
		m.Set("filter.username", p.filter.Username, true)
		m.Set("filter.deleted", p.filter.Deleted, true)
		m.Set("filter.suspended", p.filter.Suspended, true)
		m.Set("filter.sort", p.filter.Sort, true)
	}

	return m
}

// tr translates string and replaces meta value placeholder with values
//
// This function is auto-generated.
//
func (p userActionProps) Format(in string, err error) string {
	var (
		pairs = []string{"{err}"}
		// first non-empty string
		fns = func(ii ...interface{}) string {
			for _, i := range ii {
				if s := fmt.Sprintf("%v", i); len(s) > 0 {
					return s
				}
			}

			return ""
		}
	)

	if err != nil {
		pairs = append(pairs, err.Error())
	} else {
		pairs = append(pairs, "nil")
	}

	if p.user != nil {
		// replacement for "{user}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{user}",
			fns(
				p.user.Handle,
				p.user.Email,
				p.user.Name,
				p.user.Username,
				p.user.ID,
			),
		)
		pairs = append(pairs, "{user.handle}", fns(p.user.Handle))
		pairs = append(pairs, "{user.email}", fns(p.user.Email))
		pairs = append(pairs, "{user.name}", fns(p.user.Name))
		pairs = append(pairs, "{user.username}", fns(p.user.Username))
		pairs = append(pairs, "{user.ID}", fns(p.user.ID))
	}

	if p.new != nil {
		// replacement for "{new}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{new}",
			fns(
				p.new.Handle,
				p.new.Email,
				p.new.Name,
				p.new.Username,
				p.new.ID,
			),
		)
		pairs = append(pairs, "{new.handle}", fns(p.new.Handle))
		pairs = append(pairs, "{new.email}", fns(p.new.Email))
		pairs = append(pairs, "{new.name}", fns(p.new.Name))
		pairs = append(pairs, "{new.username}", fns(p.new.Username))
		pairs = append(pairs, "{new.ID}", fns(p.new.ID))
	}

	if p.update != nil {
		// replacement for "{update}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{update}",
			fns(
				p.update.Handle,
				p.update.Email,
				p.update.Name,
				p.update.Username,
				p.update.ID,
			),
		)
		pairs = append(pairs, "{update.handle}", fns(p.update.Handle))
		pairs = append(pairs, "{update.email}", fns(p.update.Email))
		pairs = append(pairs, "{update.name}", fns(p.update.Name))
		pairs = append(pairs, "{update.username}", fns(p.update.Username))
		pairs = append(pairs, "{update.ID}", fns(p.update.ID))
	}

	if p.existing != nil {
		// replacement for "{existing}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{existing}",
			fns(
				p.existing.Handle,
				p.existing.Email,
				p.existing.Name,
				p.existing.Username,
				p.existing.ID,
			),
		)
		pairs = append(pairs, "{existing.handle}", fns(p.existing.Handle))
		pairs = append(pairs, "{existing.email}", fns(p.existing.Email))
		pairs = append(pairs, "{existing.name}", fns(p.existing.Name))
		pairs = append(pairs, "{existing.username}", fns(p.existing.Username))
		pairs = append(pairs, "{existing.ID}", fns(p.existing.ID))
	}

	if p.filter != nil {
		// replacement for "{filter}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{filter}",
			fns(
				p.filter.Query,
				p.filter.UserID,
				p.filter.RoleID,
				p.filter.Handle,
				p.filter.Email,
				p.filter.Username,
				p.filter.Deleted,
				p.filter.Suspended,
				p.filter.Sort,
			),
		)
		pairs = append(pairs, "{filter.query}", fns(p.filter.Query))
		pairs = append(pairs, "{filter.userID}", fns(p.filter.UserID))
		pairs = append(pairs, "{filter.roleID}", fns(p.filter.RoleID))
		pairs = append(pairs, "{filter.handle}", fns(p.filter.Handle))
		pairs = append(pairs, "{filter.email}", fns(p.filter.Email))
		pairs = append(pairs, "{filter.username}", fns(p.filter.Username))
		pairs = append(pairs, "{filter.deleted}", fns(p.filter.Deleted))
		pairs = append(pairs, "{filter.suspended}", fns(p.filter.Suspended))
		pairs = append(pairs, "{filter.sort}", fns(p.filter.Sort))
	}
	return strings.NewReplacer(pairs...).Replace(in)
}

// *********************************************************************************************************************
// *********************************************************************************************************************
// Action methods

// String returns loggable description as string
//
// This function is auto-generated.
//
func (a *userAction) String() string {
	var props = &userActionProps{}

	if a.props != nil {
		props = a.props
	}

	return props.Format(a.log, nil)
}

func (e *userAction) ToAction() *actionlog.Action {
	return &actionlog.Action{
		Resource:    e.resource,
		Action:      e.action,
		Severity:    e.severity,
		Description: e.String(),
		Meta:        e.props.Serialize(),
	}
}

// *********************************************************************************************************************
// *********************************************************************************************************************
// Action constructors

// UserActionSearch returns "system:user.search" action
//
// This function is auto-generated.
//
func UserActionSearch(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "search",
		log:       "searched for matching users",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionLookup returns "system:user.lookup" action
//
// This function is auto-generated.
//
func UserActionLookup(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "lookup",
		log:       "looked-up for a {user}",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionCreate returns "system:user.create" action
//
// This function is auto-generated.
//
func UserActionCreate(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "create",
		log:       "created {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionUpdate returns "system:user.update" action
//
// This function is auto-generated.
//
func UserActionUpdate(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "update",
		log:       "updated {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionDelete returns "system:user.delete" action
//
// This function is auto-generated.
//
func UserActionDelete(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "delete",
		log:       "deleted {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionUndelete returns "system:user.undelete" action
//
// This function is auto-generated.
//
func UserActionUndelete(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "undelete",
		log:       "undeleted {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionSuspend returns "system:user.suspend" action
//
// This function is auto-generated.
//
func UserActionSuspend(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "suspend",
		log:       "suspended {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionUnsuspend returns "system:user.unsuspend" action
//
// This function is auto-generated.
//
func UserActionUnsuspend(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "unsuspend",
		log:       "unsuspended {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// UserActionSetPassword returns "system:user.setPassword" action
//
// This function is auto-generated.
//
func UserActionSetPassword(props ...*userActionProps) *userAction {
	a := &userAction{
		timestamp: time.Now(),
		resource:  "system:user",
		action:    "setPassword",
		log:       "password changed for {user}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// *********************************************************************************************************************
// *********************************************************************************************************************
// Error constructors

// UserErrGeneric returns "system:user.generic" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrGeneric(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("failed to complete request due to internal error", nil),

		errors.Meta("type", "generic"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "{err}"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotFound returns "system:user.notFound" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotFound(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("user not found", nil),

		errors.Meta("type", "notFound"),
		errors.Meta("resource", "system:user"),

		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrInvalidID returns "system:user.invalidID" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrInvalidID(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid ID", nil),

		errors.Meta("type", "invalidID"),
		errors.Meta("resource", "system:user"),

		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrInvalidHandle returns "system:user.invalidHandle" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrInvalidHandle(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid handle", nil),

		errors.Meta("type", "invalidHandle"),
		errors.Meta("resource", "system:user"),

		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrInvalidEmail returns "system:user.invalidEmail" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrInvalidEmail(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid email", nil),

		errors.Meta("type", "invalidEmail"),
		errors.Meta("resource", "system:user"),

		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToRead returns "system:user.notAllowedToRead" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToRead(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to read this user", nil),

		errors.Meta("type", "notAllowedToRead"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to read {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToSearch returns "system:user.notAllowedToSearch" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToSearch(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to search for users", nil),

		errors.Meta("type", "notAllowedToSearch"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to search for users; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToListUsers returns "system:user.notAllowedToListUsers" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToListUsers(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to list users", nil),

		errors.Meta("type", "notAllowedToListUsers"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to list user; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToCreate returns "system:user.notAllowedToCreate" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToCreate(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to create users", nil),

		errors.Meta("type", "notAllowedToCreate"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to create users; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToUpdate returns "system:user.notAllowedToUpdate" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToUpdate(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to update this user", nil),

		errors.Meta("type", "notAllowedToUpdate"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to update {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToDelete returns "system:user.notAllowedToDelete" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToDelete(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to delete this user", nil),

		errors.Meta("type", "notAllowedToDelete"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to delete {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToUndelete returns "system:user.notAllowedToUndelete" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToUndelete(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to undelete this user", nil),

		errors.Meta("type", "notAllowedToUndelete"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to undelete {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToSuspend returns "system:user.notAllowedToSuspend" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToSuspend(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to suspend this user", nil),

		errors.Meta("type", "notAllowedToSuspend"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to suspend {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrNotAllowedToUnsuspend returns "system:user.notAllowedToUnsuspend" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrNotAllowedToUnsuspend(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to unsuspend this user", nil),

		errors.Meta("type", "notAllowedToUnsuspend"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "failed to unsuspend {user.handle}; insufficient permissions"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrHandleNotUnique returns "system:user.handleNotUnique" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrHandleNotUnique(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("handle not unique", nil),

		errors.Meta("type", "handleNotUnique"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "used duplicate handle ({user.handle}) for user"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrEmailNotUnique returns "system:user.emailNotUnique" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrEmailNotUnique(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("email not unique", nil),

		errors.Meta("type", "emailNotUnique"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "used duplicate email ({user.email}) for user"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrUsernameNotUnique returns "system:user.usernameNotUnique" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrUsernameNotUnique(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("username not unique", nil),

		errors.Meta("type", "usernameNotUnique"),
		errors.Meta("resource", "system:user"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(userLogMetaKey{}, "used duplicate username ({user.username}) for user"),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// UserErrPasswordNotSecure returns "system:user.passwordNotSecure" as *errors.Error
//
//
// This function is auto-generated.
//
func UserErrPasswordNotSecure(mm ...*userActionProps) *errors.Error {
	var p = &userActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("provided password is not secure; use longer password with more non-alphanumeric character", nil),

		errors.Meta("type", "passwordNotSecure"),
		errors.Meta("resource", "system:user"),

		// link to documentation; formatting applies in case we need some special link formatting
		errors.Meta("documentation", p.Format("foo", nil)),

		// details, used in detailed eror reporting
		errors.Meta("details", p.Format("foo bar", nil)),
		errors.Meta(userPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// *********************************************************************************************************************
// *********************************************************************************************************************

// recordAction is a service helper function wraps function that can return error
//
// It will wrap unrecognized/internal errors with generic errors.
//
// This function is auto-generated.
//
func (svc user) recordAction(ctx context.Context, props *userActionProps, actionFn func(...*userActionProps) *userAction, err error) error {
	if svc.actionlog == nil || actionFn == nil {
		// action log disabled or no action fn passed, return error as-is
		return err
	} else if err == nil {
		// action completed w/o error, record it
		svc.actionlog.Record(ctx, actionFn(props).ToAction())
		return nil
	}

	a := actionFn(props).ToAction()

	// Extracting error information and recording it as action
	a.Error = err.Error()

	switch c := err.(type) {
	case *errors.Error:
		m := c.Meta()

		a.Error = err.Error()
		a.Severity = actionlog.Severity(m.AsInt("severity"))
		a.Description = props.Format(m.AsString(userLogMetaKey{}), err)

		if p, has := m[userPropsMetaKey{}]; has {
			a.Meta = p.(*userActionProps).Serialize()
		}

		svc.actionlog.Record(ctx, a)
	default:
		svc.actionlog.Record(ctx, a)
	}

	// Original error is passed on
	return err
}
