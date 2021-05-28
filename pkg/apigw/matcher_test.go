package apigw

import (
	"context"
	"net/http"
	"testing"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationOriginMatcher(t *testing.T) {
	type (
		tf struct {
			name   string
			origin string
			exp    string
			req    *http.Request
		}
	)

	var (
		ctx = context.Background()

		tcc = []tf{
			{
				name:   "fail on origin",
				origin: "http://fail.ed",
				exp:    "workflow 0 step 0 execution failed: origin fail",

				req: &http.Request{
					Header: http.Header{
						"Origin": []string{
							"http://localhost",
						},
					},
				},
			},
			{
				name:   "success on origin",
				origin: "http://localhost",
				exp:    "",

				req: &http.Request{
					Header: http.Header{
						"Origin": []string{
							"http://localhost",
						},
					},
				},
			},
		}
	)

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			var (
				req   = require.New(t)
				input = &expr.Vars{}
			)

			input.Set("origin", tc.origin)

			err := execFn(t, tc.req, authenticationOriginMatcher(ctx, input))

			if tc.exp != "" {
				req.EqualError(err, tc.exp)
			} else {
				req.NoError(err)
			}
		})
	}

}
