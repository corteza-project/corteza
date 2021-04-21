package service

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/label"
	"github.com/cortezaproject/corteza-server/pkg/rand"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/service/event"
	"github.com/cortezaproject/corteza-server/system/types"
	"strings"
)

type (
	authClient struct {
		ac        authClientAccessController
		eventbus  eventDispatcher
		actionlog actionlog.Recorder
		store     store.Storer
	}

	authClientAccessController interface {
		CanCreateAuthClient(context.Context) bool
		CanReadAuthClient(context.Context, *types.AuthClient) bool
		CanUpdateAuthClient(context.Context, *types.AuthClient) bool
		CanDeleteAuthClient(context.Context, *types.AuthClient) bool
	}
)

var (
	validGrantTypes = map[string]bool{
		"authorization_code": true,
		"client_credentials": true,
	}

	validScopes = map[string]bool{
		"profile":   true,
		"api":       true,
		"discovery": true,
	}
)

// AuthClient is a default authClient service initializer
func AuthClient(s store.Storer, ac authClientAccessController, al actionlog.Recorder, eb eventDispatcher) *authClient {
	return &authClient{store: s, ac: ac, actionlog: al, eventbus: eb}
}

func (svc *authClient) LookupByID(ctx context.Context, ID uint64) (client *types.AuthClient, err error) {
	var (
		aaProps = &authClientActionProps{authClient: &types.AuthClient{ID: ID}}
	)

	client, err = svc.lookupByID(ctx, ID)

	if client != nil {
		client.Secret = ""
	}

	return client, svc.recordAction(ctx, aaProps, AuthClientActionLookup, err)
}

func (svc *authClient) ExposeSecret(ctx context.Context, ID uint64) (secret string, err error) {
	var (
		client  *types.AuthClient
		aaProps = &authClientActionProps{authClient: &types.AuthClient{ID: ID}}
	)

	client, err = svc.lookupByID(ctx, ID)
	if client != nil {
		secret = client.Secret
	}

	return secret, svc.recordAction(ctx, aaProps, AuthClientActionExposeSecret, err)
}

func (svc *authClient) RegenerateSecret(ctx context.Context, ID uint64) (secret string, err error) {
	var (
		client  *types.AuthClient
		aaProps = &authClientActionProps{authClient: &types.AuthClient{ID: ID}}
	)

	client, err = svc.lookupByID(ctx, ID)
	if client != nil {
		secret = string(rand.Bytes(64))
		client.Secret = secret
		err = store.UpdateAuthClient(ctx, svc.store, client)
	}

	return secret, svc.recordAction(ctx, aaProps, AuthClientActionRegenerateSecret, err)
}

func (svc *authClient) lookupByID(ctx context.Context, ID uint64) (client *types.AuthClient, err error) {
	err = func() error {
		if ID == 0 {
			return AuthClientErrInvalidID()
		}

		if client, err = store.LookupAuthClientByID(ctx, svc.store, ID); err != nil {
			return AuthClientErrInvalidID().Wrap(err)
		}

		if !svc.ac.CanReadAuthClient(ctx, client) {
			return AuthClientErrNotAllowedToRead()
		}

		return nil
	}()

	return client, err
}

func (svc *authClient) Search(ctx context.Context, af types.AuthClientFilter) (aa types.AuthClientSet, f types.AuthClientFilter, err error) {
	var (
		aaProps = &authClientActionProps{filter: &af}
	)

	// For each fetched item, store backend will check if it is valid or not
	af.Check = func(res *types.AuthClient) (bool, error) {
		if !svc.ac.CanReadAuthClient(ctx, res) {
			return false, nil
		}

		return true, nil
	}

	err = func() error {
		if af.Deleted > filter.StateExcluded {
			// If list with deleted authClients is requested
			// user must have access permissions to system (ie: is admin)
			//
			// not the best solution but ATM it allows us to have at least
			// some kind of control over who can see deleted authClients
			//if !svc.ac.CanAccess(ctx) {
			//	return AuthClientErrNotAllowedToListAuthClients()
			//}
		}

		if len(af.Labels) > 0 {
			af.LabeledIDs, err = label.Search(
				ctx,
				svc.store,
				types.AuthClient{}.LabelResourceKind(),
				af.Labels,
			)

			if err != nil {
				return err
			}

			// labels specified but no labeled resources found
			if len(af.LabeledIDs) == 0 {
				return nil
			}
		}

		if aa, f, err = store.SearchAuthClients(ctx, svc.store, af); err != nil {
			return err
		}

		if err = label.Load(ctx, svc.store, toLabeledAuthClients(aa)...); err != nil {
			return err
		}

		_ = aa.Walk(func(a *types.AuthClient) error {
			// make sure we do not leak client's secret without explicit request
			a.Secret = ""
			return nil
		})

		return nil

	}()

	return aa, f, svc.recordAction(ctx, aaProps, AuthClientActionSearch, err)
}

func (svc *authClient) Create(ctx context.Context, new *types.AuthClient) (ac *types.AuthClient, err error) {
	var (
		aaProps = &authClientActionProps{new: new}
	)

	err = func() (err error) {
		if !svc.ac.CanCreateAuthClient(ctx) {
			return AuthClientErrNotAllowedToCreate()
		}

		if err = svc.eventbus.WaitFor(ctx, event.AuthClientBeforeCreate(new, nil)); err != nil {
			return
		}

		// Set new values after beforeCreate events are emitted
		new.ID = nextID()
		new.CreatedAt = *now()
		new.Secret = string(rand.Bytes(64))

		if new.ValidGrant == "" {
			new.ValidGrant = "authorization_code"
		} else if !validGrantTypes[new.ValidGrant] {
			return AuthClientErrUnknownGrantType()
		}

		if !validateScopes(new.Scope) {
			return AuthClientErrUnknownGrantType()
		}

		if new.Security == nil {
			new.Security = &types.AuthClientSecurity{}
		}

		if new.Meta == nil {
			new.Meta = &types.AuthClientMeta{}
		}

		if err = store.CreateAuthClient(ctx, svc.store, new); err != nil {
			return
		}

		if err = label.Create(ctx, svc.store, new); err != nil {
			return
		}

		ac = new

		_ = svc.eventbus.WaitFor(ctx, event.AuthClientAfterCreate(new, nil))
		return nil
	}()

	return ac, svc.recordAction(ctx, aaProps, AuthClientActionCreate, err)
}

func (svc *authClient) Update(ctx context.Context, upd *types.AuthClient) (app *types.AuthClient, err error) {
	var (
		aaProps = &authClientActionProps{update: upd}
	)

	err = func() (err error) {
		if upd.ID == 0 {
			return AuthClientErrInvalidID()
		}

		if app, err = store.LookupAuthClientByID(ctx, svc.store, upd.ID); err != nil {
			return
		}

		aaProps.setAuthClient(app)

		if !svc.ac.CanUpdateAuthClient(ctx, app) {
			return AuthClientErrNotAllowedToUpdate()
		}

		if err = svc.eventbus.WaitFor(ctx, event.AuthClientBeforeUpdate(upd, app)); err != nil {
			return
		}

		if app.ValidGrant != upd.ValidGrant && !validGrantTypes[upd.ValidGrant] {
			return AuthClientErrUnknownGrantType()
		}

		if app.Scope != upd.Scope && !validateScopes(upd.Scope) {
			return AuthClientErrUnknownGrantType()
		}

		// Assign changed values after afterUpdate events are emitted
		app.Handle = upd.Handle
		app.ValidGrant = upd.ValidGrant
		app.RedirectURI = upd.RedirectURI
		app.Scope = upd.Scope
		app.Enabled = upd.Enabled
		app.Trusted = upd.Trusted
		app.ValidFrom = upd.ValidFrom
		app.ExpiresAt = upd.ExpiresAt
		app.UpdatedAt = now()

		if upd.Meta != nil {
			app.Meta = upd.Meta
		}

		if upd.Security != nil {
			app.Security = upd.Security
		}

		if err = store.UpdateAuthClient(ctx, svc.store, app); err != nil {
			return err
		}

		if label.Changed(app.Labels, upd.Labels) {
			if err = label.Update(ctx, svc.store, upd); err != nil {
				return
			}
			app.Labels = upd.Labels
		}

		_ = svc.eventbus.WaitFor(ctx, event.AuthClientAfterUpdate(upd, app))
		return nil
	}()

	return app, svc.recordAction(ctx, aaProps, AuthClientActionUpdate, err)
}

func (svc *authClient) Delete(ctx context.Context, ID uint64) (err error) {
	var (
		aaProps = &authClientActionProps{}
		app     *types.AuthClient
	)

	err = func() (err error) {
		if ID == 0 {
			return AuthClientErrInvalidID()
		}

		if app, err = store.LookupAuthClientByID(ctx, svc.store, ID); err != nil {
			return
		}

		aaProps.setAuthClient(app)

		if !svc.ac.CanDeleteAuthClient(ctx, app) {
			return AuthClientErrNotAllowedToDelete()
		}

		if err = svc.eventbus.WaitFor(ctx, event.AuthClientBeforeDelete(nil, app)); err != nil {
			return
		}

		app.DeletedAt = now()
		if err = store.UpdateAuthClient(ctx, svc.store, app); err != nil {
			return
		}

		_ = svc.eventbus.WaitFor(ctx, event.AuthClientAfterDelete(nil, app))
		return nil
	}()

	return svc.recordAction(ctx, aaProps, AuthClientActionDelete, err)
}

func (svc *authClient) Undelete(ctx context.Context, ID uint64) (err error) {
	var (
		aaProps = &authClientActionProps{}
		app     *types.AuthClient
	)

	err = func() (err error) {
		if ID == 0 {
			return AuthClientErrInvalidID()
		}

		if app, err = store.LookupAuthClientByID(ctx, svc.store, ID); err != nil {
			return
		}

		aaProps.setAuthClient(app)

		if !svc.ac.CanDeleteAuthClient(ctx, app) {
			return AuthClientErrNotAllowedToUndelete()
		}

		// @todo add event
		//       if err = svc.eventbus.WaitFor(ctx, event.AuthClientBeforeUndelete(nil, app)); err != nil {
		//       	return
		//       }

		app.DeletedAt = nil
		if err = store.UpdateAuthClient(ctx, svc.store, app); err != nil {
			return
		}

		// @todo add event
		//       _ = svc.eventbus.WaitFor(ctx, event.AuthClientAfterUndelete(nil, app))
		return nil
	}()

	return svc.recordAction(ctx, aaProps, AuthClientActionUndelete, err)
}

// toLabeledAuthClients converts to []label.LabeledResource
//
// This function is auto-generated.
func toLabeledAuthClients(set []*types.AuthClient) []label.LabeledResource {
	if len(set) == 0 {
		return nil
	}

	ll := make([]label.LabeledResource, len(set))
	for i := range set {
		ll[i] = set[i]
	}

	return ll
}

func validateScopes(scope string) bool {
	for _, s := range strings.Split(scope, " ") {
		if !validScopes[s] {
			return false
		}
	}

	return true
}
