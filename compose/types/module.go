package types

import (
	"database/sql/driver"
	"encoding/json"
	discovery "github.com/cortezaproject/corteza-server/discovery/types"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
	"time"
)

type (
	Module struct {
		ID     uint64         `json:"moduleID,string"`
		Handle string         `json:"handle"`
		Name   string         `json:"name"`
		Meta   types.JSONText `json:"meta"`
		Fields ModuleFieldSet `json:"fields"`

		Labels map[string]string `json:"labels,omitempty"`

		NamespaceID uint64 `json:"namespaceID,string"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
	}

	ModuleMeta struct {
		Discovery struct {
			// Index access restrictions for module and it's sub-resources
			IndexAccessRestrictions struct {
				// Can modules in this namespace be indexed
				Module discovery.Access `json:"module"`

				// Can records in this namespace be indexed
				Records discovery.Access `json:"records"`
			} `json:"indexAccessRestrictions"`
		}
	}

	ModuleFilter struct {
		ModuleID    []uint64 `json:"moduleID"`
		NamespaceID uint64   `json:"namespaceID,string"`
		Query       string   `json:"query"`
		Handle      string   `json:"handle"`
		Name        string   `json:"name"`

		LabeledIDs []uint64          `json:"-"`
		Labels     map[string]string `json:"labels,omitempty"`

		Deleted filter.State `json:"deleted"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*Module) (bool, error) `json:"-"`

		// Standard helpers for paging and sorting
		filter.Sorting
		filter.Paging
	}
)

// Resource returns a system resource ID for this type
func (m Module) RBACResource() rbac.Resource {
	return ModuleRBACResource.AppendID(m.ID)
}

func (m Module) DynamicRoles(userID uint64) []uint64 {
	return nil
}

func (m Module) Clone() *Module {
	c := &m
	c.Fields = m.Fields.Clone()
	return c
}

// FindByHandle finds module by it's handle
func (set ModuleSet) FindByHandle(handle string) *Module {
	for i := range set {
		if set[i].Handle == handle {
			return set[i]
		}
	}

	return nil
}

func (nm *ModuleMeta) Scan(value interface{}) error {
	//lint:ignore S1034 This typecast is intentional, we need to get []byte out of a []uint8
	switch value.(type) {
	case nil:
		*nm = ModuleMeta{}
	case []uint8:
		b := value.([]byte)
		if err := json.Unmarshal(b, nm); err != nil {
			return errors.Wrapf(err, "cannot scan '%v' into ModuleMeta", string(b))
		}
	}

	return nil
}

func (nm ModuleMeta) Value() (driver.Value, error) {
	return json.Marshal(nm)
}
