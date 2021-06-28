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
	Function struct {
		svc functionService
		ac  templateAccessController
	}

	functionPayload struct {
		*types.Function
	}

	functionSetPayload struct {
		Filter types.FunctionFilter `json:"filter"`
		Set    []*functionPayload   `json:"set"`
	}

	functionService interface {
		FindByID(ctx context.Context, ID uint64) (*types.Function, error)
		Search(ctx context.Context, filter types.FunctionFilter) (types.FunctionSet, types.FunctionFilter, error)
		Create(ctx context.Context, new *types.Function) (*types.Function, error)
		Update(ctx context.Context, upd *types.Function) (*types.Function, error)
		DeleteByID(ctx context.Context, ID uint64) error
		UndeleteByID(ctx context.Context, ID uint64) error
	}
)

func (Function) New() *Function {
	return &Function{
		svc: service.DefaultFunction,
		ac:  service.DefaultAccessControl,
	}
}

func (ctrl *Function) List(ctx context.Context, r *request.FunctionList) (interface{}, error) {
	var (
		err error
		f   = types.FunctionFilter{
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

func (ctrl *Function) Create(ctx context.Context, r *request.FunctionCreate) (interface{}, error) {
	var (
		err error
		q   = &types.Function{
			Route:  r.RouteID,
			Weight: r.Weight,
			Kind:   string(r.Kind),
			Ref:    r.Ref,
			Params: r.Params,
		}
	)

	q, err = ctrl.svc.Create(ctx, q)

	return ctrl.makePayload(ctx, q, err)
}

func (ctrl *Function) Read(ctx context.Context, r *request.FunctionRead) (interface{}, error) {
	return ctrl.svc.FindByID(ctx, r.FunctionID)
}

func (ctrl *Function) Update(ctx context.Context, r *request.FunctionUpdate) (interface{}, error) {
	var (
		err error
		q   = &types.Function{
			ID:     r.FunctionID,
			Route:  r.RouteID,
			Weight: r.Weight,
			Kind:   string(r.Kind),
			Ref:    r.Ref,
			Params: r.Params,
		}
	)

	q, err = ctrl.svc.Update(ctx, q)

	return ctrl.makePayload(ctx, q, err)
}

func (ctrl *Function) Delete(ctx context.Context, r *request.FunctionDelete) (interface{}, error) {
	return api.OK(), ctrl.svc.DeleteByID(ctx, r.FunctionID)
}

func (ctrl *Function) Undelete(ctx context.Context, r *request.FunctionUndelete) (interface{}, error) {
	return api.OK(), ctrl.svc.UndeleteByID(ctx, r.FunctionID)
}

func (ctrl *Function) makePayload(ctx context.Context, q *types.Function, err error) (*functionPayload, error) {
	if err != nil || q == nil {
		return nil, err
	}

	qq := &functionPayload{
		Function: q,
	}

	return qq, nil
}

func (ctrl *Function) makeFilterPayload(ctx context.Context, nn types.FunctionSet, f types.FunctionFilter, err error) (*functionSetPayload, error) {
	if err != nil {
		return nil, err
	}

	msp := &functionSetPayload{Filter: f, Set: make([]*functionPayload, len(nn))}

	for i := range nn {
		msp.Set[i], _ = ctrl.makePayload(ctx, nn[i], nil)
	}

	return msp, nil
}
