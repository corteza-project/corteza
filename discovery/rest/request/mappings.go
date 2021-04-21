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
	MappingsList struct {
	}
)

// NewMappingsList request
func NewMappingsList() *MappingsList {
	return &MappingsList{}
}

// Auditable returns all auditable/loggable parameters
func (r MappingsList) Auditable() map[string]interface{} {
	return map[string]interface{}{}
}

// Fill processes request and fills internal variables
func (r *MappingsList) Fill(req *http.Request) (err error) {

	return err
}
