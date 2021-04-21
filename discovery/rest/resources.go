package rest

import (
	"context"
	"github.com/cortezaproject/corteza-server/discovery/rest/internal/documents"
	"github.com/cortezaproject/corteza-server/discovery/rest/request"
)

type (
	resources struct {
		sys interface {
			Users(ctx context.Context, limit uint, cur string) (*documents.Response, error)
		}

		cmp interface {
			Namespaces(ctx context.Context, limit uint, cur string) (*documents.Response, error)
			Modules(ctx context.Context, namespaceID uint64, limit uint, cur string) (*documents.Response, error)
			Records(ctx context.Context, namespaceID, moduleID uint64, limit uint, cur string) (*documents.Response, error)
		}
	}
)

func Resources() *resources {
	return &resources{
		sys: documents.SystemResources(),
		cmp: documents.ComposeResources(),
	}
}

func (ctrl resources) SystemUsers(ctx context.Context, r *request.ResourcesSystemUsers) (interface{}, error) {
	return ctrl.sys.Users(ctx, r.Limit, r.PageCursor)
}

func (ctrl resources) ComposeNamespaces(ctx context.Context, r *request.ResourcesComposeNamespaces) (interface{}, error) {
	return ctrl.cmp.Namespaces(ctx, r.Limit, r.PageCursor)
}

func (ctrl resources) ComposeModules(ctx context.Context, r *request.ResourcesComposeModules) (interface{}, error) {
	return ctrl.cmp.Modules(ctx, r.NamespaceID, r.Limit, r.PageCursor)
}

func (ctrl resources) ComposeRecords(ctx context.Context, r *request.ResourcesComposeRecords) (interface{}, error) {
	return ctrl.cmp.Records(ctx, r.NamespaceID, r.ModuleID, r.Limit, r.PageCursor)
}
