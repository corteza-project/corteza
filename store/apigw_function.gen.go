package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/apigw_function.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	ApigwFunctions interface {
		SearchApigwFunctions(ctx context.Context, f types.FunctionFilter) (types.FunctionSet, types.FunctionFilter, error)
		LookupApigwFunctionByID(ctx context.Context, id uint64) (*types.Function, error)
		LookupApigwFunctionByRoute(ctx context.Context, route string) (*types.Function, error)

		CreateApigwFunction(ctx context.Context, rr ...*types.Function) error

		UpdateApigwFunction(ctx context.Context, rr ...*types.Function) error

		DeleteApigwFunction(ctx context.Context, rr ...*types.Function) error
		DeleteApigwFunctionByID(ctx context.Context, ID uint64) error

		TruncateApigwFunctions(ctx context.Context) error
	}
)

var _ *types.Function
var _ context.Context

// SearchApigwFunctions returns all matching ApigwFunctions from store
func SearchApigwFunctions(ctx context.Context, s ApigwFunctions, f types.FunctionFilter) (types.FunctionSet, types.FunctionFilter, error) {
	return s.SearchApigwFunctions(ctx, f)
}

// LookupApigwFunctionByID searches for function by ID
func LookupApigwFunctionByID(ctx context.Context, s ApigwFunctions, id uint64) (*types.Function, error) {
	return s.LookupApigwFunctionByID(ctx, id)
}

// LookupApigwFunctionByRoute searches for function by route
func LookupApigwFunctionByRoute(ctx context.Context, s ApigwFunctions, route string) (*types.Function, error) {
	return s.LookupApigwFunctionByRoute(ctx, route)
}

// CreateApigwFunction creates one or more ApigwFunctions in store
func CreateApigwFunction(ctx context.Context, s ApigwFunctions, rr ...*types.Function) error {
	return s.CreateApigwFunction(ctx, rr...)
}

// UpdateApigwFunction updates one or more (existing) ApigwFunctions in store
func UpdateApigwFunction(ctx context.Context, s ApigwFunctions, rr ...*types.Function) error {
	return s.UpdateApigwFunction(ctx, rr...)
}

// DeleteApigwFunction Deletes one or more ApigwFunctions from store
func DeleteApigwFunction(ctx context.Context, s ApigwFunctions, rr ...*types.Function) error {
	return s.DeleteApigwFunction(ctx, rr...)
}

// DeleteApigwFunctionByID Deletes ApigwFunction from store
func DeleteApigwFunctionByID(ctx context.Context, s ApigwFunctions, ID uint64) error {
	return s.DeleteApigwFunctionByID(ctx, ID)
}

// TruncateApigwFunctions Deletes all ApigwFunctions from store
func TruncateApigwFunctions(ctx context.Context, s ApigwFunctions) error {
	return s.TruncateApigwFunctions(ctx)
}
