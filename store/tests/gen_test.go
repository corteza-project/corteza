package tests

// This file is auto-generated.
//
// Template:	pkg/codegen/assets/store_test_all.gen.go.tpl
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

import (
	"github.com/cortezaproject/corteza-server/store"
	"testing"
)

func testAllGenerated(t *testing.T, s store.Storer) {
	// Run generated tests for Actionlog
	t.Run("Actionlog", func(t *testing.T) {
		testActionlog(t, s)
	})

	// Run generated tests for ApigwFunction
	t.Run("ApigwFunction", func(t *testing.T) {
		testApigwFunction(t, s)
	})

	// Run generated tests for ApigwRoute
	t.Run("ApigwRoute", func(t *testing.T) {
		testApigwRoute(t, s)
	})

	// Run generated tests for Applications
	t.Run("Applications", func(t *testing.T) {
		testApplications(t, s)
	})

	// Run generated tests for Attachment
	t.Run("Attachment", func(t *testing.T) {
		testAttachment(t, s)
	})

	// Run generated tests for AuthClients
	t.Run("AuthClients", func(t *testing.T) {
		testAuthClients(t, s)
	})

	// Run generated tests for AuthConfirmedClients
	t.Run("AuthConfirmedClients", func(t *testing.T) {
		testAuthConfirmedClients(t, s)
	})

	// Run generated tests for AuthOa2tokens
	t.Run("AuthOa2tokens", func(t *testing.T) {
		testAuthOa2tokens(t, s)
	})

	// Run generated tests for AuthSessions
	t.Run("AuthSessions", func(t *testing.T) {
		testAuthSessions(t, s)
	})

	// Run generated tests for AutomationSessions
	t.Run("AutomationSessions", func(t *testing.T) {
		testAutomationSessions(t, s)
	})

	// Run generated tests for AutomationTriggers
	t.Run("AutomationTriggers", func(t *testing.T) {
		testAutomationTriggers(t, s)
	})

	// Run generated tests for AutomationWorkflows
	t.Run("AutomationWorkflows", func(t *testing.T) {
		testAutomationWorkflows(t, s)
	})

	// Run generated tests for ComposeAttachments
	t.Run("ComposeAttachments", func(t *testing.T) {
		testComposeAttachments(t, s)
	})

	// Run generated tests for ComposeCharts
	t.Run("ComposeCharts", func(t *testing.T) {
		testComposeCharts(t, s)
	})

	// Run generated tests for ComposeModuleFields
	t.Run("ComposeModuleFields", func(t *testing.T) {
		testComposeModuleFields(t, s)
	})

	// Run generated tests for ComposeModules
	t.Run("ComposeModules", func(t *testing.T) {
		testComposeModules(t, s)
	})

	// Run generated tests for ComposeNamespaces
	t.Run("ComposeNamespaces", func(t *testing.T) {
		testComposeNamespaces(t, s)
	})

	// Run generated tests for ComposePages
	t.Run("ComposePages", func(t *testing.T) {
		testComposePages(t, s)
	})

	// Run generated tests for ComposeRecordValues
	t.Run("ComposeRecordValues", func(t *testing.T) {
		testComposeRecordValues(t, s)
	})

	// Run generated tests for ComposeRecords
	t.Run("ComposeRecords", func(t *testing.T) {
		testComposeRecords(t, s)
	})

	// Run generated tests for Credentials
	t.Run("Credentials", func(t *testing.T) {
		testCredentials(t, s)
	})

	// Run generated tests for FederationExposedModules
	t.Run("FederationExposedModules", func(t *testing.T) {
		testFederationExposedModules(t, s)
	})

	// Run generated tests for FederationModuleMappings
	t.Run("FederationModuleMappings", func(t *testing.T) {
		testFederationModuleMappings(t, s)
	})

	// Run generated tests for FederationNodes
	t.Run("FederationNodes", func(t *testing.T) {
		testFederationNodes(t, s)
	})

	// Run generated tests for FederationNodesSync
	t.Run("FederationNodesSync", func(t *testing.T) {
		testFederationNodesSync(t, s)
	})

	// Run generated tests for FederationSharedModules
	t.Run("FederationSharedModules", func(t *testing.T) {
		testFederationSharedModules(t, s)
	})

	// Run generated tests for Flags
	t.Run("Flags", func(t *testing.T) {
		testFlags(t, s)
	})

	// Run generated tests for Labels
	t.Run("Labels", func(t *testing.T) {
		testLabels(t, s)
	})

	// Run generated tests for MessagebusQueueMessage
	t.Run("MessagebusQueueMessage", func(t *testing.T) {
		testMessagebusQueueMessage(t, s)
	})

	// Run generated tests for MessagebusQueueSettings
	t.Run("MessagebusQueueSettings", func(t *testing.T) {
		testMessagebusQueueSettings(t, s)
	})

	// Run generated tests for RbacRules
	t.Run("RbacRules", func(t *testing.T) {
		testRbacRules(t, s)
	})

	// Run generated tests for Reminders
	t.Run("Reminders", func(t *testing.T) {
		testReminders(t, s)
	})

	// Run generated tests for RoleMembers
	t.Run("RoleMembers", func(t *testing.T) {
		testRoleMembers(t, s)
	})

	// Run generated tests for Roles
	t.Run("Roles", func(t *testing.T) {
		testRoles(t, s)
	})

	// Run generated tests for Settings
	t.Run("Settings", func(t *testing.T) {
		testSettings(t, s)
	})

	// Run generated tests for Templates
	t.Run("Templates", func(t *testing.T) {
		testTemplates(t, s)
	})

	// Run generated tests for Users
	t.Run("Users", func(t *testing.T) {
		testUsers(t, s)
	})
}
