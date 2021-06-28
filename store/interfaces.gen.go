package store

// This file is auto-generated.
//
// Template:	pkg/codegen/assets/store_interfaces_joined.gen.go.tpl
// Definitions:
//  - store/actionlog.yaml
//  - store/apigw_function.yaml
//  - store/apigw_route.yaml
//  - store/applications.yaml
//  - store/attachments.yaml
//  - store/auth_clients.yaml
//  - store/auth_confirmed_clients.yaml
//  - store/auth_oa2tokens.yaml
//  - store/auth_sessions.yaml
//  - store/automation_sessions.yaml
//  - store/automation_triggers.yaml
//  - store/automation_workflows.yaml
//  - store/compose_attachments.yaml
//  - store/compose_charts.yaml
//  - store/compose_module_fields.yaml
//  - store/compose_modules.yaml
//  - store/compose_namespaces.yaml
//  - store/compose_pages.yaml
//  - store/compose_record_values.yaml
//  - store/compose_records.yaml
//  - store/credentials.yaml
//  - store/federation_exposed_modules.yaml
//  - store/federation_module_mappings.yaml
//  - store/federation_nodes.yaml
//  - store/federation_nodes_sync.yaml
//  - store/federation_shared_modules.yaml
//  - store/flags.yaml
//  - store/labels.yaml
//  - store/messagebus_queue_message.yaml
//  - store/messagebus_queue_settings.yaml
//  - store/rbac_rules.yaml
//  - store/reminders.yaml
//  - store/role_members.yaml
//  - store/roles.yaml
//  - store/settings.yaml
//  - store/templates.yaml
//  - store/users.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

type (
	// Sortable interface combines interfaces of all supported store interfaces
	storerGenerated interface {
		Actionlogs
		ApigwFunctions
		ApigwRoutes
		Applications
		Attachments
		AuthClients
		AuthConfirmedClients
		AuthOa2tokens
		AuthSessions
		AutomationSessions
		AutomationTriggers
		AutomationWorkflows
		ComposeAttachments
		ComposeCharts
		ComposeModuleFields
		ComposeModules
		ComposeNamespaces
		ComposePages
		ComposeRecordValues
		ComposeRecords
		Credentials
		FederationExposedModules
		FederationModuleMappings
		FederationNodes
		FederationNodesSyncs
		FederationSharedModules
		Flags
		Labels
		MessagebusQueueMessages
		MessagebusQueueSettings
		RbacRules
		Reminders
		RoleMembers
		Roles
		Settings
		Templates
		Users
	}
)
