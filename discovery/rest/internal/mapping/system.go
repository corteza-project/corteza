package mapping

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/system/service"
)

type (
	systemAccessControl interface {
		CanSearchUsers(ctx context.Context) bool
		CanSearchRoles(ctx context.Context) bool
		CanSearchApplications(ctx context.Context) bool
		CanSearchAuthClient(ctx context.Context) bool
	}

	systemMapping struct {
		ac systemAccessControl
	}
)

func SystemMapping() *systemMapping {
	return &systemMapping{
		ac: service.DefaultAccessControl,
	}
}

func (svc systemMapping) Users(_ context.Context) ([]*Mapping, error) {
	return []*Mapping{{
		Index: fmt.Sprintf("users"),
		Mapping: map[string]*property{
			"resourceType": {Type: "keyword"},

			"userID": {Type: "unsigned_long"},

			"email":  {Type: "keyword", Boost: 2},
			"name":   {Type: "keyword", Boost: 10},
			"handle": {Type: "keyword", Boost: 2},

			"created": change(),
			"updated": change(),
			"deleted": change(),

			"suspendedAt": {Type: "date"},

			"security": security(),
		},
	}}, nil
}
