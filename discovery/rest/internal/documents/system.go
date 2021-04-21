package documents

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/cortezaproject/corteza-server/system/service"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	systemResources struct {
		settings *types.AppSettings

		rbac interface {
			SignificantRoles(res rbac.Resource, op rbac.Operation) (aRR, dRR []uint64)
		}

		ac interface {
			CanReadUser(ctx context.Context, r *types.User) bool
		}
		usr interface {
			Find(context.Context, types.UserFilter) (types.UserSet, types.UserFilter, error)
		}
	}
)

func SystemResources() *systemResources {
	return &systemResources{
		settings: service.CurrentSettings,
		rbac:     rbac.Global(),
		ac:       service.DefaultAccessControl,
		usr:      service.DefaultUser,
	}
}

func (d systemResources) Users(ctx context.Context, limit uint, cur string) (rsp *Response, err error) {
	return rsp, func() (err error) {
		if !d.settings.Discovery.SystemUsers.Enabled {
			return nil
		}

		var (
			uu types.UserSet
			f  = types.UserFilter{}
		)

		if f.Paging, err = filter.NewPaging(limit, cur); err != nil {
			return err
		}

		if uu, f, err = d.usr.Find(ctx, f); err != nil {
			return err
		}

		rsp = &Response{
			Documents: make([]Document, len(uu)),
			Filter: Filter{
				Limit:    limit,
				NextPage: f.NextPage,
			},
		}

		for i, u := range uu {
			doc := &docUser{
				ResourceType: "system:user",
				UserID:       u.ID,
				Email:        u.Email,
				Name:         u.Name,
				Handle:       u.Handle,
				Suspended:    u.SuspendedAt,
				Created:      makePartialChange(&u.CreatedAt),
				Updated:      makePartialChange(u.UpdatedAt),
				Deleted:      makePartialChange(u.DeletedAt),
			}

			doc.Security.AllowedRoles, doc.Security.DeniedRoles = d.rbac.SignificantRoles(u.RBACResource(), "read")

			rsp.Documents[i].ID = u.ID
			rsp.Documents[i].URL = "@todo"
			rsp.Documents[i].Source = doc
		}

		return nil
	}()
}
