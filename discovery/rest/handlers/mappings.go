package handlers

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"context"
	"github.com/cortezaproject/corteza-server/discovery/rest/request"
	"github.com/cortezaproject/corteza-server/pkg/api"
	"github.com/go-chi/chi"
	"net/http"
)

type (
	// Internal API interface
	MappingsAPI interface {
		List(context.Context, *request.MappingsList) (interface{}, error)
	}

	// HTTP API interface
	Mappings struct {
		List func(http.ResponseWriter, *http.Request)
	}
)

func NewMappings(h MappingsAPI) *Mappings {
	return &Mappings{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewMappingsList()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.List(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Mappings) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/mappings/", h.List)
	})
}
