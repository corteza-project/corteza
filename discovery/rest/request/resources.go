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
	ResourcesSystemUsers struct {
		// Limit GET parameter
		//
		// Limit
		Limit uint

		// PageCursor GET parameter
		//
		// Page cursor
		PageCursor string
	}

	ResourcesComposeNamespaces struct {
		// Limit GET parameter
		//
		// Limit
		Limit uint

		// PageCursor GET parameter
		//
		// Page cursor
		PageCursor string
	}

	ResourcesComposeModules struct {
		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// Limit GET parameter
		//
		// Limit
		Limit uint

		// PageCursor GET parameter
		//
		// Page cursor
		PageCursor string
	}

	ResourcesComposeRecords struct {
		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// ModuleID PATH parameter
		//
		// Module ID
		ModuleID uint64 `json:",string"`

		// Limit GET parameter
		//
		// Limit
		Limit uint

		// PageCursor GET parameter
		//
		// Page cursor
		PageCursor string
	}
)

// NewResourcesSystemUsers request
func NewResourcesSystemUsers() *ResourcesSystemUsers {
	return &ResourcesSystemUsers{}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesSystemUsers) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"limit":      r.Limit,
		"pageCursor": r.PageCursor,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesSystemUsers) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesSystemUsers) GetPageCursor() string {
	return r.PageCursor
}

// Fill processes request and fills internal variables
func (r *ResourcesSystemUsers) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

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
	}

	return err
}

// NewResourcesComposeNamespaces request
func NewResourcesComposeNamespaces() *ResourcesComposeNamespaces {
	return &ResourcesComposeNamespaces{}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeNamespaces) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"limit":      r.Limit,
		"pageCursor": r.PageCursor,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeNamespaces) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeNamespaces) GetPageCursor() string {
	return r.PageCursor
}

// Fill processes request and fills internal variables
func (r *ResourcesComposeNamespaces) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

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
	}

	return err
}

// NewResourcesComposeModules request
func NewResourcesComposeModules() *ResourcesComposeModules {
	return &ResourcesComposeModules{}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeModules) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"namespaceID": r.NamespaceID,
		"limit":       r.Limit,
		"pageCursor":  r.PageCursor,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeModules) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeModules) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeModules) GetPageCursor() string {
	return r.PageCursor
}

// Fill processes request and fills internal variables
func (r *ResourcesComposeModules) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

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
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewResourcesComposeRecords request
func NewResourcesComposeRecords() *ResourcesComposeRecords {
	return &ResourcesComposeRecords{}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeRecords) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"namespaceID": r.NamespaceID,
		"moduleID":    r.ModuleID,
		"limit":       r.Limit,
		"pageCursor":  r.PageCursor,
	}
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeRecords) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeRecords) GetModuleID() uint64 {
	return r.ModuleID
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeRecords) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r ResourcesComposeRecords) GetPageCursor() string {
	return r.PageCursor
}

// Fill processes request and fills internal variables
func (r *ResourcesComposeRecords) Fill(req *http.Request) (err error) {

	{
		// GET params
		tmp := req.URL.Query()

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
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "moduleID")
		r.ModuleID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}
