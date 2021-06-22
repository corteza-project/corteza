package service

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	a "github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	route struct {
		actionlog actionlog.Recorder
		store     store.Storer
		ac        routeAccessController
	}

	routeAccessController interface {
	}
)

func Route() *route {
	return (&route{
		ac:        DefaultAccessControl,
		actionlog: DefaultActionlog,
		store:     DefaultStore,
	})
}

func (svc *route) FindByID(ctx context.Context, ID uint64) (q *types.Route, err error) {
	var (
		rProps = &routeActionProps{}
	)

	err = func() error {
		if ID == 0 {
			return RouteErrInvalidID()
		}

		if q, err = store.LookupApigwRouteByID(ctx, svc.store, ID); err != nil {
			return RouteErrInvalidID().Wrap(err)
		}

		rProps.setRoute(q)

		// if !svc.ac.CanReadMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToRead(qProps)
		// }

		return nil
	}()

	return q, svc.recordAction(ctx, rProps, RouteActionLookup, err)
}

func (svc *route) Create(ctx context.Context, new *types.Route) (q *types.Route, err error) {
	var (
		qProps = &routeActionProps{new: new}
	)

	err = func() (err error) {
		// if !svc.ac.CanCreateMessagebusQueue(ctx) {
		// 	return QueueErrNotAllowedToCreate(qProps)
		// }

		// Set new values after beforeCreate events are emitted
		new.ID = nextID()
		new.CreatedAt = *now()
		new.CreatedBy = a.GetIdentityFromContext(ctx).Identity()

		// todo
		new.Group = 0

		if err = store.CreateApigwRoute(ctx, svc.store, new); err != nil {
			return err
		}

		q = new

		// send the signal to reload all routes
		// if new.Enabled {
		// 	apigw.Service().Reload(ctx)
		// }

		return nil
	}()

	return q, svc.recordAction(ctx, qProps, RouteActionCreate, err)
}

func (svc *route) Update(ctx context.Context, upd *types.Route) (q *types.Route, err error) {
	var (
		qProps = &routeActionProps{update: upd}
		qq     *types.Route
		e      error
	)

	err = func() (err error) {
		// if !svc.ac.CanUpdateMessagebusQueue(ctx, upd) {
		// 	return QueueErrNotAllowedToUpdate(qProps)
		// }

		if qq, e = store.LookupApigwRouteByID(ctx, svc.store, upd.ID); e != nil {
			return RouteErrNotFound(qProps)
		}

		// temp todo - update itself with the same endpoint
		// if qq, e = store.LookupApigwRouteByEndpoint(ctx, svc.store, upd.Endpoint); e == nil && qq == nil {
		// 	return RouteErrExistsEndpoint(qProps)
		// }

		// Set new values after beforeCreate events are emitted
		upd.UpdatedAt = now()
		upd.CreatedAt = qq.CreatedAt
		upd.UpdatedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwRoute(ctx, svc.store, upd); err != nil {
			return
		}

		q = upd

		// send the signal to reload all route
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return q, svc.recordAction(ctx, qProps, RouteActionUpdate, err)
}

func (svc *route) DeleteByID(ctx context.Context, ID uint64) (err error) {
	var (
		qProps = &routeActionProps{}
		q      *types.Route
	)

	err = func() (err error) {
		if ID == 0 {
			return RouteErrInvalidID()
		}

		if q, err = store.LookupApigwRouteByID(ctx, svc.store, ID); err != nil {
			return
		}

		qProps.setRoute(q)

		// if !svc.ac.CanDeleteMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToDelete(qProps)
		// }

		q.DeletedAt = now()
		q.DeletedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwRoute(ctx, svc.store, q); err != nil {
			return
		}

		// send the signal to reload all queues
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return svc.recordAction(ctx, qProps, RouteActionDelete, err)
}

func (svc *route) UndeleteByID(ctx context.Context, ID uint64) (err error) {
	var (
		qProps = &routeActionProps{}
		q      *types.Route
	)

	err = func() (err error) {
		if ID == 0 {
			return RouteErrInvalidID()
		}

		if q, err = store.LookupApigwRouteByID(ctx, svc.store, ID); err != nil {
			return
		}

		qProps.setRoute(q)

		// if !svc.ac.CanDeleteMessagebusQueue(ctx, q) {
		// 	return QueueErrNotAllowedToDelete(qProps)
		// }

		q.DeletedAt = nil
		q.UpdatedBy = a.GetIdentityFromContext(ctx).Identity()

		if err = store.UpdateApigwRoute(ctx, svc.store, q); err != nil {
			return
		}

		// send the signal to reload all queues
		// apigw.Service().Reload(ctx)

		return nil
	}()

	return svc.recordAction(ctx, qProps, RouteActionDelete, err)
}

func (svc *route) Search(ctx context.Context, filter types.RouteFilter) (r types.RouteSet, f types.RouteFilter, err error) {
	var (
		aProps = &routeActionProps{search: &filter}
	)

	// For each fetched item, store backend will check if it is valid or not
	// filter.Check = func(res *messagebus.QueueSettings) (bool, error) {
	// 	if !svc.ac.CanReadMessagebusQueue(ctx, res) {
	// 		return false, nil
	// 	}

	// 	return true, nil
	// }

	err = func() error {
		if r, f, err = store.SearchApigwRoutes(ctx, svc.store, filter); err != nil {
			return err
		}

		return nil
	}()

	return r, f, svc.recordAction(ctx, aProps, RouteActionSearch, err)
}
