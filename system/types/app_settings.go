package types

import (
	"errors"
	"strings"
)

type (
	// AppSettings type is structured representation of all application settings
	//
	// Raw settings keys are hypen (kebab) case, separated with a dot (.) that indicates sub-level
	// JSON properties for settings are NOT converted (lower-cased, etc...)
	// Use `json:"-"` tag to hide settings on REST endpoint
	AppSettings struct {
		Privacy struct {
			Mask struct {
				// Enable masking of user's email (value replaced with ######)
				Email bool

				// Enable masking of user's name (value replaced with ######)
				Name bool
			}
		} `json:"-"`

		General struct {
			Mail struct {
				Logo   string
				Header string `kv:"header.en"`
				Footer string `kv:"footer.en"`
			}
		} `json:"-"`

		Auth struct {
			Internal struct {
				// Is internal authentication (username + password) enabled
				Enabled bool

				Signup struct {
					// Can users register
					Enabled bool

					// Users must confirm their emails when signing-up
					EmailConfirmationRequired bool `kv:"email-confirmation-required"`
				}

				// Can users reset their passwords
				PasswordReset struct{ Enabled bool } `kv:"password-reset"`
			}

			External struct {
				// Is external authentication
				Enabled bool

				// all external providers we know
				Providers ExternalAuthProviderSet
			}

			MultiFactor struct {
				EmailOTP struct {
					// Can users use email for MFA
					Enabled bool

					// Is MFA with email enforced?
					Enforced bool

					// Require fresh Email OTP on every client authorization
					//Strict bool

					Expires uint
				} `kv:"email-otp"`

				TOTP struct {
					// Can users use TOTP for MFA
					Enabled bool

					// Is MFA with TOTP enforced?
					Enforced bool

					// Require fresh TOTP on every client authorization
					//Strict bool

					// TOTP issuer, defaults to "Corteza"
					Issuer string
				} `kv:"totp"`
			} `kv:"multi-factor"`

			Mail struct {
				FromAddress string `kv:"from-address"`
				FromName    string `kv:"from-name"`

				EmailConfirmation struct {
					Subject string `kv:"subject.en"`
					Body    string `kv:"body.en"`
				} `kv:"email-confirmation"`

				PasswordReset struct {
					Subject string `kv:"subject.en"`
					Body    string `kv:"body.en"`
				} `kv:"password-reset"`
			} `json:"-"`
		} `json:"-"`

		UI struct {
			// Corteza One configuration settings
			One struct {
				Logo  string `kv:"logo" json:"logo"`
				Title string `kv:"title" json:"title"`

				PanelsEnabled bool `kv:"panels-enabled" json:"panelsEnabled"`
				Panels        []struct {
					Visible        bool `json:"visible"`
					Sticky         bool `json:"sticky"`
					Width          uint `json:"width"`
					Height         uint `json:"height"`
					ActiveTabIndex uint `json:"activeTabIndex"`
					Tabs           []struct {
						Title  string `json:"title"`
						Url    string `json:"url"`
						Logo   string `json:"logo"`
						Icon   string `json:"icon"`
						Sticky bool   `json:"sticky"`
					} `json:"tabs"`
				} `kv:"panels,final" json:"panels"`
			} `kv:"one" json:"one"`

			// Admin struct {} `kv:"admin"`
		} `kv:"ui" json:"ui"`

		Compose struct {
			// UI related settings
			UI struct {
				// Emoji
				// @todo implementation
				NamespaceSwitcher struct {
					Enabled     bool
					DefaultOpen bool
				} `kv:"namespace-switcher"`
			} `kv:"ui"`

			// Message related settings
			Record struct {
				// @todo implementation
				Attachments struct {
					// What is max size (in MB, so: MaxSize x 2^20)
					MaxSize uint `kv:"max-size"`

					// List of mime-types we support,
					Mimetypes []string
				}
			}

			// Page related settings
			Page struct {
				// @todo implementation
				Attachments struct {
					// What is max size (in MB, so: MaxSize x 2^20)
					MaxSize uint `kv:"max-size"`

					// List of mime-types we support,
					Mimetypes []string
				}
			}
		} `kv:"compose" json:"compose"`

		Discovery struct {
			// Enable indexing
			Enabled bool

			SystemUsers struct {
				// Enable indexing of users
				Enabled bool
			} `kv:"system-users"`

			SystemApplications struct {
				// Enable indexing of applications
				Enabled bool
			} `kv:"system-applications"`

			SystemRoles struct {
				// Enable indexing of roles
				Enabled bool
			} `kv:"system-roles"`

			SystemTemplates struct {
				// Enable indexing of templates
				Enabled bool
			} `kv:"system-templates"`

			AutomationWorkflows struct {
				// Enable indexing of workflows
				Enabled bool
			} `kv:"automation-workflows"`

			ComposeNamespaces struct {
				// Enable indexing of compose namespaces
				Enabled bool
			} `kv:"compose-namespaces"`

			ComposeCharts struct {
				// Enable indexing of compose charts
				Enabled bool
			} `kv:"compose-charts"`

			ComposePages struct {
				// Enable indexing of compose pages
				Enabled bool
			} `kv:"compose-pages"`

			ComposeModules struct {
				// Enable indexing of compose modules
				Enabled bool
			} `kv:"compose-modules"`

			ComposeRecords struct {
				// Enable indexing of compose records
				Enabled bool
			} `kv:"compose-records"`
		} `kv:"discovery"`
	}

	ExternalAuthProviderSet []*ExternalAuthProvider

	ExternalAuthProvider struct {
		Enabled     bool   `json:"enabled"`
		Handle      string `json:"handle"`
		Label       string `json:"label"`
		Key         string `json:"-"`
		Secret      string `json:"-"`
		RedirectUrl string `json:"-" kv:"redirect"`
		IssuerUrl   string `json:"-" kv:"issuer"`
		Weight      int    `json:"-"`
	}
)

// DecodeKV translates settings' KV into internal system external auth settings
func (set *ExternalAuthProviderSet) DecodeKV(kv SettingsKV, prefix string) (err error) {
	if *set == nil {
		*set = ExternalAuthProviderSet{}
	}

	// create standard provider set
	providers := map[string]bool{"github": true, "facebook": true, "google": true, "linkedin": true}

	// remove prefix
	kv = kv.CutPrefix(prefix + ".")

	// add all additional providers (prefixed with "openid-connect.")
	oidcPrefix := "openid-connect."
	for p := range kv {
		if !strings.HasPrefix(p, oidcPrefix) {
			continue
		}

		l := len(oidcPrefix)
		dotPos := strings.Index(p[l:], ".") + l
		if dotPos > 0 {
			providers[p[:dotPos]] = true
		}
	}

	// go over all added providers again add decode KV into each one
	for handle := range providers {
		p := (*set).FindByHandle(handle)
		if p == nil {
			p = &ExternalAuthProvider{Handle: handle}
			(*set) = append((*set), p)
		}

		err = DecodeKV(kv.CutPrefix(handle+"."), p)
		if err != nil {
			return
		}

		if p.Label == "" {
			switch p.Handle {
			case "github":
				p.Label = "GitHub"
			case "linkedin":
				p.Label = "LinkedIn"
			case "corteza-iam", "corteza", "corteza-one":
				p.Label = "Corteza IAM"
			case "crust-iam", "crust", "crust-unify":
				p.Label = "Crust IAM"
			default:
				if strings.HasPrefix(p.Handle, oidcPrefix) {
					p.Label = strings.Title(p.Handle[len(oidcPrefix):])
				} else {
					p.Label = strings.Title(p.Handle)
				}
			}
		}
	}

	return
}

func (set ExternalAuthProviderSet) FindByHandle(handle string) *ExternalAuthProvider {
	for p := range set {
		if set[p].Handle == handle {
			return set[p]
		}
	}

	return nil
}

func (set ExternalAuthProviderSet) Len() int      { return len(set) }
func (set ExternalAuthProviderSet) Swap(i, j int) { set[i], set[j] = set[j], set[i] }
func (set ExternalAuthProviderSet) Less(i, j int) bool {
	if set[i].Weight != set[j].Weight {
		// Sort by weight
		return set[i].Weight < set[j].Weight
	}

	if set[i].Label+set[j].Label != "" {
		// If at least one of the
		return set[i].Label < set[j].Label
	}

	return set[i].Handle < set[j].Handle
}

// Returns enabled providers, sorted with their redirect-URLs set...
func (set ExternalAuthProviderSet) Valid() (out ExternalAuthProviderSet) {
	for _, eap := range set {
		if !eap.Enabled {
			continue
		}

		out = append(out, eap)
	}

	return
}

var _ KVDecoder = &ExternalAuthProviderSet{}

func (p ExternalAuthProvider) EncodeKV() (vv SettingValueSet, err error) {
	if p.Handle == "" {
		return nil, errors.New("cannot encode external auth provider without handle")
	}
	var (
		prefix = "auth.external.providers." + p.Handle + "."
		pairs  = map[string]interface{}{
			"enabled":  p.Enabled,
			"label":    p.Label,
			"key":      p.Key,
			"secret":   p.Secret,
			"issuer":   p.IssuerUrl,
			"redirect": p.RedirectUrl,
			"weight":   p.Weight,
		}
	)

	for key, value := range pairs {
		v := &SettingValue{Name: prefix + key}

		if err = v.SetValue(value); err != nil {
			return
		}

		vv = append(vv, v)
	}

	return
}
