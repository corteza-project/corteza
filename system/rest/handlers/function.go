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
	FunctionAPI interface {
		List(context.Context, *request.FunctionList) (interface{}, error)
		Create(context.Context, *request.FunctionCreate) (interface{}, error)
		Update(context.Context, *request.FunctionUpdate) (interface{}, error)
		Read(context.Context, *request.FunctionRead) (interface{}, error)
		Delete(context.Context, *request.FunctionDelete) (interface{}, error)
		Undelete(context.Context, *request.FunctionUndelete) (interface{}, error)
		Definitions(context.Context, *request.FunctionDefinitions) (interface{}, error)
	}

	// HTTP API interface
	Function struct {
		List        func(http.ResponseWriter, *http.Request)
		Create      func(http.ResponseWriter, *http.Request)
		Update      func(http.ResponseWriter, *http.Request)
		Read        func(http.ResponseWriter, *http.Request)
		Delete      func(http.ResponseWriter, *http.Request)
		Undelete    func(http.ResponseWriter, *http.Request)
		Definitions func(http.ResponseWriter, *http.Request)
	}
)

func NewFunction(h FunctionAPI) *Function {
	return &Function{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewFunctionList()
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
			params := request.NewFunctionCreate()
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
			params := request.NewFunctionUpdate()
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
			params := request.NewFunctionRead()
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
			params := request.NewFunctionDelete()
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
			params := request.NewFunctionUndelete()
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
		Definitions: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewFunctionDefinitions()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Definitions(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Function) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/apigw/function/", h.List)
		r.Put("/apigw/function", h.Create)
		r.Post("/apigw/function/{functionID}", h.Update)
		r.Get("/apigw/function/{functionID}", h.Read)
		r.Delete("/apigw/function/{functionID}", h.Delete)
		r.Post("/apigw/function/{functionID}/undelete", h.Undelete)
		r.Get("/apigw/function/def", h.Definitions)
	})
}
