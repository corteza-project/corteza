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
	"github.com/cortezaproject/corteza-server/pkg/api"
	"github.com/cortezaproject/corteza-server/system/rest/request"
	"github.com/go-chi/chi"
	"net/http"
)

type (
	// Internal API interface
	RouteAPI interface {
		List(context.Context, *request.RouteList) (interface{}, error)
		Create(context.Context, *request.RouteCreate) (interface{}, error)
		Update(context.Context, *request.RouteUpdate) (interface{}, error)
		Read(context.Context, *request.RouteRead) (interface{}, error)
		Delete(context.Context, *request.RouteDelete) (interface{}, error)
		Undelete(context.Context, *request.RouteUndelete) (interface{}, error)
	}

	// HTTP API interface
	Route struct {
		List     func(http.ResponseWriter, *http.Request)
		Create   func(http.ResponseWriter, *http.Request)
		Update   func(http.ResponseWriter, *http.Request)
		Read     func(http.ResponseWriter, *http.Request)
		Delete   func(http.ResponseWriter, *http.Request)
		Undelete func(http.ResponseWriter, *http.Request)
	}
)

func NewRoute(h RouteAPI) *Route {
	return &Route{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteList()
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
		Create: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteCreate()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Create(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		Update: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteUpdate()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Update(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		Read: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteRead()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Read(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		Delete: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteDelete()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Delete(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
		Undelete: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRouteUndelete()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Undelete(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Route) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/apigw/route/", h.List)
		r.Post("/apigw/route", h.Create)
		r.Put("/apigw/route/{routeID}", h.Update)
		r.Get("/apigw/route/{routeID}", h.Read)
		r.Delete("/apigw/route/{routeID}", h.Delete)
		r.Post("/apigw/route/{routeID}/undelete", h.Undelete)
	})
}
