package importer

import (
	"testing"
)

func TestRoleImport_CastSet(t *testing.T) {
	t.Skip()
	// 	tests := []struct {
	// 		name     string
	// 		resource string
	// 		yaml     string
	// 		set      types.RoleSet
	// 	}{
	// 		{name: "empty", yaml: ``},
	// 		{name: "empty map", yaml: `{}`},
	// 		{name: "empty slice", yaml: `[]`},
	// 		{
	// 			name: "full",
	// 			yaml: `
	// admins: Admins
	// foo:
	//   name: Foo
	// bar:
	// `,
	// 			set: []*types.Role{
	// 				{Handle: "admins", Name: "Admins"},
	// 				{Handle: "foo", Name: "Foo"},
	// 				{Handle: "bar", Name: "bar"},
	// 			},
	// 		},
	// 	}
	// 	for _, tt := range tests {
	// 		t.Run(tt.name, func(t *testing.T) {
	// 			imp := &RoleImport{}
	//
	// 			aux, err := importer.ParseYAML([]byte(tt.yaml))
	// 			require.NoError(t, err)
	// 			require.NoError(t, imp.CastSet(aux))
	// 			require.Equal(t, tt.set, imp.set)
	// 		})
	// 	}
}