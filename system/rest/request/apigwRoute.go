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
	"github.com/cortezaproject/corteza-server/pkg/label"
	"github.com/cortezaproject/corteza-server/pkg/payload"
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
	ApigwRouteList struct {
		// RouteID GET parameter
		//
		// Filter by route ID
		RouteID []uint64

		// Query GET parameter
		//
		// Filter routes
		Query string

		// Deleted GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) deleted routes
		Deleted uint64 `json:",string"`

		// Disabled GET parameter
		//
		// Exclude (0, default), include (1) or return only (2) disabled routes
		Disabled uint64 `json:",string"`

		// Labels GET parameter
		//
		// Labels
		Labels map[string]string

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

	ApigwRouteCreate struct {
		// Endpoint POST parameter
		//
		// Route endpoint
		Endpoint string

		// Method POST parameter
		//
		// Route method
		Method string

		// Debug POST parameter
		//
		// Debug route
		Debug bool

		// Enabled POST parameter
		//
		// Is route enabled
		Enabled bool

		// Group POST parameter
		//
		// Route group
		Group uint64 `json:",string"`
	}

	ApigwRouteUpdate struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`

		// Endpoint POST parameter
		//
		// Route endpoint
		Endpoint string

		// Method POST parameter
		//
		// Route method
		Method string

		// Debug POST parameter
		//
		// Debug route
		Debug bool

		// Enabled POST parameter
		//
		// Is route enabled
		Enabled bool

		// Group POST parameter
		//
		// Route group
		Group uint64 `json:",string"`
	}

	ApigwRouteRead struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}

	ApigwRouteDelete struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}

	ApigwRouteUndelete struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}
)

// NewApigwRouteList request
func NewApigwRouteList() *ApigwRouteList {
	return &ApigwRouteList{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID":    r.RouteID,
		"query":      r.Query,
		"deleted":    r.Deleted,
		"disabled":   r.Disabled,
		"labels":     r.Labels,
		"limit":      r.Limit,
		"pageCursor": r.PageCursor,
		"sort":       r.Sort,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetRouteID() []uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetQuery() string {
	return r.Query
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetDeleted() uint64 {
	return r.Deleted
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetDisabled() uint64 {
	return r.Disabled
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetLabels() map[string]string {
	return r.Labels
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetPageCursor() string {
	return r.PageCursor
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteList) GetSort() string {
	return r.Sort
}

// Fill processes request and fills internal variables
func (r *ApigwRouteList) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["routeID[]"]; ok {
			r.RouteID, err = payload.ParseUint64s(val), nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["routeID"]; ok {
			r.RouteID, err = payload.ParseUint64s(val), nil
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
		if val, ok := tmp["labels[]"]; ok {
			r.Labels, err = label.ParseStrings(val)
			if err != nil {
				return err
			}
		} else if val, ok := tmp["labels"]; ok {
			r.Labels, err = label.ParseStrings(val)
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

// NewApigwRouteCreate request
func NewApigwRouteCreate() *ApigwRouteCreate {
	return &ApigwRouteCreate{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"endpoint": r.Endpoint,
		"method":   r.Method,
		"debug":    r.Debug,
		"enabled":  r.Enabled,
		"group":    r.Group,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) GetEndpoint() string {
	return r.Endpoint
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) GetMethod() string {
	return r.Method
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) GetDebug() bool {
	return r.Debug
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) GetEnabled() bool {
	return r.Enabled
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteCreate) GetGroup() uint64 {
	return r.Group
}

// Fill processes request and fills internal variables
func (r *ApigwRouteCreate) Fill(req *http.Request) (err error) {

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

		if val, ok := req.Form["endpoint"]; ok && len(val) > 0 {
			r.Endpoint, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["method"]; ok && len(val) > 0 {
			r.Method, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["debug"]; ok && len(val) > 0 {
			r.Debug, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["enabled"]; ok && len(val) > 0 {
			r.Enabled, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["group"]; ok && len(val) > 0 {
			r.Group, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewApigwRouteUpdate request
func NewApigwRouteUpdate() *ApigwRouteUpdate {
	return &ApigwRouteUpdate{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID":  r.RouteID,
		"endpoint": r.Endpoint,
		"method":   r.Method,
		"debug":    r.Debug,
		"enabled":  r.Enabled,
		"group":    r.Group,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetEndpoint() string {
	return r.Endpoint
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetMethod() string {
	return r.Method
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetDebug() bool {
	return r.Debug
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetEnabled() bool {
	return r.Enabled
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUpdate) GetGroup() uint64 {
	return r.Group
}

// Fill processes request and fills internal variables
func (r *ApigwRouteUpdate) Fill(req *http.Request) (err error) {

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

		if val, ok := req.Form["endpoint"]; ok && len(val) > 0 {
			r.Endpoint, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["method"]; ok && len(val) > 0 {
			r.Method, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["debug"]; ok && len(val) > 0 {
			r.Debug, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["enabled"]; ok && len(val) > 0 {
			r.Enabled, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["group"]; ok && len(val) > 0 {
			r.Group, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "routeID")
		r.RouteID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwRouteRead request
func NewApigwRouteRead() *ApigwRouteRead {
	return &ApigwRouteRead{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteRead) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *ApigwRouteRead) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "routeID")
		r.RouteID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwRouteDelete request
func NewApigwRouteDelete() *ApigwRouteDelete {
	return &ApigwRouteDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteDelete) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *ApigwRouteDelete) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "routeID")
		r.RouteID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewApigwRouteUndelete request
func NewApigwRouteUndelete() *ApigwRouteUndelete {
	return &ApigwRouteUndelete{}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUndelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ApigwRouteUndelete) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *ApigwRouteUndelete) Fill(req *http.Request) (err error) {

	{
		var val string
		// path params

		val = chi.URLParam(req, "routeID")
		r.RouteID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}
