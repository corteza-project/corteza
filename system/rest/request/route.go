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
	RouteList struct {
		// RouteID GET parameter
		//
		// Filter by route ID
		RouteID []string

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

	RouteCreate struct {
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

	RouteUpdate struct {
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

	RouteRead struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}

	RouteDelete struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}

	RouteUndelete struct {
		// RouteID PATH parameter
		//
		// Route ID
		RouteID uint64 `json:",string"`
	}
)

// NewRouteList request
func NewRouteList() *RouteList {
	return &RouteList{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) Auditable() map[string]interface{} {
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
func (r RouteList) GetRouteID() []string {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetQuery() string {
	return r.Query
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetDeleted() uint64 {
	return r.Deleted
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetDisabled() uint64 {
	return r.Disabled
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetLabels() map[string]string {
	return r.Labels
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetPageCursor() string {
	return r.PageCursor
}

// Auditable returns all auditable/loggable parameters
func (r RouteList) GetSort() string {
	return r.Sort
}

// Fill processes request and fills internal variables
func (r *RouteList) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["routeID[]"]; ok {
			r.RouteID, err = val, nil
			if err != nil {
				return err
			}
		} else if val, ok := tmp["routeID"]; ok {
			r.RouteID, err = val, nil
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

// NewRouteCreate request
func NewRouteCreate() *RouteCreate {
	return &RouteCreate{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"endpoint": r.Endpoint,
		"method":   r.Method,
		"debug":    r.Debug,
		"enabled":  r.Enabled,
		"group":    r.Group,
	}
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) GetEndpoint() string {
	return r.Endpoint
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) GetMethod() string {
	return r.Method
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) GetDebug() bool {
	return r.Debug
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) GetEnabled() bool {
	return r.Enabled
}

// Auditable returns all auditable/loggable parameters
func (r RouteCreate) GetGroup() uint64 {
	return r.Group
}

// Fill processes request and fills internal variables
func (r *RouteCreate) Fill(req *http.Request) (err error) {

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

// NewRouteUpdate request
func NewRouteUpdate() *RouteUpdate {
	return &RouteUpdate{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) Auditable() map[string]interface{} {
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
func (r RouteUpdate) GetRouteID() uint64 {
	return r.RouteID
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) GetEndpoint() string {
	return r.Endpoint
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) GetMethod() string {
	return r.Method
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) GetDebug() bool {
	return r.Debug
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) GetEnabled() bool {
	return r.Enabled
}

// Auditable returns all auditable/loggable parameters
func (r RouteUpdate) GetGroup() uint64 {
	return r.Group
}

// Fill processes request and fills internal variables
func (r *RouteUpdate) Fill(req *http.Request) (err error) {

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

// NewRouteRead request
func NewRouteRead() *RouteRead {
	return &RouteRead{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r RouteRead) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *RouteRead) Fill(req *http.Request) (err error) {

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

// NewRouteDelete request
func NewRouteDelete() *RouteDelete {
	return &RouteDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r RouteDelete) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *RouteDelete) Fill(req *http.Request) (err error) {

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

// NewRouteUndelete request
func NewRouteUndelete() *RouteUndelete {
	return &RouteUndelete{}
}

// Auditable returns all auditable/loggable parameters
func (r RouteUndelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"routeID": r.RouteID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r RouteUndelete) GetRouteID() uint64 {
	return r.RouteID
}

// Fill processes request and fills internal variables
func (r *RouteUndelete) Fill(req *http.Request) (err error) {

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
