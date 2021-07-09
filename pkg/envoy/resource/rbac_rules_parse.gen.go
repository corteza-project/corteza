package resource

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

// Definitions file that controls how this file is generated:
// - automation.workflow.yaml
// - automation.yaml
// - compose.chart.yaml
// - compose.module-field.yaml
// - compose.module.yaml
// - compose.namespace.yaml
// - compose.page.yaml
// - compose.record.yaml
// - compose.yaml
// - system.application.yaml
// - system.auth-client.yaml
// - system.role.yaml
// - system.route.yaml
// - system.template.yaml
// - system.user.yaml
// - system.yaml

import (
	"fmt"
	automationTypes "github.com/cortezaproject/corteza-server/automation/types"
	composeTypes "github.com/cortezaproject/corteza-server/compose/types"
	systemTypes "github.com/cortezaproject/corteza-server/system/types"
	"strings"
)

// Parse generates resource setting logic for each resource
//
// Resources with "envoy: false" are skipped
//
// This function is auto-generated
func ParseRule(res string) (string, *Ref, []*Ref, error) {
	if res == "" {
		return "", nil, nil, fmt.Errorf("empty resource")
	}

	sp := "/"

	res = strings.TrimSpace(res)
	res = strings.TrimRight(res, sp)
	rr := strings.Split(res, sp)

	// only service defined (corteza::system, corteza::compose, ...)
	if len(rr) == 1 {
		return res, nil, nil, nil
	}

	// full thing
	resourceType, path := rr[0], rr[1:]

	for p := 1; p < len(path); p++ {
		if path[p] != "*" && path[p-1] == "*" {
			return "", nil, nil, fmt.Errorf("invalid path wildcard combination for '%s'", res)
		}
	}

	// make the resource provide the slice of parent resources we should nest under
	switch resourceType {
	case automationTypes.WorkflowResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := AutomationWorkflowRbacReferences(
			// workflow
			path[0],
		)
		return automationTypes.WorkflowResourceType, ref, pp, err

	case automationTypes.ComponentResourceType:
		if len(path) != 0 {
			return "", nil, nil, fmt.Errorf("expecting 0 reference components in path, got %d", len(path))
		}

		// Component resource, no path
		return automationTypes.ComponentResourceType, nil, nil, nil
	case composeTypes.ChartResourceType:
		if len(path) != 2 {
			return "", nil, nil, fmt.Errorf("expecting 2 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposeChartRbacReferences(
			// namespace
			path[0],

			// chart
			path[1],
		)
		return composeTypes.ChartResourceType, ref, pp, err

	case composeTypes.ModuleFieldResourceType:
		if len(path) != 3 {
			return "", nil, nil, fmt.Errorf("expecting 3 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposeModuleFieldRbacReferences(
			// namespace
			path[0],

			// module
			path[1],

			// moduleField
			path[2],
		)
		return composeTypes.ModuleFieldResourceType, ref, pp, err

	case composeTypes.ModuleResourceType:
		if len(path) != 2 {
			return "", nil, nil, fmt.Errorf("expecting 2 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposeModuleRbacReferences(
			// namespace
			path[0],

			// module
			path[1],
		)
		return composeTypes.ModuleResourceType, ref, pp, err

	case composeTypes.NamespaceResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposeNamespaceRbacReferences(
			// namespace
			path[0],
		)
		return composeTypes.NamespaceResourceType, ref, pp, err

	case composeTypes.PageResourceType:
		if len(path) != 2 {
			return "", nil, nil, fmt.Errorf("expecting 2 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposePageRbacReferences(
			// namespace
			path[0],

			// page
			path[1],
		)
		return composeTypes.PageResourceType, ref, pp, err

	case composeTypes.RecordResourceType:
		if len(path) != 3 {
			return "", nil, nil, fmt.Errorf("expecting 3 reference components in path, got %d", len(path))
		}
		ref, pp, err := ComposeRecordRbacReferences(
			// namespace
			path[0],

			// module
			path[1],

			// record
			path[2],
		)
		return composeTypes.RecordResourceType, ref, pp, err

	case composeTypes.ComponentResourceType:
		if len(path) != 0 {
			return "", nil, nil, fmt.Errorf("expecting 0 reference components in path, got %d", len(path))
		}

		// Component resource, no path
		return composeTypes.ComponentResourceType, nil, nil, nil
	case systemTypes.ApplicationResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemApplicationRbacReferences(
			// application
			path[0],
		)
		return systemTypes.ApplicationResourceType, ref, pp, err

	case systemTypes.AuthClientResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemAuthClientRbacReferences(
			// authClient
			path[0],
		)
		return systemTypes.AuthClientResourceType, ref, pp, err

	case systemTypes.RoleResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemRoleRbacReferences(
			// role
			path[0],
		)
		return systemTypes.RoleResourceType, ref, pp, err

	case systemTypes.RouteResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemRouteRbacReferences(
			// route
			path[0],
		)
		return systemTypes.RouteResourceType, ref, pp, err

	case systemTypes.TemplateResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemTemplateRbacReferences(
			// template
			path[0],
		)
		return systemTypes.TemplateResourceType, ref, pp, err

	case systemTypes.UserResourceType:
		if len(path) != 1 {
			return "", nil, nil, fmt.Errorf("expecting 1 reference components in path, got %d", len(path))
		}
		ref, pp, err := SystemUserRbacReferences(
			// user
			path[0],
		)
		return systemTypes.UserResourceType, ref, pp, err

	case systemTypes.ComponentResourceType:
		if len(path) != 0 {
			return "", nil, nil, fmt.Errorf("expecting 0 reference components in path, got %d", len(path))
		}

		// Component resource, no path
		return systemTypes.ComponentResourceType, nil, nil, nil
	}

	// return unhandled resource as-is
	return resourceType, nil, nil, nil
}
