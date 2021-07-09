package resource

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

// Definitions file that controls how this file is generated:
// - system.application.yaml
// - system.auth-client.yaml
// - system.role.yaml
// - system.route.yaml
// - system.template.yaml
// - system.user.yaml
// - system.yaml

import (
	"github.com/cortezaproject/corteza-server/system/types"
)

// SystemApplicationRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemApplicationRbacReferences(application string) (res *Ref, pp []*Ref, err error) {
	if application != "*" {
		res = &Ref{ResourceType: types.ApplicationResourceType, Identifiers: MakeIdentifiers(application)}
	}

	return
}

// SystemAuthClientRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemAuthClientRbacReferences(authClient string) (res *Ref, pp []*Ref, err error) {
	if authClient != "*" {
		res = &Ref{ResourceType: types.AuthClientResourceType, Identifiers: MakeIdentifiers(authClient)}
	}

	return
}

// SystemRoleRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemRoleRbacReferences(role string) (res *Ref, pp []*Ref, err error) {
	if role != "*" {
		res = &Ref{ResourceType: types.RoleResourceType, Identifiers: MakeIdentifiers(role)}
	}

	return
}

// SystemRouteRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemRouteRbacReferences(route string) (res *Ref, pp []*Ref, err error) {
	if route != "*" {
		res = &Ref{ResourceType: types.RouteResourceType, Identifiers: MakeIdentifiers(route)}
	}

	return
}

// SystemTemplateRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemTemplateRbacReferences(template string) (res *Ref, pp []*Ref, err error) {
	if template != "*" {
		res = &Ref{ResourceType: types.TemplateResourceType, Identifiers: MakeIdentifiers(template)}
	}

	return
}

// SystemUserRbacReferences generates RBAC references
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func SystemUserRbacReferences(user string) (res *Ref, pp []*Ref, err error) {
	if user != "*" {
		res = &Ref{ResourceType: types.UserResourceType, Identifiers: MakeIdentifiers(user)}
	}

	return
}
