package request

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"encoding/json"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/payload"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/go-chi/chi"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// dummy vars to prevent
// unused imports complain
var (
	_ = chi.URLParam
	_ = multipart.ErrMessageTooLarge
	_ = payload.ParseUint64s
	_ = strings.ToLower
	_ = io.EOF
	_ = fmt.Errorf
	_ = json.NewEncoder
)

type (
	// Internal API interface
	ApigwFunctionList struct {
		// RouteID GET parameter
		//
		// Filter by route ID
		RouteID uint64 `json:",string"`

		// Query GET parameter
		//
		// Filter functions
		Query string

		// Deleted GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) deleted functions
		Deleted uint64 `json:",string"`

		// Disabled GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) disabled functions
		Disabled uint64 `json:",string"`

		// Limit GET parameter
		//
		// Limit
		Limit uint

		// PageCursor GET parameter
		//
		// Page cursor
		PageCursor string

		// Sort GET parameter
		//
		// Sort items
		Sort string
	}

	ApigwFunctionCreate struct {
		// RouteID POST parameter
		//
		// Route
		RouteID uint64 `json:",string"`

		// Weight POST parameter
		//
		// Function priority
		Weight uint64 `json:",string"`

		// Kind POST parameter
		//
		// Function kind
		Kind types.ApigwFunctionKind

		// Ref POST parameter
		//
		// Function ref
		Ref string

		// Params POST parameter
		//
		// Function parameters
		Params types.ApigwFuncParams
	}

	ApigwFunctionUpdate struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`

		// RouteID POST parameter
		//
		// Route
		RouteID uint64 `json:",string"`

		// Weight POST parameter
		//
		// Function priority
		Weight uint64 `json:",string"`

		// Kind POST parameter
		//
		// Function kind
		Kind types.ApigwFunctionKind

		// Ref POST parameter
		//
		// Function ref
		Ref string

		// Params POST parameter
		//
		// Function parameters
		Params types.ApigwFuncParams
	}

	ApigwFunctionRead struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}

	ApigwFunctionDelete struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}

	ApigwFunctionUndelete struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}

	ApigwFunctionDefinitions struct {
		// Kind GET parameter
		//
		// Filter functions by kind
		Kind string
	}
)

// NewApigwFunctionList request
func NewApigwFunctionList() *ApigwFunctionList {
	return &ApigwFunctionList{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID":    r.RouteID,
		"query":      r.Query,
		"deleted":    r.Deleted,
		"disabled":   r.Disabled,
		"limit":      r.Limit,
		"pageCursor": r.PageCursor,
		"sort":       r.Sort,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetQuery() string {
	return r.Query
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetDeleted() uint64 {
	return r.Deleted
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetDisabled() uint64 {
	return r.Disabled
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetPageCursor() string {
	return r.PageCursor
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionList) GetSort() string {
	return r.Sort
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionList) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["routeID"]; ok && len(val) > 0 {
			r.RouteID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["query"]; ok && len(val) > 0 {
			r.Query, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["deleted"]; ok && len(val) > 0 {
			r.Deleted, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["disabled"]; ok && len(val) > 0 {
			r.Disabled, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["limit"]; ok && len(val) > 0 {
			r.Limit, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["pageCursor"]; ok && len(val) > 0 {
			r.PageCursor, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["sort"]; ok && len(val) > 0 {
			r.Sort, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewApigwFunctionCreate request
func NewApigwFunctionCreate() *ApigwFunctionCreate {
	return &ApigwFunctionCreate{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
		"weight":  r.Weight,
		"kind":    r.Kind,
		"ref":     r.Ref,
		"params":  r.Params,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) GetWeight() uint64 {
	return r.Weight
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) GetKind() types.ApigwFunctionKind {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) GetRef() string {
	return r.Ref
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionCreate) GetParams() types.ApigwFuncParams {
	return r.Params
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionCreate) Fill(req *http.Request) (err error) {

	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["routeID"]; ok && len(val) > 0 {
			r.RouteID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["weight"]; ok && len(val) > 0 {
			r.Weight, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["kind"]; ok && len(val) > 0 {
			r.Kind, err = types.ApigwFunctionKind(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["ref"]; ok && len(val) > 0 {
			r.Ref, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["params[]"]; ok {
			r.Params, err = types.ParseApigwfFunctionParams(val)
			if err != nil {
				return err
			}
		} else if val, ok := req.Form["params"]; ok {
			r.Params, err = types.ParseApigwfFunctionParams(val)
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewApigwFunctionUpdate request
func NewApigwFunctionUpdate() *ApigwFunctionUpdate {
	return &ApigwFunctionUpdate{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
		"routeID":    r.RouteID,
		"weight":     r.Weight,
		"kind":       r.Kind,
		"ref":        r.Ref,
		"params":     r.Params,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetFunctionID() uint64 {
	return r.FunctionID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetWeight() uint64 {
	return r.Weight
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetKind() types.ApigwFunctionKind {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetRef() string {
	return r.Ref
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUpdate) GetParams() types.ApigwFuncParams {
	return r.Params
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionUpdate) Fill(req *http.Request) (err error) {

	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["routeID"]; ok && len(val) > 0 {
			r.RouteID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["weight"]; ok && len(val) > 0 {
			r.Weight, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["kind"]; ok && len(val) > 0 {
			r.Kind, err = types.ApigwFunctionKind(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["ref"]; ok && len(val) > 0 {
			r.Ref, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["params[]"]; ok {
			r.Params, err = types.ParseApigwfFunctionParams(val)
			if err != nil {
				return err
			}
		} else if val, ok := req.Form["params"]; ok {
			r.Params, err = types.ParseApigwfFunctionParams(val)
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "functionID")
		r.FunctionID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwFunctionRead request
func NewApigwFunctionRead() *ApigwFunctionRead {
	return &ApigwFunctionRead{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionRead) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionRead) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "functionID")
		r.FunctionID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwFunctionDelete request
func NewApigwFunctionDelete() *ApigwFunctionDelete {
	return &ApigwFunctionDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionDelete) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionDelete) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "functionID")
		r.FunctionID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwFunctionUndelete request
func NewApigwFunctionUndelete() *ApigwFunctionUndelete {
	return &ApigwFunctionUndelete{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUndelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionUndelete) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionUndelete) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "functionID")
		r.FunctionID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwFunctionDefinitions request
func NewApigwFunctionDefinitions() *ApigwFunctionDefinitions {
	return &ApigwFunctionDefinitions{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionDefinitions) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind": r.Kind,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwFunctionDefinitions) GetKind() string {
	return r.Kind
}

// Fill processes request and fills internal variables
func (r *ApigwFunctionDefinitions) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["kind"]; ok && len(val) > 0 {
			r.Kind, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}
