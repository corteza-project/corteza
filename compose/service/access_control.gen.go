package service

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

// Definitions file that controls how this file is generated:
// - compose.chart.yaml
// - compose.module-field.yaml
// - compose.module.yaml
// - compose.namespace.yaml
// - compose.page.yaml
// - compose.record.yaml
// - compose.yaml

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/spf13/cast"
	"strings"
)

type (
	accessControl struct {
		actionlog actionlog.Recorder

		rbac interface {
			Can(rbac.Session, string, rbac.Resource) bool
			Grant(context.Context, ...*rbac.Rule) error
			FindRulesByRoleID(roleID uint64) (rr rbac.RuleSet)
		}
	}
)

func AccessControl() *accessControl {
	return &accessControl{
		rbac:      rbac.Global(),
		actionlog: DefaultActionlog,
	}
}

func (svc accessControl) can(ctx context.Context, op string, res rbac.Resource) bool {
	return svc.rbac.Can(rbac.ContextToSession(ctx), op, res)
}

// Effective returns a list of effective permissions for all given resource
func (svc accessControl) Effective(ctx context.Context, rr ...rbac.Resource) (ee rbac.EffectiveSet) {
	for _, res := range rr {
		r := res.RbacResource()
		for op := range rbacResourceOperations(r) {
			ee.Push(r, op, svc.can(ctx, op, res))
		}
	}

	return
}

func (svc accessControl) List() (out []map[string]string) {
	def := []map[string]string{
		{
			"type": types.ChartResourceType,
			"any":  types.ChartRbacResource(0, 0),
			"op":   "read",
		},
		{
			"type": types.ChartResourceType,
			"any":  types.ChartRbacResource(0, 0),
			"op":   "update",
		},
		{
			"type": types.ChartResourceType,
			"any":  types.ChartRbacResource(0, 0),
			"op":   "delete",
		},
		{
			"type": types.ModuleFieldResourceType,
			"any":  types.ModuleFieldRbacResource(0, 0, 0),
			"op":   "record.value.read",
		},
		{
			"type": types.ModuleFieldResourceType,
			"any":  types.ModuleFieldRbacResource(0, 0, 0),
			"op":   "record.value.update",
		},
		{
			"type": types.ModuleResourceType,
			"any":  types.ModuleRbacResource(0, 0),
			"op":   "read",
		},
		{
			"type": types.ModuleResourceType,
			"any":  types.ModuleRbacResource(0, 0),
			"op":   "update",
		},
		{
			"type": types.ModuleResourceType,
			"any":  types.ModuleRbacResource(0, 0),
			"op":   "delete",
		},
		{
			"type": types.ModuleResourceType,
			"any":  types.ModuleRbacResource(0, 0),
			"op":   "record.create",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "read",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "update",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "delete",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "manage",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "module.create",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "chart.create",
		},
		{
			"type": types.NamespaceResourceType,
			"any":  types.NamespaceRbacResource(0),
			"op":   "page.create",
		},
		{
			"type": types.PageResourceType,
			"any":  types.PageRbacResource(0, 0),
			"op":   "read",
		},
		{
			"type": types.PageResourceType,
			"any":  types.PageRbacResource(0, 0),
			"op":   "update",
		},
		{
			"type": types.PageResourceType,
			"any":  types.PageRbacResource(0, 0),
			"op":   "delete",
		},
		{
			"type": types.RecordResourceType,
			"any":  types.RecordRbacResource(0, 0, 0),
			"op":   "read",
		},
		{
			"type": types.RecordResourceType,
			"any":  types.RecordRbacResource(0, 0, 0),
			"op":   "update",
		},
		{
			"type": types.RecordResourceType,
			"any":  types.RecordRbacResource(0, 0, 0),
			"op":   "delete",
		},
		{
			"type": types.ComponentResourceType,
			"any":  types.ComponentRbacResource(),
			"op":   "grant",
		},
		{
			"type": types.ComponentResourceType,
			"any":  types.ComponentRbacResource(),
			"op":   "namespace.create",
		},
		{
			"type": types.ComponentResourceType,
			"any":  types.ComponentRbacResource(),
			"op":   "settings.read",
		},
		{
			"type": types.ComponentResourceType,
			"any":  types.ComponentRbacResource(),
			"op":   "settings.manage",
		},
	}

	func(svc interface{}) {
		if svc, is := svc.(interface{}).(interface{ list() []map[string]string }); is {
			def = append(def, svc.list()...)
		}
	}(svc)

	return def
}

// Grant applies one or more RBAC rules
//
// This function is auto-generated
func (svc accessControl) Grant(ctx context.Context, rr ...*rbac.Rule) error {
	if !svc.CanGrant(ctx) {
		// @todo should be altered to check grant permissions PER resource
		return AccessControlErrNotAllowedToSetPermissions()
	}

	for _, r := range rr {
		err := rbacResourceValidator(r.Resource, r.Operation)
		if err != nil {
			return err
		}
	}

	if err := svc.rbac.Grant(ctx, rr...); err != nil {
		return AccessControlErrGeneric().Wrap(err)
	}

	svc.logGrants(ctx, rr)

	return nil
}

// This function is auto-generated
func (svc accessControl) logGrants(ctx context.Context, rr []*rbac.Rule) {
	if svc.actionlog == nil {
		return
	}

	for _, r := range rr {
		g := AccessControlActionGrant(&accessControlActionProps{r})
		g.log = r.String()
		g.resource = r.Resource

		svc.actionlog.Record(ctx, g.ToAction())
	}
}

// FindRulesByRoleID find all rules for a specific role
//
// This function is auto-generated
func (svc accessControl) FindRulesByRoleID(ctx context.Context, roleID uint64) (rbac.RuleSet, error) {
	if !svc.CanGrant(ctx) {
		return nil, AccessControlErrNotAllowedToSetPermissions()
	}

	return svc.rbac.FindRulesByRoleID(roleID), nil
}

// CanReadChart checks if current user can read chart
//
// This function is auto-generated
func (svc accessControl) CanReadChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, "read", r)
}

// CanUpdateChart checks if current user can update chart
//
// This function is auto-generated
func (svc accessControl) CanUpdateChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, "update", r)
}

// CanDeleteChart checks if current user can delete chart
//
// This function is auto-generated
func (svc accessControl) CanDeleteChart(ctx context.Context, r *types.Chart) bool {
	return svc.can(ctx, "delete", r)
}

// CanReadRecordValue checks if current user can read field value on records
//
// This function is auto-generated
func (svc accessControl) CanReadRecordValue(ctx context.Context, r *types.ModuleField) bool {
	return svc.can(ctx, "record.value.read", r)
}

// CanUpdateRecordValue checks if current user can update field value on records
//
// This function is auto-generated
func (svc accessControl) CanUpdateRecordValue(ctx context.Context, r *types.ModuleField) bool {
	return svc.can(ctx, "record.value.update", r)
}

// CanReadModule checks if current user can read module
//
// This function is auto-generated
func (svc accessControl) CanReadModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, "read", r)
}

// CanUpdateModule checks if current user can update module
//
// This function is auto-generated
func (svc accessControl) CanUpdateModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, "update", r)
}

// CanDeleteModule checks if current user can delete module
//
// This function is auto-generated
func (svc accessControl) CanDeleteModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, "delete", r)
}

// CanCreateRecordOnModule checks if current user can create record
//
// This function is auto-generated
func (svc accessControl) CanCreateRecordOnModule(ctx context.Context, r *types.Module) bool {
	return svc.can(ctx, "record.create", r)
}

// CanReadNamespace checks if current user can read namespace
//
// This function is auto-generated
func (svc accessControl) CanReadNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "read", r)
}

// CanUpdateNamespace checks if current user can update namespace
//
// This function is auto-generated
func (svc accessControl) CanUpdateNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "update", r)
}

// CanDeleteNamespace checks if current user can delete namespace
//
// This function is auto-generated
func (svc accessControl) CanDeleteNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "delete", r)
}

// CanManageNamespace checks if current user can access to namespace admin panel
//
// This function is auto-generated
func (svc accessControl) CanManageNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "manage", r)
}

// CanCreateModuleOnNamespace checks if current user can create module on namespace
//
// This function is auto-generated
func (svc accessControl) CanCreateModuleOnNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "module.create", r)
}

// CanCreateChartOnNamespace checks if current user can create chart on namespace
//
// This function is auto-generated
func (svc accessControl) CanCreateChartOnNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "chart.create", r)
}

// CanCreatePageOnNamespace checks if current user can create page on namespace
//
// This function is auto-generated
func (svc accessControl) CanCreatePageOnNamespace(ctx context.Context, r *types.Namespace) bool {
	return svc.can(ctx, "page.create", r)
}

// CanReadPage checks if current user can read page
//
// This function is auto-generated
func (svc accessControl) CanReadPage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, "read", r)
}

// CanUpdatePage checks if current user can update page
//
// This function is auto-generated
func (svc accessControl) CanUpdatePage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, "update", r)
}

// CanDeletePage checks if current user can delete page
//
// This function is auto-generated
func (svc accessControl) CanDeletePage(ctx context.Context, r *types.Page) bool {
	return svc.can(ctx, "delete", r)
}

// CanReadRecord checks if current user can read record
//
// This function is auto-generated
func (svc accessControl) CanReadRecord(ctx context.Context, r *types.Record) bool {
	return svc.can(ctx, "read", r)
}

// CanUpdateRecord checks if current user can update record
//
// This function is auto-generated
func (svc accessControl) CanUpdateRecord(ctx context.Context, r *types.Record) bool {
	return svc.can(ctx, "update", r)
}

// CanDeleteRecord checks if current user can delete record
//
// This function is auto-generated
func (svc accessControl) CanDeleteRecord(ctx context.Context, r *types.Record) bool {
	return svc.can(ctx, "delete", r)
}

// CanGrant checks if current user can manage compose permissions
//
// This function is auto-generated
func (svc accessControl) CanGrant(ctx context.Context) bool {
	return svc.can(ctx, "grant", &types.Component{})
}

// CanCreateNamespace checks if current user can create namespace
//
// This function is auto-generated
func (svc accessControl) CanCreateNamespace(ctx context.Context) bool {
	return svc.can(ctx, "namespace.create", &types.Component{})
}

// CanReadSettings checks if current user can read settings
//
// This function is auto-generated
func (svc accessControl) CanReadSettings(ctx context.Context) bool {
	return svc.can(ctx, "settings.read", &types.Component{})
}

// CanManageSettings checks if current user can manage settings
//
// This function is auto-generated
func (svc accessControl) CanManageSettings(ctx context.Context) bool {
	return svc.can(ctx, "settings.manage", &types.Component{})
}

// rbacResourceValidator validates known component's resource by routing it to the appropriate validator
//
// This function is auto-generated
func rbacResourceValidator(r string, oo ...string) error {
	switch rbac.ResourceType(r) {
	case types.ChartResourceType:
		return rbacChartResourceValidator(r, oo...)
	case types.ModuleFieldResourceType:
		return rbacModuleFieldResourceValidator(r, oo...)
	case types.ModuleResourceType:
		return rbacModuleResourceValidator(r, oo...)
	case types.NamespaceResourceType:
		return rbacNamespaceResourceValidator(r, oo...)
	case types.PageResourceType:
		return rbacPageResourceValidator(r, oo...)
	case types.RecordResourceType:
		return rbacRecordResourceValidator(r, oo...)
	case types.ComponentResourceType:
		return rbacComponentResourceValidator(r, oo...)
	}

	return fmt.Errorf("unknown resource type '%q'", r)
}

// rbacResourceOperations returns defined operations for a requested resource
//
// This function is auto-generated
func rbacResourceOperations(r string) map[string]bool {
	switch rbac.ResourceType(r) {
	case types.ChartResourceType:
		return map[string]bool{
			"read":   true,
			"update": true,
			"delete": true,
		}
	case types.ModuleFieldResourceType:
		return map[string]bool{
			"record.value.read":   true,
			"record.value.update": true,
		}
	case types.ModuleResourceType:
		return map[string]bool{
			"read":          true,
			"update":        true,
			"delete":        true,
			"record.create": true,
		}
	case types.NamespaceResourceType:
		return map[string]bool{
			"read":          true,
			"update":        true,
			"delete":        true,
			"manage":        true,
			"module.create": true,
			"chart.create":  true,
			"page.create":   true,
		}
	case types.PageResourceType:
		return map[string]bool{
			"read":   true,
			"update": true,
			"delete": true,
		}
	case types.RecordResourceType:
		return map[string]bool{
			"read":   true,
			"update": true,
			"delete": true,
		}
	case types.ComponentResourceType:
		return map[string]bool{
			"grant":            true,
			"namespace.create": true,
			"settings.read":    true,
			"settings.manage":  true,
		}
	}

	return nil
}

// rbacChartResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacChartResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose Chart resource", o)
		}
	}

	if !strings.HasPrefix(r, types.ChartResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.ChartResourceType):], sep), sep)
		prc = []string{
			"namespaceID",
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for Chart", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacModuleFieldResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacModuleFieldResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose ModuleField resource", o)
		}
	}

	if !strings.HasPrefix(r, types.ModuleFieldResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.ModuleFieldResourceType):], sep), sep)
		prc = []string{
			"namespaceID",
			"moduleID",
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for ModuleField", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacModuleResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacModuleResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose Module resource", o)
		}
	}

	if !strings.HasPrefix(r, types.ModuleResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.ModuleResourceType):], sep), sep)
		prc = []string{
			"namespaceID",
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for Module", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacNamespaceResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacNamespaceResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose Namespace resource", o)
		}
	}

	if !strings.HasPrefix(r, types.NamespaceResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.NamespaceResourceType):], sep), sep)
		prc = []string{
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for Namespace", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacPageResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacPageResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose Page resource", o)
		}
	}

	if !strings.HasPrefix(r, types.PageResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.PageResourceType):], sep), sep)
		prc = []string{
			"namespaceID",
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for Page", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacRecordResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacRecordResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose Record resource", o)
		}
	}

	if !strings.HasPrefix(r, types.RecordResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	const sep = "/"
	var (
		specIdUsed = true

		pp  = strings.Split(strings.Trim(r[len(types.RecordResourceType):], sep), sep)
		prc = []string{
			"namespaceID",
			"moduleID",
			"ID",
		}
	)

	if len(pp) != len(prc) {
		return fmt.Errorf("invalid resource path structure")
	}

	for i, p := range pp {
		if p == "*" {
			if !specIdUsed {
				return fmt.Errorf("invalid resource path wildcard level (%d) for Record", i)
			}

			specIdUsed = false
			continue
		}

		specIdUsed = true
		if _, err := cast.ToUint64E(p); err != nil {
			return fmt.Errorf("invalid reference for %s: '%s'", prc[i], p)
		}
	}

	return nil
}

// rbacComponentResourceValidator checks validity of rbac resource and operations
//
// Can be called without operations to check for validity of resource string only
//
// This function is auto-generated
func rbacComponentResourceValidator(r string, oo ...string) error {
	defOps := rbacResourceOperations(r)
	for _, o := range oo {
		if !defOps[o] {
			return fmt.Errorf("invalid operation '%s' for compose resource", o)
		}
	}

	if !strings.HasPrefix(r, types.ComponentResourceType) {
		// expecting resource to always include path
		return fmt.Errorf("invalid resource type")
	}

	return nil
}
