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
	FunctionList struct {
		// FunctionID GET parameter
		//
		// Filter by function ID
		FunctionID []string

		// RouteID GET parameter
		//
		// Filter by route ID
		RouteID string

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

	FunctionCreate struct {
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
		Params types.FuncParams
	}

	FunctionUpdate struct {
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
		Params types.FuncParams
	}

	FunctionRead struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}

	FunctionDelete struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}

	FunctionUndelete struct {
		// FunctionID PATH parameter
		//
		// Function ID
		FunctionID uint64 `json:",string"`
	}
)

// NewFunctionList request
func NewFunctionList() *FunctionList {
	return &FunctionList{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
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
func (r FunctionList) GetFunctionID() []string {
	return r.FunctionID
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetRouteID() string {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetQuery() string {
	return r.Query
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetDeleted() uint64 {
	return r.Deleted
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetDisabled() uint64 {
	return r.Disabled
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetPageCursor() string {
	return r.PageCursor
}

// Auditable returns all auditable/loggable parameters
func (r FunctionList) GetSort() string {
	return r.Sort
}

// Fill processes request and fills internal variables
func (r *FunctionList) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["functionID[]"]; ok {
			r.FunctionID, err = val, nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["functionID"]; ok {
			r.FunctionID, err = val, nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["routeID"]; ok && len(val) > 0 {
			r.RouteID, err = val[0], nil
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

// NewFunctionCreate request
func NewFunctionCreate() *FunctionCreate {
	return &FunctionCreate{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
		"weight":  r.Weight,
		"kind":    r.Kind,
		"ref":     r.Ref,
		"params":  r.Params,
	}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) GetWeight() uint64 {
	return r.Weight
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) GetKind() types.ApigwFunctionKind {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) GetRef() string {
	return r.Ref
}

// Auditable returns all auditable/loggable parameters
func (r FunctionCreate) GetParams() types.FuncParams {
	return r.Params
}

// Fill processes request and fills internal variables
func (r *FunctionCreate) Fill(req *http.Request) (err error) {

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

// NewFunctionUpdate request
func NewFunctionUpdate() *FunctionUpdate {
	return &FunctionUpdate{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) Auditable() map[string]interface{} {
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
func (r FunctionUpdate) GetFunctionID() uint64 {
	return r.FunctionID
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) GetWeight() uint64 {
	return r.Weight
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) GetKind() types.ApigwFunctionKind {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) GetRef() string {
	return r.Ref
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUpdate) GetParams() types.FuncParams {
	return r.Params
}

// Fill processes request and fills internal variables
func (r *FunctionUpdate) Fill(req *http.Request) (err error) {

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

// NewFunctionRead request
func NewFunctionRead() *FunctionRead {
	return &FunctionRead{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionRead) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *FunctionRead) Fill(req *http.Request) (err error) {

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

// NewFunctionDelete request
func NewFunctionDelete() *FunctionDelete {
	return &FunctionDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionDelete) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *FunctionDelete) Fill(req *http.Request) (err error) {

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

// NewFunctionUndelete request
func NewFunctionUndelete() *FunctionUndelete {
	return &FunctionUndelete{}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUndelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"functionID": r.FunctionID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r FunctionUndelete) GetFunctionID() uint64 {
	return r.FunctionID
}

// Fill processes request and fills internal variables
func (r *FunctionUndelete) Fill(req *http.Request) (err error) {

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
