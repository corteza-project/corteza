package rest

import (
	"github.com/go-chi/chi"

	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/system/rest/handlers"
	"github.com/cortezaproject/corteza-server/system/service"
)

func MountRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		handlers.NewAttachment(Attachment{}.New()).MountRoutes(r)
		handlers.NewAuth((Auth{}).New()).MountRoutes(r)

		// A special case that, we do not add this through standard request, handlers & controllers
		// combo but directly -- we need access to r.Body
		r.Handle(service.SinkBaseURL+"*", &Sink{
			svc:  service.DefaultSink,
			sign: auth.DefaultSigner,
		})
	})

	// Protect all _private_ routes
	r.Group(func(r chi.Router) {
		r.Use(auth.MiddlewareValidOnly)

		handlers.NewAuthClient(AuthClient{}.New()).MountRoutes(r)
		handlers.NewAutomation(Automation{}.New()).MountRoutes(r)
		handlers.NewUser(User{}.New()).MountRoutes(r)
		handlers.NewRole(Role{}.New()).MountRoutes(r)
		handlers.NewPermissions(Permissions{}.New()).MountRoutes(r)
		handlers.NewApplication(Application{}.New()).MountRoutes(r)
		handlers.NewTemplate(Template{}.New()).MountRoutes(r)
		handlers.NewSettings(Settings{}.New()).MountRoutes(r)
		handlers.NewStats(Stats{}.New()).MountRoutes(r)
		handlers.NewReminder(Reminder{}.New()).MountRoutes(r)
		handlers.NewActionlog(Actionlog{}.New()).MountRoutes(r)
		handlers.NewQueues(Queue{}.New()).MountRoutes(r)
		handlers.NewRoute(Route{}.New()).MountRoutes(r)
		handlers.NewFunction(Function{}.New()).MountRoutes(r)
	})
}
