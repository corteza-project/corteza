package app

import (
	"context"
	automationRest "github.com/cortezaproject/corteza-server/automation/rest"
	composeRest "github.com/cortezaproject/corteza-server/compose/rest"
	discoveryRest "github.com/cortezaproject/corteza-server/discovery/rest"
	"github.com/cortezaproject/corteza-server/docs"
	federationRest "github.com/cortezaproject/corteza-server/federation/rest"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/api/server"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/webapp"
	systemRest "github.com/cortezaproject/corteza-server/system/rest"
	"github.com/cortezaproject/corteza-server/system/scim"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func (app *CortezaApp) Serve(ctx context.Context) (err error) {
	wg := &sync.WaitGroup{}

	{ // @todo refactor wait-for out of HTTP API server.
		app.HttpServer = server.New(app.Log, app.Opt.Environment, app.Opt.HTTPServer, app.Opt.WaitFor)
		app.HttpServer.MountRoutes(app.mountHttpRoutes)

		wg.Add(1)
		go func() {
			app.HttpServer.Serve(actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_API_REST))
			wg.Done()
		}()
	}

	{
		//wg.Add(1)
		//go func(ctx context.Context) {
		//	grpcApi.Serve(actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_API_GRPC))
		//	wg.Done()
		//}(ctx)
	}

	// Wait for all servers to be done
	wg.Wait()

	return nil
}

func (app *CortezaApp) mountHttpRoutes(r chi.Router) {
	var (
		apiBaseUrl    = strings.Trim(app.Opt.HTTPServer.ApiBaseUrl, "/")
		webappBaseUrl = strings.Trim(app.Opt.HTTPServer.WebappBaseUrl, "/")
	)

	app.AuthService.MountHttpRoutes(r)

	if app.Opt.HTTPServer.ApiEnabled {

		r.Route("/"+apiBaseUrl, func(r chi.Router) {
			r.Route("/system", systemRest.MountRoutes)
			r.Route("/automation", automationRest.MountRoutes)
			r.Route("/compose", composeRest.MountRoutes)

			if app.Opt.Discovery.Enabled {
				r.Route("/discovery", discoveryRest.MountRoutes)
			}

			if app.Opt.Federation.Enabled {
				r.Route("/federation", federationRest.MountRoutes)
			}

			r.Handle("/docs", http.RedirectHandler("/"+apiBaseUrl+"/docs/", http.StatusPermanentRedirect))
			r.Handle("/docs*", http.StripPrefix("/"+apiBaseUrl+"/docs", http.FileServer(docs.GetFS())))
		})

		app.Log.Info(
			"JSON REST API enabled",
			zap.String("baseUrl", app.Opt.HTTPServer.ApiBaseUrl),
		)

		app.Log.Info(
			"API docs enabled",
			zap.String("baseUrl", app.Opt.HTTPServer.ApiBaseUrl+"/docs"),
		)
	} else {
		app.Log.Info("JSON REST API disabled")
	}

	func() {
		if !app.Opt.SCIM.Enabled {
			return
		}

		if app.Opt.SCIM.Secret == "" {
			app.Log.
				WithOptions(zap.AddStacktrace(zap.PanicLevel)).
				Error("SCIM secret empty")
		}

		var (
			baseUrl         = "/" + strings.Trim(app.Opt.SCIM.BaseURL, "/")
			extIdValidation *regexp.Regexp
			err             error
		)

		if len(app.Opt.SCIM.ExternalIdValidation) > 0 {
			extIdValidation, err = regexp.Compile(app.Opt.SCIM.ExternalIdValidation)
		}

		if err != nil {
			app.Log.Error("failed to compile SCIM external ID validation", zap.Error(err))
			return
		}

		app.Log.Debug(
			"SCIM enabled",
			zap.String("baseUrl", baseUrl),
			logger.Mask("secret", app.Opt.SCIM.Secret),
		)

		r.Route(baseUrl, func(r chi.Router) {
			if !app.Opt.Environment.IsDevelopment() {
				r.Use(scim.Guard(app.Opt.SCIM))
			}

			scim.Routes(r, scim.Config{
				ExternalIdAsPrimary: app.Opt.SCIM.ExternalIdAsPrimary,
				ExternalIdValidator: extIdValidation,
			})
		})
	}()

	if app.Opt.HTTPServer.WebappEnabled {
		r.Route("/"+webappBaseUrl, webapp.MakeWebappServer(app.Opt.HTTPServer, app.Opt.Auth, app.Opt.Federation))

		app.Log.Info(
			"client web applications enabled",
			zap.String("baseUrl", app.Opt.HTTPServer.WebappBaseUrl),
			zap.String("baseDir", app.Opt.HTTPServer.WebappBaseDir),
			zap.Strings("apps", strings.Split(app.Opt.HTTPServer.WebappList, ",")),
		)
	} else {
		app.Log.Info("client web applications disabled")
	}
}
