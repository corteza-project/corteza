package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cortezaproject/corteza-server/pkg/filter"
)

type (
	ApigwFunctionKind string

	FuncParams map[string]interface{}

	Function struct {
		ID     uint64     `json:"functionID,string"`
		Route  uint64     `json:"routeID,string"`
		Weight uint64     `json:"weight"`
		Ref    string     `json:"ref,omitempty"`
		Kind   string     `json:"kind,omitempty"`
		Params FuncParams `json:"params"`

		CreatedAt time.Time  `json:"createdAt,omitempty"`
		CreatedBy uint64     `json:"createdBy,string" `
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
		UpdatedBy uint64     `json:"updatedBy,string,omitempty" `
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
		DeletedBy uint64     `json:"deletedBy,string,omitempty" `
	}

	FunctionFilter struct {
		RouteID  uint64 `json:"routeID,string"`
		Endpoint string `json:"endpoint"`
		Group    string `json:"group"`
		Enabled  bool   `json:"enabled"`

		Deleted filter.State `json:"deleted"`

		// Check fn is called by store backend for each resource found function can
		// modify the resource and return false if store should not return it
		//
		// Store then loads additional resources to satisfy the paging parameters
		Check func(*Route) (bool, error) `json:"-"`

		filter.Sorting
		filter.Paging
	}
)

const (
	ApigwFunctionKindVerifier  ApigwFunctionKind = "functionVerifier"
	ApigwFunctionKindValidator ApigwFunctionKind = "functionValidator"
	ApigwFunctionKindMatcher   ApigwFunctionKind = "functionMatcher"
	ApigwFunctionKindProcesser ApigwFunctionKind = "functionProcesser"
	ApigwFunctionKindExpediter ApigwFunctionKind = "functionExpediter"
)

func (vv *FuncParams) Scan(value interface{}) (err error) {
	if err := json.Unmarshal(value.([]byte), vv); err != nil {
		return fmt.Errorf("cannot scan '%v' into FuncParams", value)
	}

	return
}

func (vv FuncParams) Value() (driver.Value, error) {
	return json.Marshal(vv)
}
