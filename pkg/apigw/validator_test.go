package apigw

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/stretchr/testify/require"
)

func TestContentLengthValidator(t *testing.T) {
	type (
		tf struct {
			name  string
			limit int
			exp   string
			body  string
		}
	)

	var (
		ctx = context.Background()

		tcc = []tf{
			{
				name:  "fail on content length > limit",
				limit: 10,
				exp:   "workflow 0 step 0 execution failed: content length overriden",
				body:  "A message that is 31 bytes long",
			},
			{
				name:  "success on content length < limit",
				limit: 10,
				exp:   "",
				body:  "Below 10",
			},
		}
	)

	for _, tc := range tcc {
		t.Run(tc.name, func(t *testing.T) {
			var (
				req   = require.New(t)
				input = &expr.Vars{}
			)

			input.Set("length", tc.limit)

			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))

			err := execFn(t, r, contentLengthValidator(ctx, input))

			if tc.exp != "" {
				req.EqualError(err, tc.exp)
			} else {
				req.NoError(err)
			}
		})
	}
}
