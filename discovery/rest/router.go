package rest

import (
	"github.com/cortezaproject/corteza-server/discovery/rest/handlers"
	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/go-chi/chi"
)

func MountRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(auth.AccessTokenCheck("discovery"))

		handlers.NewResources(Resources()).MountRoutes(r)
		handlers.NewFeed(Feed()).MountRoutes(r)
		handlers.NewMappings(Mappings()).MountRoutes(r)
	})
}
