package service

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
)

type (
	accessControl struct {
		permissions accessControlRBACServicer
		actionlog   actionlog.Recorder
	}

	accessControlRBACServicer interface {
		Can([]uint64, rbac.Resource, rbac.Operation, ...rbac.CheckAccessFunc) bool
		Grant(context.Context, rbac.Whitelist, ...*rbac.Rule) error
		FindRulesByRoleID(roleID uint64) (rr rbac.RuleSet)
	}

	secureResource interface {
		RBACResource() rbac.Resource
		DynamicRoles(uint64) []uint64
	}
)

func AccessControl(perm accessControlRBACServicer) *accessControl {
	return &accessControl{
		permissions: perm,
		actionlog:   DefaultActionlog,
	}
}

// Effective returns a list of effective service-level permissions
func (svc accessControl) Effective(ctx context.Context) (ee rbac.EffectiveSet) {
	ee = rbac.EffectiveSet{}

	ee.Push(types.ComposeRBACResource, "grant", svc.CanGrant(ctx))
	ee.Push(types.ComposeRBACResource, "namespace.create", svc.CanCreateNamespace(ctx))
	ee.Push(types.ComposeRBACResource, "settings.read", svc.CanReadSettings(ctx))
	ee.Push(types.ComposeRBACResource, "settings.manage", svc.CanManageSettings(ctx))

	return
}

func (svc accessControl) CanGrant(ctx context.Context) bool {
	return svc.can(ctx, types.ComposeRBACResource, "grant")
}

func (svc accessControl) CanReadSettings(ctx context.Context) bool {
	return svc.can(ctx, types.ComposeRBACResource, "settings.read")
}

func (svc accessControl) CanManageSettings(ctx context.Context) bool {
	return svc.can(ctx, types.ComposeRBACResource, "settings.manage")
}

func (svc accessControl) CanCreateNamespace(ctx context.Context) bool {
	return svc.can(ctx, types.ComposeRBACResource, "namespace.create")
}

func (svc accessControl) CanSearchNamespaces(ctx context.Context) bool {
	return svc.can(ctx, types.ComposeRBACResource, "namespace.search")
}

func (svc accessControl) CanReadNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "read", rbac.Allowed)
}

func (svc accessControl) CanUpdateNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "update")
}

func (svc accessControl) CanDeleteNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "delete")
}

func (svc accessControl) CanManageNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "manage")
}

func (svc accessControl) CanCreateModule(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "module.create")
}

func (svc accessControl) CanSearchModules(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "module.search")
}

func (svc accessControl) CanReadModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "read")
}

func (svc accessControl) CanUpdateModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "update")
}

func (svc accessControl) CanDeleteModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "delete")
}

func (svc accessControl) CanReadRecordValue(ctx context.Context, r *types.ModuleField) bool {
	return svc.can(ctx, r, "record.value.read", rbac.Allowed)
}

func (svc accessControl) CanUpdateRecordValue(ctx context.Context, r *types.ModuleField) bool {
	return svc.can(ctx, r, "record.value.update", rbac.Allowed)
}

func (svc accessControl) CanCreateRecord(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "record.create")
}

func (svc accessControl) CanReadRecord(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "record.read")
}

func (svc accessControl) CanSearchRecords(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "record.search")
}

func (svc accessControl) CanUpdateRecord(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "record.update")
}

func (svc accessControl) CanDeleteRecord(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, r, "record.delete")
}

func (svc accessControl) CanCreateChart(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "chart.create")
}

func (svc accessControl) CanSearchChart(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "chart.search")
}

func (svc accessControl) CanReadChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, r, "read")
}

func (svc accessControl) CanUpdateChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, r, "update")
}

func (svc accessControl) CanDeleteChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, r, "delete")
}

func (svc accessControl) CanCreatePage(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "page.create")
}

func (svc accessControl) CanSearchPage(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, r, "page.search")
}

func (svc accessControl) CanReadPage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, r, "read")
}

func (svc accessControl) CanUpdatePage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, r, "update")
}

func (svc accessControl) CanDeletePage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, r, "delete")
}

func (svc accessControl) can(ctx context.Context, res secureResource, op rbac.Operation, ff ...rbac.CheckAccessFunc) bool {
	var u = auth.GetIdentityFromContext(ctx)

	if auth.IsSuperUser(u) {
		// Temp solution to allow migration from passing context to ResourceFilter
		// and checking "superuser" privileges there to more sustainable solution
		// (eg: creating super-role with allow-all)
		return true
	}

	return svc.permissions.Can(
		append(u.Roles(), res.DynamicRoles(u.Identity())...),
		res.RBACResource(),
		op,
		ff...,
	)
}

func (svc accessControl) Grant(ctx context.Context, rr ...*rbac.Rule) error {
	if !svc.CanGrant(ctx) {
		return AccessControlErrNotAllowedToSetPermissions()
	}

	if err := svc.permissions.Grant(ctx, svc.Whitelist(), rr...); err != nil {
		return AccessControlErrGeneric().Wrap(err)
	}

	svc.logGrants(ctx, rr)

	return nil
}

func (svc accessControl) logGrants(ctx context.Context, rr []*rbac.Rule) {
	if svc.actionlog == nil {
		return
	}

	for _, r := range rr {
		g := AccessControlActionGrant(&accessControlActionProps{r})
		g.log = r.String()
		g.resource = r.Resource.String()

		svc.actionlog.Record(ctx, g.ToAction())
	}
}

func (svc accessControl) FindRulesByRoleID(ctx context.Context, roleID uint64) (rbac.RuleSet, error) {
	if !svc.CanGrant(ctx) {
		return nil, AccessControlErrNotAllowedToSetPermissions()
	}

	return svc.permissions.FindRulesByRoleID(roleID), nil
}

func (svc accessControl) Whitelist() rbac.Whitelist {
	var wl = rbac.Whitelist{}

	wl.Set(
		types.ComposeRBACResource,
		"grant",
		"namespace.create",
		"namespace.search",
		"settings.read",
		"settings.manage",
	)

	wl.Set(
		types.NamespaceRBACResource,
		"read",
		"update",
		"delete",
		"manage",
		"module.create",
		"module.search",
		"chart.create",
		"chart.search",
		"page.create",
		"page.search",
	)

	wl.Set(
		types.ModuleRBACResource,
		"read",
		"update",
		"delete",
		"record.create",
		"record.read",
		"record.update",
		"record.delete",
		"record.search",
	)

	wl.Set(
		types.ModuleFieldRBACResource,
		"record.value.read",
		"record.value.update",
	)

	wl.Set(
		types.ChartRBACResource,
		"read",
		"update",
		"delete",
	)

	wl.Set(
		types.PageRBACResource,
		"read",
		"update",
		"delete",
	)

	return wl
}
