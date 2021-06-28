package service

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/service/function_actions.yaml

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
	functionActionProps struct {
		function *types.Function
		search   *types.FunctionFilter
	}

	functionAction struct {
		timestamp time.Time
		resource  string
		action    string
		log       string
		severity  actionlog.Severity

		// prefix for error when action fails
		errorMessage string

		props *functionActionProps
	}

	functionLogMetaKey   struct{}
	functionPropsMetaKey struct{}
)

var (
	// just a placeholder to cover template cases w/o fmt package use
	_ = fmt.Println
)

// *********************************************************************************************************************
// *********************************************************************************************************************
// Props methods
// setFunction updates functionActionProps's function
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *functionActionProps) setFunction(function *types.Function) *functionActionProps {
	p.function = function
	return p
}

// setSearch updates functionActionProps's search
//
// Allows method chaining
//
// This function is auto-generated.
//
func (p *functionActionProps) setSearch(search *types.FunctionFilter) *functionActionProps {
	p.search = search
	return p
}

// Serialize converts functionActionProps to actionlog.Meta
//
// This function is auto-generated.
//
func (p functionActionProps) Serialize() actionlog.Meta {
	var (
		m = make(actionlog.Meta)
	)

	if p.function != nil {
		m.Set("function.ID", p.function.ID, true)
		m.Set("function.ref", p.function.Ref, true)
	}
	if p.search != nil {
	}

	return m
}

// tr translates string and replaces meta value placeholder with values
//
// This function is auto-generated.
//
func (p functionActionProps) Format(in string, err error) string {
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

	if p.function != nil {
		// replacement for "{function}" (in order how fields are defined)
		pairs = append(
			pairs,
			"{function}",
			fns(
				p.function.ID,
				p.function.Ref,
			),
		)
		pairs = append(pairs, "{function.ID}", fns(p.function.ID))
		pairs = append(pairs, "{function.ref}", fns(p.function.Ref))
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
func (a *functionAction) String() string {
	var props = &functionActionProps{}

	if a.props != nil {
		props = a.props
	}

	return props.Format(a.log, nil)
}

func (e *functionAction) ToAction() *actionlog.Action {
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

// FunctionActionSearch returns "system:function.search" action
//
// This function is auto-generated.
//
func FunctionActionSearch(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "search",
		log:       "searched for function",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// FunctionActionLookup returns "system:function.lookup" action
//
// This function is auto-generated.
//
func FunctionActionLookup(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "lookup",
		log:       "looked-up for a {function}",
		severity:  actionlog.Info,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// FunctionActionCreate returns "system:function.create" action
//
// This function is auto-generated.
//
func FunctionActionCreate(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "create",
		log:       "created {function}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// FunctionActionUpdate returns "system:function.update" action
//
// This function is auto-generated.
//
func FunctionActionUpdate(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "update",
		log:       "updated {function}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// FunctionActionDelete returns "system:function.delete" action
//
// This function is auto-generated.
//
func FunctionActionDelete(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "delete",
		log:       "deleted {function}",
		severity:  actionlog.Notice,
	}

	if len(props) > 0 {
		a.props = props[0]
	}

	return a
}

// FunctionActionUndelete returns "system:function.undelete" action
//
// This function is auto-generated.
//
func FunctionActionUndelete(props ...*functionActionProps) *functionAction {
	a := &functionAction{
		timestamp: time.Now(),
		resource:  "system:function",
		action:    "undelete",
		log:       "undeleted {function}",
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

// FunctionErrGeneric returns "system:function.generic" as *errors.Error
//
//
// This function is auto-generated.
//
func FunctionErrGeneric(mm ...*functionActionProps) *errors.Error {
	var p = &functionActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("failed to complete request due to internal error", nil),

		errors.Meta("type", "generic"),
		errors.Meta("resource", "system:function"),

		// action log entry; no formatting, it will be applied inside recordAction fn.
		errors.Meta(functionLogMetaKey{}, "{err}"),
		errors.Meta(functionPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// FunctionErrNotFound returns "system:function.notFound" as *errors.Error
//
//
// This function is auto-generated.
//
func FunctionErrNotFound(mm ...*functionActionProps) *errors.Error {
	var p = &functionActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("function not found", nil),

		errors.Meta("type", "notFound"),
		errors.Meta("resource", "system:function"),

		errors.Meta(functionPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// FunctionErrInvalidID returns "system:function.invalidID" as *errors.Error
//
//
// This function is auto-generated.
//
func FunctionErrInvalidID(mm ...*functionActionProps) *errors.Error {
	var p = &functionActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid ID", nil),

		errors.Meta("type", "invalidID"),
		errors.Meta("resource", "system:function"),

		errors.Meta(functionPropsMetaKey{}, p),

		errors.StackSkip(1),
	)

	if len(mm) > 0 {
	}

	return e
}

// FunctionErrInvalidRoute returns "system:function.invalidRoute" as *errors.Error
//
//
// This function is auto-generated.
//
func FunctionErrInvalidRoute(mm ...*functionActionProps) *errors.Error {
	var p = &functionActionProps{}
	if len(mm) > 0 {
		p = mm[0]
	}

	var e = errors.New(
		errors.KindInternal,

		p.Format("invalid route", nil),

		errors.Meta("type", "invalidRoute"),
		errors.Meta("resource", "system:function"),

		errors.Meta(functionPropsMetaKey{}, p),

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
func (svc function) recordAction(ctx context.Context, props *functionActionProps, actionFn func(...*functionActionProps) *functionAction, err error) error {
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
		a.Description = props.Format(m.AsString(functionLogMetaKey{}), err)

		if p, has := m[functionPropsMetaKey{}]; has {
			a.Meta = p.(*functionActionProps).Serialize()
		}

		svc.actionlog.Record(ctx, a)
	default:
		svc.actionlog.Record(ctx, a)
	}

	// Original error is passed on
	return err
}
