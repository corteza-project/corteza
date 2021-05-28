package apigw

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/stretchr/testify/require"
)

func execFn(t *testing.T, r *http.Request, fn wfHandler) error {
	var (
		req      = require.New(t)
		ctx      = context.Background()
		scope    = &expr.Vars{}
		graph    = wfexec.NewGraph()
		recorder = httptest.NewRecorder()
	)

	scope.Set("envelope", envelope{
		Request: r,
		Writer:  recorder,
	})

	step := wfexec.NewGenericStep(fn.self())

	graph.AddStep(step)

	sess := wfexec.NewSession(ctx, graph, wfexec.SetLogger(logger.Default()))

	err := sess.Exec(ctx, step, scope)

	req.NoError(err)

	return sess.Wait(ctx)
}
