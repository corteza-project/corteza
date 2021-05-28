package apigw

import (
	"context"
	"net/http"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/davecgh/go-spew/spew"
)

type (
	route struct {
		endpoint string
		method   string
		graph    *wfexec.Graph
		steps    []wfexec.Step
		fns      wfHandlerList
	}
)

func (r route) validate(req *http.Request) (err error) {
	// if req.Method != r.method {
	// 	err = errors.New("http method invalid")
	// }

	return
}

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := r.validate(req); err != nil {
		spew.Dump("ERR", err)
		return
	}

	sess := wfexec.NewSession(context.Background(), r.graph, wfexec.SetLogger(logger.Default()), wfexec.SetHandler(func(ss wfexec.SessionStatus, s1 *wfexec.State, s2 *wfexec.Session) {
		// spew.Dump("event handler here!", ss)
	}))

	scope := &expr.Vars{}

	scope.Set("envelope", envelope{
		Request: req,
		Writer:  w,
	})

	if len(r.steps) == 0 {
		// dont serve, do what? return default response?
		return
	}

	err := sess.Exec(context.Background(), r.steps[0], scope)

	// if err != nil {
	// 	fmt.Fprintf(w, "no go, err on exec: %s", err)
	// 	return
	// }

	err = sess.Wait(context.Background())

	if err != nil {
	}
}
