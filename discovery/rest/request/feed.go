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
	FeedChanges struct {
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

// NewFeedChanges request
func NewFeedChanges() *FeedChanges {
	return &FeedChanges{}
}

// Auditable returns all auditable/loggable parameters
func (r FeedChanges) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"limit":      r.Limit,
		"pageCursor": r.PageCursor,
	}
}

// Auditable returns all auditable/loggable parameters
func (r FeedChanges) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r FeedChanges) GetPageCursor() string {
	return r.PageCursor
}

// Fill processes request and fills internal variables
func (r *FeedChanges) Fill(req *http.Request) (err error) {

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
