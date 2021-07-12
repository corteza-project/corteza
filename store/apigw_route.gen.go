package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/apigw_route.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	ApigwRoutes interface {
		SearchApigwRoutes(ctx context.Context, f types.ApigwRouteFilter) (types.ApigwRouteSet, types.ApigwRouteFilter, error)
		LookupApigwRouteByID(ctx context.Context, id uint64) (*types.ApigwRoute, error)
		LookupApigwRouteByEndpoint(ctx context.Context, endpoint string) (*types.ApigwRoute, error)

		CreateApigwRoute(ctx context.Context, rr ...*types.ApigwRoute) error

		UpdateApigwRoute(ctx context.Context, rr ...*types.ApigwRoute) error

		DeleteApigwRoute(ctx context.Context, rr ...*types.ApigwRoute) error
		DeleteApigwRouteByID(ctx context.Context, ID uint64) error

		TruncateApigwRoutes(ctx context.Context) error
	}
)

var _ *types.ApigwRoute
var _ context.Context

// SearchApigwRoutes returns all matching ApigwRoutes from store
func SearchApigwRoutes(ctx context.Context, s ApigwRoutes, f types.ApigwRouteFilter) (types.ApigwRouteSet, types.ApigwRouteFilter, error) {
	return s.SearchApigwRoutes(ctx, f)
}

// LookupApigwRouteByID searches for route by ID
func LookupApigwRouteByID(ctx context.Context, s ApigwRoutes, id uint64) (*types.ApigwRoute, error) {
	return s.LookupApigwRouteByID(ctx, id)
}

// LookupApigwRouteByEndpoint searches for route by endpoint
func LookupApigwRouteByEndpoint(ctx context.Context, s ApigwRoutes, endpoint string) (*types.ApigwRoute, error) {
	return s.LookupApigwRouteByEndpoint(ctx, endpoint)
}

// CreateApigwRoute creates one or more ApigwRoutes in store
func CreateApigwRoute(ctx context.Context, s ApigwRoutes, rr ...*types.ApigwRoute) error {
	return s.CreateApigwRoute(ctx, rr...)
}

// UpdateApigwRoute updates one or more (existing) ApigwRoutes in store
func UpdateApigwRoute(ctx context.Context, s ApigwRoutes, rr ...*types.ApigwRoute) error {
	return s.UpdateApigwRoute(ctx, rr...)
}

// DeleteApigwRoute Deletes one or more ApigwRoutes from store
func DeleteApigwRoute(ctx context.Context, s ApigwRoutes, rr ...*types.ApigwRoute) error {
	return s.DeleteApigwRoute(ctx, rr...)
}

// DeleteApigwRouteByID Deletes ApigwRoute from store
func DeleteApigwRouteByID(ctx context.Context, s ApigwRoutes, ID uint64) error {
	return s.DeleteApigwRouteByID(ctx, ID)
}

// TruncateApigwRoutes Deletes all ApigwRoutes from store
func TruncateApigwRoutes(ctx context.Context, s ApigwRoutes) error {
	return s.TruncateApigwRoutes(ctx)
}
