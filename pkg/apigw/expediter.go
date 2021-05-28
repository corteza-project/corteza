package apigw

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/davecgh/go-spew/spew"
)

type (
	redirectExpediterArgs struct {
		Location string
	}
)

func redirectExpediter(c context.Context, params *expr.Vars) wfHandler {
	var (
		clv = redirectExpediterArgs{}
	)

	return func(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {
		params.Decode(&clv)
		spew.Dump("redirect expediter fn()", clv)
		e := er.Scope.GetValue()["envelope"]
		ee := e.Get().(envelope)

		http.Redirect(ee.Writer, ee.Request, clv.Location, http.StatusTemporaryRedirect)

		r = &expr.Vars{}
		return
	}
}

func expediterErrorFn(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {
	// spew.Dump("expediter error fn()", er)

	e := er.Scope.GetValue()["error"]
	values := er.Scope.GetValue()["writer"]

	writer := values.Get()
	eValue := e.Get()

	fmt.Fprintf(writer.(*httptest.ResponseRecorder), fmt.Sprintf(`{"msg": "%s"}`, eValue))
	writer.(*httptest.ResponseRecorder).Code = http.StatusBadGateway

	r = &expr.Vars{}
	return
}
