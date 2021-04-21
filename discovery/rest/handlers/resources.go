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
	ResourcesAPI interface {
		SystemUsers(context.Context, *request.ResourcesSystemUsers) (interface{}, error)
		ComposeNamespaces(context.Context, *request.ResourcesComposeNamespaces) (interface{}, error)
		ComposeModules(context.Context, *request.ResourcesComposeModules) (interface{}, error)
		ComposeRecords(context.Context, *request.ResourcesComposeRecords) (interface{}, error)
	}

	// HTTP API interface
	Resources struct {
		SystemUsers       func(http.ResponseWriter, *http.Request)
		ComposeNamespaces func(http.ResponseWriter, *http.Request)
		ComposeModules    func(http.ResponseWriter, *http.Request)
		ComposeRecords    func(http.ResponseWriter, *http.Request)
	}
)

func NewResources(h ResourcesAPI) *Resources {
	return &Resources{
		SystemUsers: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewResourcesSystemUsers()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.SystemUsers(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		ComposeNamespaces: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewResourcesComposeNamespaces()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.ComposeNamespaces(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		ComposeModules: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewResourcesComposeModules()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.ComposeModules(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		ComposeRecords: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewResourcesComposeRecords()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.ComposeRecords(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Resources) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/resources/system/users", h.SystemUsers)
		r.Get("/resources/compose/namespaces", h.ComposeNamespaces)
		r.Get("/resources/compose/namespaces/{namespaceID}/modules", h.ComposeModules)
		r.Get("/resources/compose/namespaces/{namespaceID}/modules/{moduleID}/records", h.ComposeRecords)
	})
}
