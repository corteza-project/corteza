package service

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/service/route_actions.yaml

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
	routeActionProps struct {
		route  *types.Route
		new    *types.Route
		update *types.Route
		search *types.RouteFilter
	}

	routeAction struct {
		timestamp time.Time
		resource  string
		action    string
		log       string
		severity  actionlog.Severity

		// prefix for error when action fails
		errorMessage string

		props *routeActionProps
	}

	routeLogMetaKey   struct{}
	routePropsMetaKey struct{}
)

var (
	// just a placeholder to cover template cases w/o fmt package use
	_ = fmt.Println
)

// *********************************************************************************************************************
// *********************************************************************************************************************
// Props methods
// setRoute updates routeActionProps's route
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *routeActionProps) setRoute(route *types.Route) *routeActionProps {
	p.route = route
	return p
}

// setNew updates routeActionProps's new
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *routeActionProps) setNew(new *types.Route) *routeActionProps {
	p.new = new
	return p
}

// setUpdate updates routeActionProps's update
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *routeActionProps) setUpdate(update *types.Route) *routeActionProps {
	p.update = update
	return p
}

// setSearch updates routeActionProps's search
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *routeActionProps) setSearch(search *types.RouteFilter) *routeActionProps {
	p.search = search
	return p
}

// Serialize converts routeActionProps to actionlog.Meta
//
// This function is auto-generated.
//
func (p routeActionProps) Serialize() actionlog.Meta {
	var (
		m = make(actionlog.Meta)
	)

	if p.route != nil {
		m.Set("route.endpoint", p.route.Endpoint, true)
		m.Set("route.ID", p.route.ID, true)
	}
	if p.new != nil {
		m.Set("new.endpoint", p.new.Endpoint, true)
		m.Set("new.ID", p.new.ID, true)
	}
	if p.update != nil {
		m.Set("update.endpoint", p.update.Endpoint, true)
		m.Set("update.ID", p.update.ID, true)
	}
	if p.search != nil {
	}

	return m
}

// tr translates string and replaces meta value placeholder with values
//
// This function is auto-generated.
//
func (p routeActionProps) Format(in string, err error) string {
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

	if p.route != nil {
		// replacement for "{route}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{route}",
			fns(
				p.route.Endpoint,
				p.route.ID,
			),
		)
		pairs = append(pairs, "{route.endpoint}", fns(p.route.Endpoint))
		pairs = append(pairs, "{route.ID}", fns(p.route.ID))
	}

	if p.new != nil {
		// replacement for "{new}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{new}",
			fns(
				p.new.Endpoint,
				p.new.ID,
			),
		)
		pairs = append(pairs, "{new.endpoint}", fns(p.new.Endpoint))
		pairs = append(pairs, "{new.ID}", fns(p.new.ID))
	}

	if p.update != nil {
		// replacement for "{update}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{update}",
			fns(
				p.update.Endpoint,
				p.update.ID,
			),
		)
		pairs = append(pairs, "{update.endpoint}", fns(p.update.Endpoint))
		pairs = append(pairs, "{update.ID}", fns(p.update.ID))
	}

	if p.search != nil {
		// replacement for "{search}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{search}",
			fns(),
		)
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
func (a *routeAction) String() string {
	var props = &routeActionProps{}

	if a.props != nil {
		props = a.props
	}

	return props.Format(a.log, nil)
}

func (e *routeAction) ToAction() *actionlog.Action {
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

// RouteActionSearch returns "system:route.search" action
//
// This function is auto-generated.
//
func RouteActionSearch(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "search",
		log:       "searched for route",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// RouteActionLookup returns "system:route.lookup" action
//
// This function is auto-generated.
//
func RouteActionLookup(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "lookup",
		log:       "looked-up for a {route}",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// RouteActionCreate returns "system:route.create" action
//
// This function is auto-generated.
//
func RouteActionCreate(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "create",
		log:       "created {route}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// RouteActionUpdate returns "system:route.update" action
//
// This function is auto-generated.
//
func RouteActionUpdate(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "update",
		log:       "updated {route}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// RouteActionDelete returns "system:route.delete" action
//
// This function is auto-generated.
//
func RouteActionDelete(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "delete",
		log:       "deleted {route}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// RouteActionUndelete returns "system:route.undelete" action
//
// This function is auto-generated.
//
func RouteActionUndelete(props ...*routeActionProps) *routeAction {
	a := &routeAction{
		timestamp: time.Now(),
		resource:  "system:route",
		action:    "undelete",
		log:       "undeleted {route}",
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

// RouteErrGeneric returns "system:route.generic" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrGeneric(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("failed to complete request due to internal error", nil),

		errors.Meta("type", "generic"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "{err}"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotFound returns "system:route.notFound" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotFound(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("route not found", nil),

		errors.Meta("type", "notFound"),
		errors.Meta("resource", "system:route"),

		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrInvalidID returns "system:route.invalidID" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrInvalidID(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid ID", nil),

		errors.Meta("type", "invalidID"),
		errors.Meta("resource", "system:route"),

		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrInvalidEndpoint returns "system:route.invalidEndpoint" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrInvalidEndpoint(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid endpoint", nil),

		errors.Meta("type", "invalidEndpoint"),
		errors.Meta("resource", "system:route"),

		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrExistsEndpoint returns "system:route.existsEndpoint" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrExistsEndpoint(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("route with this endpoint already exists", nil),

		errors.Meta("type", "existsEndpoint"),
		errors.Meta("resource", "system:route"),

		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrAlreadyExists returns "system:route.alreadyExists" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrAlreadyExists(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("route by that endpoint already exists", nil),

		errors.Meta("type", "alreadyExists"),
		errors.Meta("resource", "system:route"),

		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToCreate returns "system:route.notAllowedToCreate" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToCreate(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to create a route", nil),

		errors.Meta("type", "notAllowedToCreate"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to create a route; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToRead returns "system:route.notAllowedToRead" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToRead(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to read this route", nil),

		errors.Meta("type", "notAllowedToRead"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to read {route.endpoint}; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToUpdate returns "system:route.notAllowedToUpdate" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToUpdate(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to update this route", nil),

		errors.Meta("type", "notAllowedToUpdate"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to update {route.endpoint}; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToDelete returns "system:route.notAllowedToDelete" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToDelete(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to delete this route", nil),

		errors.Meta("type", "notAllowedToDelete"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to delete {route.endpoint}; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToUndelete returns "system:route.notAllowedToUndelete" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToUndelete(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to undelete this route", nil),

		errors.Meta("type", "notAllowedToUndelete"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to undelete {route.endpoint}; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// RouteErrNotAllowedToExec returns "system:route.notAllowedToExec" as *errors.Error
//
//
// This function is auto-generated.
//
func RouteErrNotAllowedToExec(mm ...*routeActionProps) *errors.Error {
	var p = &routeActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("not allowed to execute this route", nil),

		errors.Meta("type", "notAllowedToExec"),
		errors.Meta("resource", "system:route"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(routeLogMetaKey{}, "failed to exec {route.endpoint}; insufficient permissions"),
		errors.Meta(routePropsMetaKey{}, p),

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
func (svc route) recordAction(ctx context.Context, props *routeActionProps, actionFn func(...*routeActionProps) *routeAction, err error) error {
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
		a.Description = props.Format(m.AsString(routeLogMetaKey{}), err)

		if p, has := m[routePropsMetaKey{}]; has {
			a.Meta = p.(*routeActionProps).Serialize()
		}

		svc.actionlog.Record(ctx, a)
	default:
		svc.actionlog.Record(ctx, a)
	}

	// Original error is passed on
	return err
}
