package types

import (
	"time"

	"github.com/cortezaproject/corteza-server/pkg/filter"
)

type (
	ApigwRoute struct {
		ID       uint64 `json:"routeID,string"`
		Endpoint string `json:"endpoint"`
		Method   string `json:"method"`
		Debug    bool   `json:"debug"`
		Enabled  bool   `json:"enabled"`
		Group    uint64 `json:"group,string"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		CreatedBy uint64     `json:"createdBy,string" `
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		UpdatedBy uint64     `json:"updatedBy,string,omitempty" `
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
		DeletedBy uint64     `json:"deletedBy,string,omitempty" `
	}

	ApigwRouteFilter struct {
		Route   string `json:"route"`
		Group   string `json:"group"`
		Enabled bool   `json:"enabled"`

		Deleted filter.State `json:"deleted"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*ApigwRoute) (bool, error) `json:"-"`

		filter.Sorting
		filter.Paging
	}
)
