package service

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/service/auth_client_actions.yaml

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
	authClientActionProps struct {
		authClient *types.AuthClient
		new        *types.AuthClient
		update     *types.AuthClient
		filter     *types.AuthClientFilter
	}

	authClientAction struct {
		timestamp time.Time
		resource  string
		action    string
		log       string
		severity  actionlog.Severity

		// prefix for error when action fails
		errorMessage string

		props *authClientActionProps
	}

	authClientLogMetaKey   struct{}
	authClientPropsMetaKey struct{}
)

var (
	// just a placeholder to cover template cases w/o fmt package use
	_ = fmt.Println
)

// *********************************************************************************************************************
// *********************************************************************************************************************
// Props methods
// setAuthClient updates authClientActionProps's authClient
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *authClientActionProps) setAuthClient(authClient *types.AuthClient) *authClientActionProps {
	p.authClient = authClient
	return p
}

// setNew updates authClientActionProps's new
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *authClientActionProps) setNew(new *types.AuthClient) *authClientActionProps {
	p.new = new
	return p
}

// setUpdate updates authClientActionProps's update
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *authClientActionProps) setUpdate(update *types.AuthClient) *authClientActionProps {
	p.update = update
	return p
}

// setFilter updates authClientActionProps's filter
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *authClientActionProps) setFilter(filter *types.AuthClientFilter) *authClientActionProps {
	p.filter = filter
	return p
}

// Serialize converts authClientActionProps to actionlog.Meta
//
// This function is auto-generated.
//
func (p authClientActionProps) Serialize() actionlog.Meta {
	var (
		m = make(actionlog.Meta)
	)

	if p.authClient != nil {
		m.Set("authClient.handle", p.authClient.Handle, true)
		m.Set("authClient.ID", p.authClient.ID, true)
	}
	if p.new != nil {
		m.Set("new.handle", p.new.Handle, true)
		m.Set("new.ID", p.new.ID, true)
	}
	if p.update != nil {
		m.Set("update.handle", p.update.Handle, true)
		m.Set("update.ID", p.update.ID, true)
	}
	if p.filter != nil {
		m.Set("filter.handle", p.filter.Handle, true)
		m.Set("filter.deleted", p.filter.Deleted, true)
		m.Set("filter.sort", p.filter.Sort, true)
	}

	return m
}

// tr translates string and replaces meta value placeholder with values
//
// This function is auto-generated.
//
func (p authClientActionProps) Format(in string, err error) string {
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

	if p.authClient != nil {
		// replacement for "{authClient}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{authClient}",
			fns(
				p.authClient.Handle,
				p.authClient.ID,
			),
		)
		pairs = append(pairs, "{authClient.handle}", fns(p.authClient.Handle))
		pairs = append(pairs, "{authClient.ID}", fns(p.authClient.ID))
	}

	if p.new != nil {
		// replacement for "{new}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{new}",
			fns(
				p.new.Handle,
				p.new.ID,
			),
		)
		pairs = append(pairs, "{new.handle}", fns(p.new.Handle))
		pairs = append(pairs, "{new.ID}", fns(p.new.ID))
	}

	if p.update != nil {
		// replacement for "{update}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{update}",
			fns(
				p.update.Handle,
				p.update.ID,
			),
		)
		pairs = append(pairs, "{update.handle}", fns(p.update.Handle))
		pairs = append(pairs, "{update.ID}", fns(p.update.ID))
	}

	if p.filter != nil {
		// replacement for "{filter}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{filter}",
			fns(
				p.filter.Handle,
				p.filter.Deleted,
				p.filter.Sort,
			),
		)
		pairs = append(pairs, "{filter.handle}", fns(p.filter.Handle))
		pairs = append(pairs, "{filter.deleted}", fns(p.filter.Deleted))
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
func (a *authClientAction) String() string {
	var props = &authClientActionProps{}

	if a.props != nil {
		props = a.props
	}

	return props.Format(a.log, nil)
}

func (e *authClientAction) ToAction() *actionlog.Action {
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

// AuthClientActionSearch returns "system:auth-client.search" action
//
// This function is auto-generated.
//
func AuthClientActionSearch(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "search",
		log:       "searched for auth clients",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionLookup returns "system:auth-client.lookup" action
//
// This function is auto-generated.
//
func AuthClientActionLookup(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "lookup",
		log:       "looked-up for a {authClient}",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionCreate returns "system:auth-client.create" action
//
// This function is auto-generated.
//
func AuthClientActionCreate(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "create",
		log:       "created {authClient}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionUpdate returns "system:auth-client.update" action
//
// This function is auto-generated.
//
func AuthClientActionUpdate(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "update",
		log:       "updated {authClient}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionDelete returns "system:auth-client.delete" action
//
// This function is auto-generated.
//
func AuthClientActionDelete(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "delete",
		log:       "deleted {authClient}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionUndelete returns "system:auth-client.undelete" action
//
// This function is auto-generated.
//
func AuthClientActionUndelete(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "undelete",
		log:       "undeleted {authClient}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionExposeSecret returns "system:auth-client.exposeSecret" action
//
// This function is auto-generated.
//
func AuthClientActionExposeSecret(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "exposeSecret",
		log:       "secret exposed",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// AuthClientActionRegenerateSecret returns "system:auth-client.regenerateSecret" action
//
// This function is auto-generated.
//
func AuthClientActionRegenerateSecret(props ...*authClientActionProps) *authClientAction {
	a := &authClientAction{
		timestamp: time.Now(),
		resource:  "system:auth-client",
		action:    "regenerateSecret",
		log:       "secret regenerated",
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

// AuthClientErrGeneric returns "system:auth-client.generic" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrGeneric(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("failed to complete request due to internal error", nil),

		errors.Meta("type", "generic"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "{err}"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotFound returns "system:auth-client.notFound" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotFound(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("auth client not found", nil),

		errors.Meta("type", "notFound"),
		errors.Meta("resource", "system:auth-client"),

		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrInvalidID returns "system:auth-client.invalidID" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrInvalidID(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid ID", nil),

		errors.Meta("type", "invalidID"),
		errors.Meta("resource", "system:auth-client"),

		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrUnknownGrantType returns "system:auth-client.unknownGrantType" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrUnknownGrantType(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("unknown grant type", nil),

		errors.Meta("type", "unknownGrantType"),
		errors.Meta("resource", "system:auth-client"),

		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrUnknownScope returns "system:auth-client.unknownScope" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrUnknownScope(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("unknown scope", nil),

		errors.Meta("type", "unknownScope"),
		errors.Meta("resource", "system:auth-client"),

		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToRead returns "system:auth-client.notAllowedToRead" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToRead(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to read this auth client", nil),

		errors.Meta("type", "notAllowedToRead"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to read {authClient}; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToListAuthClients returns "system:auth-client.notAllowedToListAuthClients" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToListAuthClients(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to list auth clients", nil),

		errors.Meta("type", "notAllowedToListAuthClients"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to list authClient; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToCreate returns "system:auth-client.notAllowedToCreate" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToCreate(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to create auth clients", nil),

		errors.Meta("type", "notAllowedToCreate"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to create authClient; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToUpdate returns "system:auth-client.notAllowedToUpdate" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToUpdate(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to update this auth client", nil),

		errors.Meta("type", "notAllowedToUpdate"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to update {authClient}; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToDelete returns "system:auth-client.notAllowedToDelete" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToDelete(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to delete this auth client", nil),

		errors.Meta("type", "notAllowedToDelete"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to delete {authClient}; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// AuthClientErrNotAllowedToUndelete returns "system:auth-client.notAllowedToUndelete" as *errors.Error
//
//
// This function is auto-generated.
//
func AuthClientErrNotAllowedToUndelete(mm ...*authClientActionProps) *errors.Error {
	var p = &authClientActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to undelete this auth client", nil),

		errors.Meta("type", "notAllowedToUndelete"),
		errors.Meta("resource", "system:auth-client"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(authClientLogMetaKey{}, "failed to undelete {authClient}; insufficient permissions"),
		errors.Meta(authClientPropsMetaKey{}, p),

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
func (svc authClient) recordAction(ctx context.Context, props *authClientActionProps, actionFn func(...*authClientActionProps) *authClientAction, err error) error {
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
		a.Description = props.Format(m.AsString(authClientLogMetaKey{}), err)

		if p, has := m[authClientPropsMetaKey{}]; has {
			a.Meta = p.(*authClientActionProps).Serialize()
		}

		svc.actionlog.Record(ctx, a)
	default:
		svc.actionlog.Record(ctx, a)
	}

	// Original error is passed on
	return err
}
