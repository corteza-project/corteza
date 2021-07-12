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
	ApigwFunction struct {
		svc functionService
		ac  templateAccessController
	}

	functionPayload struct {
		*types.ApigwFunction
	}

	functionSetPayload struct {
		Filter types.ApigwFunctionFilter `json:"filter"`
		Set    []*functionPayload        `json:"set"`
	}

	functionService interface {
		FindByID(ctx context.Context, ID uint64) (*types.ApigwFunction, error)
		Search(ctx context.Context, filter types.ApigwFunctionFilter) (types.ApigwFunctionSet, types.ApigwFunctionFilter, error)
		Create(ctx context.Context, new *types.ApigwFunction) (*types.ApigwFunction, error)
		Update(ctx context.Context, upd *types.ApigwFunction) (*types.ApigwFunction, error)
		DeleteByID(ctx context.Context, ID uint64) error
		UndeleteByID(ctx context.Context, ID uint64) error
		Definitions(context.Context, string) (interface{}, error)
	}
)

func (ApigwFunction) New() *ApigwFunction {
	return &ApigwFunction{
		svc: service.DefaultFunction,
		ac:  service.DefaultAccessControl,
	}
}

func (ctrl *ApigwFunction) List(ctx context.Context, r *request.ApigwFunctionList) (interface{}, error) {
	var (
		err error
		f   = types.ApigwFunctionFilter{
			RouteID: r.RouteID,
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

func (ctrl *ApigwFunction) Create(ctx context.Context, r *request.ApigwFunctionCreate) (interface{}, error) {
	var (
		err error
		q   = &types.ApigwFunction{
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

func (ctrl *ApigwFunction) Read(ctx context.Context, r *request.ApigwFunctionRead) (interface{}, error) {
	return ctrl.svc.FindByID(ctx, r.FunctionID)
}

func (ctrl *ApigwFunction) Update(ctx context.Context, r *request.ApigwFunctionUpdate) (interface{}, error) {
	var (
		err error
		q   = &types.ApigwFunction{
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

func (ctrl *ApigwFunction) Delete(ctx context.Context, r *request.ApigwFunctionDelete) (interface{}, error) {
	return api.OK(), ctrl.svc.DeleteByID(ctx, r.FunctionID)
}

func (ctrl *ApigwFunction) Definitions(ctx context.Context, r *request.ApigwFunctionDefinitions) (interface{}, error) {
	return ctrl.svc.Definitions(ctx, r.Kind)
}

func (ctrl *ApigwFunction) Undelete(ctx context.Context, r *request.ApigwFunctionUndelete) (interface{}, error) {
	return api.OK(), ctrl.svc.UndeleteByID(ctx, r.FunctionID)
}

func (ctrl *ApigwFunction) makePayload(ctx context.Context, q *types.ApigwFunction, err error) (*functionPayload, error) {
	if err != nil || q == nil {
		return nil, err
	}

	qq := &functionPayload{
		ApigwFunction: q,
	}

	return qq, nil
}

func (ctrl *ApigwFunction) makeFilterPayload(ctx context.Context, nn types.ApigwFunctionSet, f types.ApigwFunctionFilter, err error) (*functionSetPayload, error) {
	if err != nil {
		return nil, err
	}

	msp := &functionSetPayload{Filter: f, Set: make([]*functionPayload, len(nn))}

	for i := range nn {
		msp.Set[i], _ = ctrl.makePayload(ctx, nn[i], nil)
	}

	return msp, nil
}
