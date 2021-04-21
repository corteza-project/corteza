package service

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	internalAuth "github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/messagebus"

	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/cortezaproject/corteza-server/system/types"
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

	RBACResource interface {
		RBACResource() rbac.Resource
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

	ee.Push(types.SystemRBACResource, "grant", svc.CanGrant(ctx))
	ee.Push(types.SystemRBACResource, "auth-client.create", svc.CanCreateAuthClient(ctx))
	ee.Push(types.SystemRBACResource, "settings.read", svc.CanReadSettings(ctx))
	ee.Push(types.SystemRBACResource, "settings.manage", svc.CanManageSettings(ctx))
	ee.Push(types.SystemRBACResource, "application.create", svc.CanCreateApplication(ctx))
	ee.Push(types.SystemRBACResource, "application.flag.self", svc.CanSelfFlagApplication(ctx))
	ee.Push(types.SystemRBACResource, "application.flag.global", svc.CanGlobalFlagApplication(ctx))
	ee.Push(types.SystemRBACResource, "template.create", svc.CanCreateTemplate(ctx))
	ee.Push(types.SystemRBACResource, "role.create", svc.CanCreateRole(ctx))
	ee.Push(types.SystemRBACResource, "messagebus-queue.create", svc.CanCreateMessagebusQueue(ctx))

	return
}

func (svc accessControl) CanGrant(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "grant")
}

func (svc accessControl) CanReadSettings(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "settings.read")
}

func (svc accessControl) CanManageSettings(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "settings.manage")
}

func (svc accessControl) CanCreateUser(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "user.create")
}

func (svc accessControl) CanSearchUsers(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "user.search")
}

func (svc accessControl) CanSearchRoles(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "role.search")
}

func (svc accessControl) CanSearchApplications(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "application.search")
}

func (svc accessControl) CanCreateRole(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "role.create")
}

func (svc accessControl) CanCreateApplication(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "application.create")
}

func (svc accessControl) CanSelfFlagApplication(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "application.flag.self", rbac.Allowed)
}

func (svc accessControl) CanGlobalFlagApplication(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "application.flag.global")
}

func (svc accessControl) CanCreateAuthClient(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "auth-client.create")
}

func (svc accessControl) CanSearchAuthClient(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "auth-client.search")
}

func (svc accessControl) CanCreateTemplate(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "template.create")
}

func (svc accessControl) CanAssignReminder(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "reminder.assign")
}

func (svc accessControl) CanCreateMessagebusQueue(ctx context.Context) bool {
	return svc.can(ctx, types.SystemRBACResource, "messagebus-queue.create")
}

func (svc accessControl) CanReadRole(ctx context.Context, rl *types.Role) bool {
	return svc.can(ctx, rl.RBACResource(), "read", rbac.Allowed)
}

func (svc accessControl) CanUpdateRole(ctx context.Context, rl *types.Role) bool {
	if rl.ID == rbac.EveryoneRoleID {
		return false
	}

	return svc.can(ctx, rl.RBACResource(), "update")
}

func (svc accessControl) CanDeleteRole(ctx context.Context, rl *types.Role) bool {
	if rl.ID == rbac.EveryoneRoleID {
		return false
	}

	return svc.can(ctx, rl.RBACResource(), "delete")
}

func (svc accessControl) CanManageRoleMembers(ctx context.Context, rl *types.Role) bool {
	if rl.ID == rbac.EveryoneRoleID {
		return false
	}
	return svc.can(ctx, rl.RBACResource(), "members.manage")
}

func (svc accessControl) CanReadApplication(ctx context.Context, app *types.Application) bool {
	return svc.can(ctx, app.RBACResource(), "read", rbac.Allowed)
}

func (svc accessControl) CanUpdateApplication(ctx context.Context, app *types.Application) bool {
	return svc.can(ctx, app.RBACResource(), "update")
}

func (svc accessControl) CanDeleteApplication(ctx context.Context, app *types.Application) bool {
	return svc.can(ctx, app.RBACResource(), "delete")
}

func (svc accessControl) CanReadAuthClient(ctx context.Context, c *types.AuthClient) bool {
	return svc.can(ctx, c.RBACResource(), "read")
}

func (svc accessControl) CanUpdateAuthClient(ctx context.Context, c *types.AuthClient) bool {
	return svc.can(ctx, c.RBACResource(), "update")
}

func (svc accessControl) CanDeleteAuthClient(ctx context.Context, c *types.AuthClient) bool {
	return svc.can(ctx, c.RBACResource(), "delete")
}

func (svc accessControl) CanAuthorizeAuthClient(ctx context.Context, c *types.AuthClient) bool {
	return svc.can(ctx, c.RBACResource(), "authorize")
}

func (svc accessControl) CanReadTemplate(ctx context.Context, tpl *types.Template) bool {
	return svc.can(ctx, tpl.RBACResource(), "read", rbac.Allowed)
}

func (svc accessControl) CanUpdateTemplate(ctx context.Context, tpl *types.Template) bool {
	return svc.can(ctx, tpl.RBACResource(), "update")
}

func (svc accessControl) CanDeleteTemplate(ctx context.Context, tpl *types.Template) bool {
	return svc.can(ctx, tpl.RBACResource(), "delete")
}

func (svc accessControl) CanRenderTemplate(ctx context.Context, tpl *types.Template) bool {
	return svc.can(ctx, tpl.RBACResource(), "render", rbac.Allowed)
}

func (svc accessControl) CanReadUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "read")
}

func (svc accessControl) CanUpdateUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "update")
}

func (svc accessControl) CanSuspendUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "suspend")
}

func (svc accessControl) CanUnsuspendUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "unsuspend")
}

func (svc accessControl) CanDeleteUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "delete")
}

func (svc accessControl) CanImpersonateUser(ctx context.Context, u *types.User) bool {
	return svc.can(ctx, u.RBACResource(), "impersonate", rbac.Denied)
}

func (svc accessControl) CanUnmaskEmail(ctx context.Context, u *types.User) bool {
	if internalAuth.GetIdentityFromContext(ctx).Identity() == u.ID {
		// Make an exception when users are reading their own info
		return true
	}

	return svc.can(ctx, u.RBACResource(), "unmask.email")
}

func (svc accessControl) CanUnmaskName(ctx context.Context, u *types.User) bool {
	if internalAuth.GetIdentityFromContext(ctx).Identity() == u.ID {
		// Make an exception when users are reading their own info
		return true
	}

	return svc.can(ctx, u.RBACResource(), "unmask.name")
}

func (svc accessControl) CanReadMessagebusQueue(ctx context.Context, c *messagebus.QueueSettings) bool {
	return svc.can(ctx, c.RBACResource(), "read")
}

func (svc accessControl) CanUpdateMessagebusQueue(ctx context.Context, c *messagebus.QueueSettings) bool {
	return svc.can(ctx, c.RBACResource(), "update")
}

func (svc accessControl) CanDeleteMessagebusQueue(ctx context.Context, c *messagebus.QueueSettings) bool {
	return svc.can(ctx, c.RBACResource(), "delete")
}

func (svc accessControl) CanReadFromMessagebusQueue(ctx context.Context, c *messagebus.QueueSettings) bool {
	return svc.can(ctx, c.RBACResource(), "queue.read")
}

func (svc accessControl) CanWriteToMessagebusQueue(ctx context.Context, c *messagebus.QueueSettings) bool {
	return svc.can(ctx, c.RBACResource(), "queue.write")
}

func (svc accessControl) can(ctx context.Context, res rbac.Resource, op rbac.Operation, ff ...rbac.CheckAccessFunc) bool {
	var (
		u     = internalAuth.GetIdentityFromContext(ctx)
		roles = u.Roles()
	)

	if internalAuth.IsSuperUser(u) {
		// Temp solution to allow migration from passing context to ResourceFilter
		// and checking "superuser" privileges there to more sustainable solution
		// (eg: creating super-role with allow-all)
		return true
	}

	return svc.permissions.Can(roles, res.RBACResource(), op, ff...)
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
		types.SystemRBACResource,
		"grant",
		"settings.read",
		"settings.manage",
		"auth-client.create",
		"auth-client.search",
		"role.create",
		"role.search",
		"user.create",
		"user.search",
		"application.create",
		"application.search",
		"application.flag.self",
		"application.flag.global",
		"template.create",
		"template.search",
		"reminder.assign",
		"messagebus-queue.create",
	)

	wl.Set(
		types.ApplicationRBACResource,
		"read",
		"update",
		"delete",
	)

	wl.Set(
		types.TemplateRBACResource,
		"read",
		"update",
		"delete",
		"render",
	)

	wl.Set(
		types.UserRBACResource,
		"read",
		"update",
		"delete",
		"suspend",
		"unsuspend",
		"unmask.email",
		"unmask.name",
		"impersonate",
	)

	wl.Set(
		types.RoleRBACResource,
		"read",
		"update",
		"delete",
		"members.manage",
	)

	wl.Set(
		types.AuthClientRBACResource,
		"read",
		"update",
		"delete",
		"authorize",
	)

	wl.Set(
		types.MessagebusQueueRBACResource,
		"read",
		"update",
		"delete",
		"queue.read",
		"queue.write",
	)

	return wl
}
