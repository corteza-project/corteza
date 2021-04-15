package options

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// pkg/options/auth.yaml

import (
	"strings"
	"time"
)

type (
	AuthOpt struct {
		LogEnabled               bool          `env:"AUTH_LOG_ENABLED"`
		Secret                   string        `env:"AUTH_JWT_SECRET"`
		Expiry                   time.Duration `env:"AUTH_JWT_EXPIRY"`
		ExternalRedirectURL      string        `env:"AUTH_EXTERNAL_REDIRECT_URL"`
		ExternalCookieSecret     string        `env:"AUTH_EXTERNAL_COOKIE_SECRET"`
		BaseURL                  string        `env:"AUTH_BASE_URL"`
		SessionCookieName        string        `env:"AUTH_SESSION_COOKIE_NAME"`
		SessionCookiePath        string        `env:"AUTH_SESSION_COOKIE_PATH"`
		SessionCookieDomain      string        `env:"AUTH_SESSION_COOKIE_DOMAIN"`
		SessionCookieSecure      bool          `env:"AUTH_SESSION_COOKIE_SECURE"`
		SessionLifetime          time.Duration `env:"AUTH_SESSION_LIFETIME"`
		SessionPermLifetime      time.Duration `env:"AUTH_SESSION_PERM_LIFETIME"`
		GarbageCollectorInterval time.Duration `env:"AUTH_GARBAGE_COLLECTOR_INTERVAL"`
		RequestRateLimit         int           `env:"AUTH_REQUEST_RATE_LIMIT"`
		RequestRateWindowLength  time.Duration `env:"AUTH_REQUEST_RATE_WINDOW_LENGTH"`
		CsrfSecret               string        `env:"AUTH_CSRF_SECRET"`
		CsrfEnabled              bool          `env:"AUTH_CSRF_ENABLED"`
		CsrfFieldName            string        `env:"AUTH_CSRF_FIELD_NAME"`
		CsrfCookieName           string        `env:"AUTH_CSRF_COOKIE_NAME"`
		DefaultClient            string        `env:"AUTH_DEFAULT_CLIENT"`
		AssetsPath               string        `env:"AUTH_ASSETS_PATH"`
		DevelopmentMode          bool          `env:"AUTH_DEVELOPMENT_MODE"`
	}
)

// Auth initializes and returns a AuthOpt with default values
func Auth() (o *AuthOpt) {
	o = &AuthOpt{
		Secret:                   getSecretFromEnv("jwt secret"),
		Expiry:                   time.Hour * 24 * 30,
		ExternalRedirectURL:      guessBaseURL() + "/auth/external/{provider}/callback",
		ExternalCookieSecret:     getSecretFromEnv("external cookie secret"),
		BaseURL:                  guessBaseURL() + "/auth",
		SessionCookieName:        "session",
		SessionCookiePath:        "/auth",
		SessionCookieDomain:      guessHostname(),
		SessionCookieSecure:      strings.HasPrefix(guessBaseURL(), "https://"),
		SessionLifetime:          24 * time.Hour,
		SessionPermLifetime:      360 * 24 * time.Hour,
		GarbageCollectorInterval: 15 * time.Minute,
		RequestRateLimit:         30,
		RequestRateWindowLength:  time.Minute,
		CsrfSecret:               getSecretFromEnv("csrf secret"),
		CsrfEnabled:              true,
		CsrfFieldName:            "same-site-authenticity-token",
		CsrfCookieName:           "same-site-authenticity-token",
		DefaultClient:            "corteza-webapp",
		AssetsPath:               "",
	}

	fill(o)

	// Function that allows access to custom logic inside the parent function.
	// The custom logic in the other file should be like:
	// func (o *Auth) Defaults() {...}
	func(o interface{}) {
		if def, ok := o.(interface{ Defaults() }); ok {
			def.Defaults()
		}
	}(o)

	return
}
