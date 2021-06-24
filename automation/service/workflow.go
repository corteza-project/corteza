package service

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/options"
	"reflect"
	"sync"

	"github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	intAuth "github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/handle"
	"github.com/cortezaproject/corteza-server/pkg/label"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/cortezaproject/corteza-server/store"
	"go.uber.org/zap"
)

type (
	workflow struct {
		eventbus  workflowEventTriggerHandler
		store     store.Storer
		actionlog actionlog.Recorder
		ac        workflowAccessController
		triggers  *trigger
		session   *session

		opt options.WorkflowOpt

		log *zap.Logger

		// maps resolved workflow graphs to workflow ID (key, uint64)
		wfgs map[uint64]*wfexec.Graph

		// workflow function registry
		reg         *registry
		corredorOpt options.CorredorOpt

		mux    *sync.RWMutex
		parser expr.Parsable
	}

	workflowAccessController interface {
		CanAccess(context.Context) bool

		CanCreateWorkflow(context.Context) bool
		CanReadWorkflow(context.Context, *types.Workflow) bool
		CanUpdateWorkflow(context.Context, *types.Workflow) bool
		CanDeleteWorkflow(context.Context, *types.Workflow) bool
		CanUndeleteWorkflow(context.Context, *types.Workflow) bool

		CanManageWorkflowSessions(context.Context, *types.Workflow) bool

		Grant(ctx context.Context, rr ...*rbac.Rule) error

		workflowExecController
	}

	workflowExecController interface {
		CanExecuteWorkflow(context.Context, *types.Workflow) bool
	}

	workflowEventTriggerHandler interface {
		Register(h eventbus.HandlerFn, ops ...eventbus.HandlerRegOp) uintptr
		Unregister(ptrs ...uintptr)
	}

	workflowUpdateHandler func(ctx context.Context, ns *types.Workflow) (workflowChanges, error)
	workflowChanges       uint8
)

const (
	workflowUnchanged     workflowChanges = 0
	workflowChanged       workflowChanges = 1
	workflowLabelsChanged workflowChanges = 2
	workflowDefChanged    workflowChanges = 4
)

func Workflow(log *zap.Logger, opt options.WorkflowOpt, corredorOpt options.CorredorOpt) *workflow {
	return &workflow{
		log:         log,
		opt:         opt,
		actionlog:   DefaultActionlog,
		store:       DefaultStore,
		ac:          DefaultAccessControl,
		triggers:    DefaultTrigger,
		session:     DefaultSession,
		eventbus:    eventbus.Service(),
		wfgs:        make(map[uint64]*wfexec.Graph),
		mux:         &sync.RWMutex{},
		parser:      expr.NewParser(),
		reg:         Registry(),
		corredorOpt: corredorOpt,
	}
}

func (svc *workflow) Search(ctx context.Context, filter types.WorkflowFilter) (rr types.WorkflowSet, f types.WorkflowFilter, err error) {
	var (
		wap = &workflowActionProps{filter: &filter}
	)

	// For each fetched item, store backend will check if it is valid or not
	filter.Check = func(res *types.Workflow) (bool, error) {
		if !svc.ac.CanReadWorkflow(ctx, res) {
			return false, nil
		}

		return true, nil
	}

	err = func() (err error) {
		if filter.Deleted > 0 {
			// If list with deleted or suspended users is requested
			// user must have access permissions to system (ie: is admin)
			//
			// not the best solution but ATM it allows us to have at least
			// some kind of control over who can see deleted or archived workflows
			if !svc.ac.CanAccess(ctx) {
				return WorkflowErrNotAllowedToSearch()
			}
		}

		if len(filter.Labels) > 0 {
			filter.LabeledIDs, err = label.Search(
				ctx,
				svc.store,
				types.Workflow{}.LabelResourceKind(),
				filter.Labels,
			)

			if err != nil {
				return err
			}

			// labels specified but no labeled resources found
			if len(filter.LabeledIDs) == 0 {
				return nil
			}
		}

		if rr, f, err = store.SearchAutomationWorkflows(ctx, svc.store, filter); err != nil {
			return err
		}

		if err = label.Load(ctx, svc.store, toLabeledWorkflows(rr)...); err != nil {
			return err
		}

		return nil
	}()

	return rr, f, svc.recordAction(ctx, wap, WorkflowActionSearch, err)
}

func (svc *workflow) LookupByID(ctx context.Context, workflowID uint64) (wf *types.Workflow, err error) {
	var (
		wap = &workflowActionProps{workflow: &types.Workflow{ID: workflowID}}
	)

	err = store.Tx(ctx, svc.store, func(ctx context.Context, s store.Storer) error {
		if wf, err = loadWorkflow(ctx, s, workflowID); err != nil {
			return err
		}

		if !svc.ac.CanReadWorkflow(ctx, wf) {
			return WorkflowErrNotAllowedToRead()
		}

		if err = label.Load(ctx, svc.store, wf); err != nil {
			return err
		}

		return nil
	})

	return wf, svc.recordAction(ctx, wap, WorkflowActionLookup, err)
}

// Create adds new workflow resource and saves it into store
// It updates service's cache
func (svc *workflow) Create(ctx context.Context, new *types.Workflow) (wf *types.Workflow, err error) {
	var (
		wap   = &workflowActionProps{new: new}
		cUser = intAuth.GetIdentityFromContext(ctx).Identity()
	)

	err = store.Tx(ctx, svc.store, func(ctx context.Context, s store.Storer) (err error) {
		if !svc.ac.CanCreateWorkflow(ctx) {
			return WorkflowErrNotAllowedToCreate()
		}

		if !handle.IsValid(new.Handle) {
			return WorkflowErrInvalidHandle()
		}

		if err = svc.uniqueCheck(ctx, new); err != nil {
			return err
		}

		wf = &types.Workflow{
			ID:           nextID(),
			Handle:       new.Handle,
			Labels:       new.Labels,
			Meta:         new.Meta,
			Enabled:      new.Enabled,
			Trace:        new.Trace,
			KeepSessions: new.KeepSessions,

			Scope: new.Scope,
			Steps: new.Steps,
			Paths: new.Paths,

			// @todo need to check against access control if current user can modify security descriptor
			RunAs:     new.RunAs,
			OwnedBy:   cUser,
			CreatedAt: *now(),
			CreatedBy: cUser,
		}

		if _, wf.Issues = Convert(svc, wf); len(wf.Issues) == 0 {
			if err = svc.triggers.registerWorkflows(ctx, wf); err != nil {
				return err
			}
		}

		if err = store.CreateAutomationWorkflow(ctx, s, wf); err != nil {
			return
		}

		if err = label.Create(ctx, s, wf); err != nil {
			return
		}

		return
	})

	return wf, svc.recordAction(ctx, wap, WorkflowActionCreate, err)
}

// Update modifies existing workflow resource in the store
func (svc *workflow) Update(ctx context.Context, upd *types.Workflow) (*types.Workflow, error) {
	return svc.updater(ctx, upd.ID, WorkflowActionUpdate, func(ctx context.Context, res *types.Workflow) (workflowChanges, error) {
		if !svc.ac.CanUpdateWorkflow(ctx, res) {
			return workflowUnchanged, WorkflowErrNotAllowedToUpdate()
		}

		handler := svc.handleUpdate(upd)
		return handler(ctx, res)
	})
}

func (svc *workflow) DeleteByID(ctx context.Context, workflowID uint64) error {
	return trim1st(svc.updater(ctx, workflowID, WorkflowActionDelete, func(ctx context.Context, res *types.Workflow) (workflowChanges, error) {
		changes, err := svc.handleDelete(ctx, res)
		if err != nil {
			return workflowUnchanged, err
		}

		if _, res.Issues = Convert(svc, res); len(res.Issues) == 0 {
			if err := svc.triggers.registerWorkflows(ctx, res); err != nil {
				return workflowUnchanged, err
			}
		}

		return changes, err

	}))
}

func (svc *workflow) UndeleteByID(ctx context.Context, workflowID uint64) error {
	return trim1st(svc.updater(ctx, workflowID, WorkflowActionUndelete, func(ctx context.Context, res *types.Workflow) (workflowChanges, error) {
		changes, err := svc.handleUndelete(ctx, res)
		if err != nil {
			return workflowUnchanged, err
		}

		if _, res.Issues = Convert(svc, res); len(res.Issues) == 0 {
			if err := svc.triggers.registerWorkflows(ctx, res); err != nil {
				return workflowUnchanged, err
			}
		}

		return changes, err

	}))
}

func (svc workflow) uniqueCheck(ctx context.Context, res *types.Workflow) (err error) {
	if res.Handle != "" {
		if e, _ := store.LookupAutomationWorkflowByHandle(ctx, svc.store, res.Handle); e != nil && e.ID != res.ID {
			return WorkflowErrHandleNotUnique()
		}
	}

	return nil
}

func (svc *workflow) updater(ctx context.Context, workflowID uint64, action func(...*workflowActionProps) *workflowAction, fn workflowUpdateHandler) (*types.Workflow, error) {
	var (
		changes workflowChanges
		res     *types.Workflow
		aProps  = &workflowActionProps{workflow: &types.Workflow{ID: workflowID}}
		err     error
	)

	err = store.Tx(ctx, svc.store, func(ctx context.Context, s store.Storer) (err error) {
		res, err = loadWorkflow(ctx, s, workflowID)
		if err != nil {
			return
		}

		if err = label.Load(ctx, svc.store, res); err != nil {
			return err
		}

		aProps.setWorkflow(res)
		aProps.setUpdate(res)

		if changes, err = fn(ctx, res); err != nil {
			return err
		}

		if _, res.Issues = Convert(svc, res); len(res.Issues) == 0 {
			if err = svc.triggers.registerWorkflows(ctx, res); err != nil {
				return err
			}
		}

		if changes&workflowChanged > 0 {
			if err = store.UpdateAutomationWorkflow(ctx, svc.store, res); err != nil {
				return err
			}
		}

		if changes&workflowLabelsChanged > 0 {
			if err = label.Update(ctx, s, res); err != nil {
				return
			}
		}

		return err
	})

	return res, svc.recordAction(ctx, aProps, action, err)
}

func (svc workflow) handleUpdate(upd *types.Workflow) workflowUpdateHandler {
	return func(ctx context.Context, res *types.Workflow) (changes workflowChanges, err error) {
		if isStale(upd.UpdatedAt, res.UpdatedAt, res.CreatedAt) {
			return workflowUnchanged, WorkflowErrStaleData()
		}

		if upd.Handle != res.Handle && !handle.IsValid(upd.Handle) {
			return workflowUnchanged, WorkflowErrInvalidHandle()
		}

		if err := svc.uniqueCheck(ctx, upd); err != nil {
			return workflowUnchanged, err
		}

		if !svc.ac.CanUpdateWorkflow(ctx, res) {
			return workflowUnchanged, WorkflowErrNotAllowedToUpdate()
		}

		if res.Handle != upd.Handle {
			changes |= workflowChanged
			res.Handle = upd.Handle
		}

		if res.Enabled != upd.Enabled {
			changes |= workflowChanged | workflowDefChanged
			res.Enabled = upd.Enabled
		}

		if upd.Labels != nil {
			if label.Changed(res.Labels, upd.Labels) {
				changes |= workflowLabelsChanged
				res.Labels = upd.Labels
			}
		}

		if res.Trace != upd.Trace {
			changes |= workflowChanged | workflowDefChanged
			res.Trace = upd.Trace
		}

		if res.KeepSessions != upd.KeepSessions {
			changes |= workflowChanged | workflowDefChanged
			res.KeepSessions = upd.KeepSessions
		}

		if upd.Meta != nil {
			if !reflect.DeepEqual(upd.Meta, res.Meta) {
				changes |= workflowChanged
				res.Meta = upd.Meta
			}
		}

		if upd.Scope != nil {
			if !reflect.DeepEqual(upd.Scope, res.Scope) {
				changes |= workflowChanged | workflowDefChanged
				res.Scope = upd.Scope
			}
		}

		if upd.Steps != nil {
			if !reflect.DeepEqual(upd.Steps, res.Steps) {
				changes |= workflowChanged | workflowDefChanged
				res.Steps = upd.Steps
			}
		}

		if upd.Paths != nil {
			if !reflect.DeepEqual(upd.Paths, res.Paths) {
				changes |= workflowChanged | workflowDefChanged
				res.Paths = upd.Paths
			}
		}

		if res.RunAs != upd.RunAs {
			// @todo need to check against access control if current user can modify security descriptor
			changes |= workflowChanged | workflowDefChanged
			res.RunAs = upd.RunAs
		}

		if res.OwnedBy != upd.OwnedBy {
			// @todo need to check against access control if current user can modify owner
			changes |= workflowChanged
			res.OwnedBy = upd.OwnedBy
		}

		if changes&workflowChanged > 0 {
			res.UpdatedAt = now()
		}

		return
	}
}

func (svc workflow) handleDelete(ctx context.Context, res *types.Workflow) (workflowChanges, error) {
	if !svc.ac.CanDeleteWorkflow(ctx, res) {
		return workflowUnchanged, WorkflowErrNotAllowedToDelete()
	}

	if res.DeletedAt != nil {
		// workflow already deleted
		return workflowUnchanged, nil
	}

	res.DeletedAt = now()
	return workflowChanged, nil
}

func (svc workflow) handleUndelete(ctx context.Context, res *types.Workflow) (workflowChanges, error) {
	if !svc.ac.CanUndeleteWorkflow(ctx, res) {
		return workflowUnchanged, WorkflowErrNotAllowedToUndelete()
	}

	if res.DeletedAt == nil {
		// workflow not deleted
		return workflowUnchanged, nil
	}

	res.DeletedAt = nil
	return workflowChanged, nil
}

func (svc *workflow) Load(ctx context.Context) error {
	wwf, _, err := store.SearchAutomationWorkflows(ctx, svc.store, types.WorkflowFilter{
		Deleted:  filter.StateInclusive,
		Disabled: filter.StateExcluded,
	})

	if err != nil {
		return err
	}

	return svc.triggers.registerWorkflows(ctx, wwf...)
}

func (svc *workflow) Exec(ctx context.Context, workflowID uint64, p types.WorkflowExecParams) (*expr.Vars, types.Stacktrace, error) {
	var (
		runAs   intAuth.Identifiable
		wap     = &workflowActionProps{}
		wf      *types.Workflow
		t       *types.Trigger
		results *expr.Vars
		wait    WaitFn

		stacktrace types.Stacktrace
	)

	err := func() (err error) {
		wf, err = loadWorkflow(ctx, svc.store, workflowID)
		if err != nil {
			return
		}

		wap.setWorkflow(wf)
		if !svc.ac.CanExecuteWorkflow(ctx, wf) {
			return WorkflowErrNotAllowedToExecute()
		}

		// User wants to trace workflow execution
		// This means we'll allow him to specify any (orphaned) step
		// even if it's not linked to onManual trigger
		if p.Trace && !svc.ac.CanManageWorkflowSessions(ctx, wf) {
			return WorkflowErrNotAllowedToExecute()
		}

		if !wf.Enabled && !p.Trace {
			return WorkflowErrDisabled()
		}

		g, convErr := Convert(svc, wf)
		if len(convErr) > 0 {
			return convErr
		}

		// Find the trigger.
		t, err = func() (*types.Trigger, error) {
			var tt types.TriggerSet
			tt, _, err = svc.triggers.Search(ctx, types.TriggerFilter{WorkflowID: []uint64{workflowID}})
			if err != nil {
				return nil, err
			}

			if p.StepID == 0 && len(tt) > 0 {
				return tt[0], nil
			} else {
				for _, tMatch := range tt {
					if tMatch.StepID == p.StepID {
						return tMatch, nil
					}
				}
			}

			return nil, nil
		}()

		// Start with workflow scope
		scope := wf.Scope.Merge()

		callStack := wfexec.GetContextCallStack(ctx)
		if len(callStack) > svc.opt.CallStackSize {
			return WorkflowErrMaximumCallStackSizeExceeded()
		}

		ssp := types.SessionStartParams{
			WorkflowID: wf.ID,
			KeepFor:    wf.KeepSessions,
			Trace:      wf.Trace || p.Trace,
			StepID:     p.StepID,

			CallStack: callStack,
		}

		if !p.Trace {
			if t == nil {
				return WorkflowErrUnknownWorkflowStep()
			} else if !t.Enabled {
				return WorkflowErrDisabled()
			}
		}

		if t != nil {
			wap.setTrigger(t)

			// Add trigger's input to scope
			scope = scope.Merge(t.Input)
			_ = scope.AssignFieldValue("eventType", expr.Must(expr.NewString(t.EventType)))
			_ = scope.AssignFieldValue("resourceType", expr.Must(expr.NewString(t.ResourceType)))

			ssp.StepID = t.StepID
			ssp.EventType = t.EventType
			ssp.ResourceType = t.ResourceType
		} else {
			ssp.EventType = "onTrace"
			ssp.ResourceType = ""
		}

		// Finally, assign input values
		ssp.Input = scope.Merge(p.Input)

		if wf.RunAs > 0 {
			if runAs, err = DefaultUser.FindByID(ctx, wf.RunAs); err != nil {
				return
			}
		}

		if runAs == nil {
			// Default to current user
			runAs = intAuth.GetIdentityFromContext(ctx)
		}

		wait, err = svc.session.Start(g, runAs, ssp)

		if err != nil {
			return err
		}

		if !p.Async {
			if !p.Wait && wf.CheckDeferred() {
				// deferred workflow, return right away and keep the workflow session
				// running without waiting for the execution
				return nil
			}
		}

		// wait for the workflow to complete
		// reuse scope for results
		// this will be decoded back to event properties
		results, _, stacktrace, err = wait(ctx)
		return
	}()

	return results, stacktrace, svc.recordAction(ctx, wap, WorkflowActionExecute, err)
}

func makeWorkflowHandler(ac workflowExecController, s *session, t *types.Trigger, wf *types.Workflow, g *wfexec.Graph, runAs intAuth.Identifiable) eventbus.HandlerFn {
	return func(ctx context.Context, ev eventbus.Event) (err error) {
		if !ac.CanExecuteWorkflow(ctx, wf) {
			return WorkflowErrNotAllowedToExecute()
		}

		var (
			// create session scope from predefined workflow scope and trigger input
			scope   = wf.Scope.Merge(t.Input)
			evScope *expr.Vars
			wait    WaitFn

			// This needs to happen!%&!/("#")
			// The returned closure needs to have it's own instance so it doesn't
			// affect the instance bound to the workflow handler
			runAs = runAs
		)

		if enc, is := ev.(varsEncoder); is {
			if evScope, err = enc.EncodeVars(); err != nil {
				return
			}

			scope = scope.Merge(evScope)
		}

		_ = scope.AssignFieldValue("eventType", expr.Must(expr.NewString(ev.EventType())))
		_ = scope.AssignFieldValue("resourceType", expr.Must(expr.NewString(ev.ResourceType())))

		if runAs == nil {
			// @todo can/should we get alternative identity from Event?
			//       for example:
			//         - use http auth header and get username
			//         - use from/to/replyTo and use that as an identifier
			runAs = intAuth.GetIdentityFromContext(ctx)
		}

		callStack := wfexec.GetContextCallStack(ctx)
		if len(callStack) > s.opt.CallStackSize {
			return WorkflowErrMaximumCallStackSizeExceeded()
		}

		wait, err = s.Start(g, runAs, types.SessionStartParams{
			WorkflowID:   wf.ID,
			KeepFor:      wf.KeepSessions,
			Trace:        wf.Trace,
			Input:        scope,
			StepID:       t.StepID,
			EventType:    t.EventType,
			ResourceType: t.ResourceType,

			CallStack: callStack,
		})

		if err != nil {
			return err
		}

		if wf.CheckDeferred() {
			// deferred workflow, return right away and keep the workflow session
			// running without waiting for the execution
			return nil
		}

		// wait for the workflow to complete
		// reuse scope for results
		// this will be decoded back to event properties
		scope, _, _, err = wait(ctx)
		if err != nil {
			return
		}

		if dec, is := ev.(varsDecoder); is {
			return dec.DecodeVars(scope)
		}

		return nil
	}
}

func loadWorkflow(ctx context.Context, s store.Storer, workflowID uint64) (res *types.Workflow, err error) {
	if workflowID == 0 {
		return nil, WorkflowErrInvalidID()
	}

	if res, err = store.LookupAutomationWorkflowByID(ctx, s, workflowID); errors.IsNotFound(err) {
		return nil, WorkflowErrNotFound()
	}

	return
}

// toLabeledWorkflows converts to []label.LabeledResource
func toLabeledWorkflows(set []*types.Workflow) []label.LabeledResource {
	if len(set) == 0 {
		return nil
	}

	ll := make([]label.LabeledResource, len(set))
	for i := range set {
		ll[i] = set[i]
	}

	return ll
}

func exprTypeSetter(reg *registry, e *types.Expr) func(string) (expr.Type, error) {
	return func(name string) (expr.Type, error) {
		if name == "" {
			name = "Any"
		}

		if typ := reg.Type(name); typ != nil {
			return typ, nil
		} else {
			return nil, errors.NotFound(
				"unknown or unregistered type %q used for expression %q on %q",
				name,
				e.Expr,
				e.Target,
			)
		}
	}
}
