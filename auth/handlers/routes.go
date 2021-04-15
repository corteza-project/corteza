package handlers

import (
	"github.com/cortezaproject/corteza-server/auth/request"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
	"github.com/gorilla/csrf"
	"net/http"
)

func (h *AuthHandlers) MountHttpRoutes(r chi.Router) {
	var (
		l = GetLinks()
	)

	if h.Opt.DevelopmentMode {
		r.Get("/auth/dev", h.handle(h.devView))
		r.Get("/auth/dev/scenarios", h.devSceneView)
	}

	r.Handle("/auth/", http.RedirectHandler("/auth", http.StatusSeeOther))
	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := actionlog.RequestOriginToContext(r.Context(), actionlog.RequestOrigin_Auth)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		})

		if h.Opt.RequestRateLimit > 0 {
			r.Use(httprate.LimitByIP(h.Opt.RequestRateLimit, h.Opt.RequestRateWindowLength)) // @todo make configurable
		}

		r.Use(request.ExtraReqInfoMiddleware)

		r.Group(func(r chi.Router) {
			// all routes protected with CSRF:
			if h.Opt.CsrfEnabled {
				r.Use(csrf.Protect(
					[]byte(h.Opt.CsrfSecret),
					csrf.SameSite(csrf.SameSiteStrictMode),
					csrf.Secure(h.Opt.SessionCookieSecure),
					csrf.FieldName(h.Opt.CsrfFieldName),
					csrf.CookieName(h.Opt.CsrfCookieName),
				))
			}

			r.Get(l.Profile, h.handle(authOnly(h.profileForm)))
			r.Post(l.Profile, h.handle(authOnly(h.profileProc)))

			r.HandleFunc(l.Logout, h.handle(h.logoutProc))

			r.Get(l.Sessions, h.handle(authOnly(h.sessionsView)))
			r.Post(l.Sessions, h.handle(authOnly(h.sessionsProc)))

			r.Get(l.AuthorizedClients, h.handle(authOnly(h.clientsView)))
			r.Post(l.AuthorizedClients, h.handle(authOnly(h.clientsProc)))

			r.Get(l.Signup, h.handle(h.onlyIfSignupEnabled(anonyOnly(h.signupForm))))
			r.Post(l.Signup, h.handle(h.onlyIfSignupEnabled(anonyOnly(h.signupProc))))
			r.Get(l.PendingEmailConfirmation, h.handle(h.pendingEmailConfirmation))
			r.Get(l.ConfirmEmail, h.handle(h.confirmEmail))

			r.Get(l.Login, h.handle(anonyOnly(h.loginForm)))
			r.Post(l.Login, h.handle(h.onlyIfLocalEnabled(anonyOnly(h.loginProc))))

			r.Get(l.Mfa, h.handle(h.mfaForm))
			r.Post(l.Mfa, h.handle(h.mfaProc))

			r.Get(l.RequestPasswordReset, h.handle(h.onlyIfPasswordResetEnabled(anonyOnly(h.requestPasswordResetForm))))
			r.Post(l.RequestPasswordReset, h.handle(h.onlyIfPasswordResetEnabled(anonyOnly(h.requestPasswordResetProc))))
			r.Get(l.PasswordResetRequested, h.handle(h.onlyIfPasswordResetEnabled(anonyOnly(h.passwordResetRequested))))
			r.Get(l.ResetPassword, h.handle(h.onlyIfPasswordResetEnabled(h.resetPasswordForm)))
			r.Post(l.ResetPassword, h.handle(h.onlyIfPasswordResetEnabled(authOnly(h.resetPasswordProc))))

			r.Get(l.Security, h.handle(authOnly(h.securityForm)))
			r.Post(l.Security, h.handle(authOnly(h.securityProc)))
			r.Get(l.ChangePassword, h.handle(h.onlyIfLocalEnabled(authOnly(h.changePasswordForm))))
			r.Post(l.ChangePassword, h.handle(h.onlyIfLocalEnabled(authOnly(h.changePasswordProc))))

			r.Get(l.MfaTotpNewSecret, h.handle(partAuthOnly(h.mfaTotpConfigForm)))
			r.Post(l.MfaTotpNewSecret, h.handle(partAuthOnly(h.mfaTotpConfigProc)))
			r.Get(l.MfaTotpQRImage, h.handle(partAuthOnly(h.mfaTotpConfigQR)))
			r.Get(l.MfaTotpDisable, h.handle(authOnly(h.mfaTotpDisableForm)))
			r.Post(l.MfaTotpDisable, h.handle(authOnly(h.mfaTotpDisableProc)))

		})

		r.Group(func(r chi.Router) {
			// OAuth2 routes
			r.HandleFunc(l.OAuth2Authorize, h.handle(h.oauth2Authorize))
			r.Get(l.OAuth2AuthorizeClient, h.handle(authOnly(h.oauth2AuthorizeClient)))
			r.Post(l.OAuth2AuthorizeClient, h.handle(authOnly(h.oauth2AuthorizeClientProc)))
			r.Get(l.OAuth2DefaultClient, h.handle(h.oauth2authorizeDefaultClient))
			r.Post(l.OAuth2DefaultClient, h.handle(h.oauth2authorizeDefaultClientProc))
		})

		r.Route(l.External+"/{provider}", func(r chi.Router) {
			// External provider
			r.Get("/", h.externalInit)
			r.Get("/callback", h.externalCallback)
		})

		r.HandleFunc("/auth/oauth2/token", h.handle(h.oauth2Token))
		r.HandleFunc("/auth/oauth2/info", h.oauth2Info)
	})
}
