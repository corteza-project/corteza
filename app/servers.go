package app

import (
	"context"
	"net/http"
	"path"
	"regexp"
	"strings"
	"sync"

	automationRest "github.com/cortezaproject/corteza-server/automation/rest"
	composeRest "github.com/cortezaproject/corteza-server/compose/rest"
	"github.com/cortezaproject/corteza-server/compose/service"
	"github.com/cortezaproject/corteza-server/docs"
	federationRest "github.com/cortezaproject/corteza-server/federation/rest"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/api/server"
	"github.com/cortezaproject/corteza-server/pkg/apigw"
	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/options"
	"github.com/cortezaproject/corteza-server/pkg/webapp"
	systemRest "github.com/cortezaproject/corteza-server/system/rest"
	"github.com/cortezaproject/corteza-server/system/scim"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
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
		ho = app.Opt.HTTPServer
	)

	func() {
		if ho.WebappEnabled && ho.ApiEnabled && ho.ApiBaseUrl == ho.WebappBaseUrl {
			app.Log.
				Warn("client web applications and api can not use the same base URL: '" + ho.WebappBaseUrl + "'")
			ho.WebappEnabled = false
		}

		if !ho.WebappEnabled {
			app.Log.Info("client web applications disabled")
			return
		}

		r.Route(options.CleanBase(ho.WebappBaseUrl), webapp.MakeWebappServer(app.Log, ho, app.Opt.Auth))

		app.Log.Info(
			"client web applications enabled",
			zap.String("baseUrl", options.CleanBase(ho.BaseUrl, ho.WebappBaseUrl)),
			zap.String("baseDir", ho.WebappBaseDir),
			zap.Strings("apps", strings.Split(ho.WebappList, ",")),
		)
	}()

	// Auth server
	app.AuthService.MountHttpRoutes(ho.BaseUrl, r)

	func() {
		if !ho.ApiEnabled {
			app.Log.Info("JSON REST API disabled")
			return
		}

		r.Route(options.CleanBase(ho.ApiBaseUrl), func(r chi.Router) {
			var fullpathAPI = "/" + strings.TrimPrefix(options.CleanBase(ho.BaseUrl, ho.ApiBaseUrl), "/")

			app.Log.Info(
				"JSON REST API enabled",
				zap.String("baseUrl", fullpathAPI),
			)

			r.Route("/system", systemRest.MountRoutes)
			r.Route("/automation", automationRest.MountRoutes)
			r.Route("/compose", composeRest.MountRoutes)
			r.Route("/websocket", app.WsServer.MountRoutes)

			if app.Opt.Federation.Enabled {
				r.Route("/federation", federationRest.MountRoutes)
			}

			// temp api gateway support
			{
				apigw.Setup(struct{}{}, app.Log, eventbus.Service(), service.DefaultStore)
				r.Route("/gateway", apigw.Service().Router(context.Background()))
			}

			var fullpathDocs = options.CleanBase(ho.BaseUrl, ho.ApiBaseUrl, "docs")
			app.Log.Info(
				"API docs enabled",
				zap.String("baseUrl", fullpathDocs),
			)

			r.Handle("/docs", http.RedirectHandler(fullpathDocs+"/", http.StatusPermanentRedirect))
			r.Handle("/docs*", http.StripPrefix(fullpathDocs, http.FileServer(docs.GetFS())))
		})
	}()

	func() {
		if !app.Opt.SCIM.Enabled {
			return
		}

		if app.Opt.SCIM.Secret == "" {
			app.Log.
				Error("SCIM secret empty")
		}

		var (
			baseUrl         = app.Opt.SCIM.BaseURL
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
			zap.String("baseUrl", path.Join(app.Opt.HTTPServer.BaseUrl, baseUrl)),
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
}
