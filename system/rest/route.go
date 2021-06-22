package rest

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/api"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/system/rest/request"
	"github.com/cortezaproject/corteza-server/system/service"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Route struct {
		svc routeService
		ac  templateAccessController
	}

	routePayload struct {
		*types.Route
	}

	routeSetPayload struct {
		Filter types.RouteFilter `json:"filter"`
		Set    []*routePayload   `json:"set"`
	}

	routeService interface {
		FindByID(ctx context.Context, ID uint64) (*types.Route, error)
		Create(ctx context.Context, new *types.Route) (*types.Route, error)
		Update(ctx context.Context, upd *types.Route) (*types.Route, error)
		DeleteByID(ctx context.Context, ID uint64) error
		UndeleteByID(ctx context.Context, ID uint64) error
		Search(ctx context.Context, filter types.RouteFilter) (types.RouteSet, types.RouteFilter, error)
	}
)

func (Route) New() *Route {
	return &Route{
		svc: service.DefaultRoute,
		ac:  service.DefaultAccessControl,
	}
}

func (ctrl *Route) List(ctx context.Context, r *request.RouteList) (interface{}, error) {
	var (
		err error
		f   = types.RouteFilter{
			Deleted: filter.State(r.Deleted),
		}
	)

	if f.Paging, err = filter.NewPaging(r.Limit, r.PageCursor); err != nil {
		return nil, err
	}

	if f.Sorting, err = filter.NewSorting(r.Sort); err != nil {
		return nil, err
	}

	set, filter, err := ctrl.svc.Search(ctx, f)

	return ctrl.makeFilterPayload(ctx, set, filter, err)
}

func (ctrl *Route) Create(ctx context.Context, r *request.RouteCreate) (interface{}, error) {
	var (
		err error
		q   = &types.Route{
			Endpoint: r.Endpoint,
			Method:   r.Method,
			Debug:    r.Debug,
			Enabled:  r.Enabled,
		}
	)

	q, err = ctrl.svc.Create(ctx, q)

	return ctrl.makePayload(ctx, q, err)
}

func (ctrl *Route) Read(ctx context.Context, r *request.RouteRead) (interface{}, error) {
	return ctrl.svc.FindByID(ctx, r.RouteID)
}

func (ctrl *Route) Update(ctx context.Context, r *request.RouteUpdate) (interface{}, error) {
	var (
		err error
		q   = &types.Route{
			ID:       r.RouteID,
			Endpoint: r.Endpoint,
			Method:   r.Method,
			Debug:    r.Debug,
			Group:    uint64(r.Group),
			Enabled:  r.Enabled,
		}
	)

	q, err = ctrl.svc.Update(ctx, q)

	return ctrl.makePayload(ctx, q, err)
}

func (ctrl *Route) Delete(ctx context.Context, r *request.RouteDelete) (interface{}, error) {
	return api.OK(), ctrl.svc.DeleteByID(ctx, r.RouteID)
}

func (ctrl *Route) Undelete(ctx context.Context, r *request.RouteUndelete) (interface{}, error) {
	return api.OK(), ctrl.svc.UndeleteByID(ctx, r.RouteID)
}

func (ctrl *Route) makePayload(ctx context.Context, q *types.Route, err error) (*routePayload, error) {
	if err != nil || q == nil {
		return nil, err
	}

	qq := &routePayload{
		Route: q,
	}

	return qq, nil
}

func (ctrl *Route) makeFilterPayload(ctx context.Context, nn types.RouteSet, f types.RouteFilter, err error) (*routeSetPayload, error) {
	if err != nil {
		return nil, err
	}

	msp := &routeSetPayload{Filter: f, Set: make([]*routePayload, len(nn))}

	for i := range nn {
		msp.Set[i], _ = ctrl.makePayload(ctx, nn[i], nil)
	}

	return msp, nil
}
