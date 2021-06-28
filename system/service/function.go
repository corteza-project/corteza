package service

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	a "github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	function struct {
		actionlog actionlog.Recorder
		store     store.Storer
		ac        functionAccessController
	}

	functionAccessController interface{}
)

func Function() *function {
	return (&function{
		ac:        DefaultAccessControl,
		actionlog: DefaultActionlog,
		store:     DefaultStore,
	})
}

func (svc *function) FindByID(ctx context.Context, ID uint64) (q *types.Function, err error) {
	var (
		rProps = &functionActionProps{}
	)

	err = func() error {
		if ID == 0 {
			return FunctionErrInvalidID()
		}

		if q, err = store.LookupApigwFunctionByID(ctx, svc.store, ID); err != nil {
			return TemplateErrInvalidID().Wrap(err)
		}

		rProps.setFunction(q)

		// if !svc.ac.CanReadMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToRead(qProps)
		// }

		return nil
	}()

	return q, svc.recordAction(ctx, rProps, FunctionActionLookup, err)
}

func (svc *function) Create(ctx context.Context, new *types.Function) (q *types.Function, err error) {
	var (
		qProps = &functionActionProps{function: new}
	)

	err = func() (err error) {
		// if !svc.ac.CanCreateMessagebusQueue(ctx) {
		// 	return QueueErrNotAllowedToCreate(qProps)
		// }

		// Set new values after beforeCreate events are emitted
		new.ID = nextID()
		new.CreatedAt = *now()
		new.CreatedBy = a.GetIdentityFromContext(ctx).Identity()

		if _, err = DefaultRoute.FindByID(ctx, new.Route); err != nil {
			return FunctionErrInvalidRoute(qProps)
		}

		if err = store.CreateApigwFunction(ctx, svc.store, new); err != nil {
			return err
		}

		q = new
		// send the signal to reload all functions
		// 	apigw.Service().Reload(ctx)
		// }

		return nil
	}()

	return q, svc.recordAction(ctx, qProps, FunctionActionCreate, err)
}

func (svc *function) Update(ctx context.Context, upd *types.Function) (q *types.Function, err error) {
	var (
		qProps = &functionActionProps{function: upd}
		qq     *types.Function
		e      error
	)

	err = func() (err error) {
		// if !svc.ac.CanUpdateMessagebusQueue(ctx, upd) {
		// 	return QueueErrNotAllowedToUpdate(qProps)
		// }

		if qq, e = store.LookupApigwFunctionByID(ctx, svc.store, upd.ID); e != nil {
			return FunctionErrNotFound(qProps)
		}

		if _, err = DefaultRoute.FindByID(ctx, upd.Route); err != nil {
			return FunctionErrInvalidRoute(qProps)
		}

		if qq, e = store.LookupApigwFunctionByID(ctx, svc.store, upd.ID); e == nil && qq == nil {
			return FunctionErrNotFound(qProps)
		}

		// Set new values after beforeCreate events are emitted
		upd.UpdatedAt = now()
		upd.CreatedAt = qq.CreatedAt
		upd.UpdatedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwFunction(ctx, svc.store, upd); err != nil {
			return
		}

		q = upd

		// send the signal to reload all function
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return q, svc.recordAction(ctx, qProps, FunctionActionUpdate, err)
}

func (svc *function) DeleteByID(ctx context.Context, ID uint64) (err error) {
	var (
		qProps = &functionActionProps{}
		q      *types.Function
	)

	err = func() (err error) {
		if q, err = store.LookupApigwFunctionByID(ctx, svc.store, ID); err != nil {
			return FunctionErrNotFound(qProps)
		}

		qProps.setFunction(q)

		// if !svc.ac.CanDeleteMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToDelete(qProps)
		// }

		q.DeletedAt = now()
		q.DeletedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwFunction(ctx, svc.store, q); err != nil {
			return
		}

		// send the signal to reload all queues
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return svc.recordAction(ctx, qProps, FunctionActionDelete, err)
}

func (svc *function) UndeleteByID(ctx context.Context, ID uint64) (err error) {
	var (
		qProps = &functionActionProps{}
		q      *types.Function
	)

	err = func() (err error) {
		if q, err = store.LookupApigwFunctionByID(ctx, svc.store, ID); err != nil {
			return FunctionErrNotFound(qProps)
		}

		qProps.setFunction(q)

		// if !svc.ac.CanDeleteMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToDelete(qProps)
		// }

		q.DeletedAt = nil
		q.UpdatedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwFunction(ctx, svc.store, q); err != nil {
			return
		}

		// send the signal to reload all queues
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return svc.recordAction(ctx, qProps, FunctionActionDelete, err)
}

func (svc *function) Search(ctx context.Context, filter types.FunctionFilter) (r types.FunctionSet, f types.FunctionFilter, err error) {
	var (
		aProps = &functionActionProps{search: &filter}
	)

	// For each fetched item, store backend will check if it is valid or not
	// filter.Check = func(res *messagebus.QueueSettings) (bool, error) {
	// 	if !svc.ac.CanReadMessagebusQueue(ctx, res) {
	// 		return false, nil
	// 	}

	// 	return true, nil
	// }

	err = func() error {
		if r, f, err = store.SearchApigwFunctions(ctx, svc.store, filter); err != nil {
			return err
		}

		return nil
	}()

	return r, f, svc.recordAction(ctx, aProps, FunctionActionSearch, err)
}
