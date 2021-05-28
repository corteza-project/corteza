package apigw

import (
	"context"
	"errors"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/davecgh/go-spew/spew"
)

type (
	authenticationOriginMatcherArgs struct {
		Origin string
	}
)

func authenticationOriginMatcher(c context.Context, params *expr.Vars) wfHandler {
	var (
		aomp = authenticationOriginMatcherArgs{}
	)

	return func(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {

		params.Decode(&aomp)
		spew.Dump("authentication origin matcher fn()", aomp)
		e := er.Scope.GetValue()["envelope"]
		ee := e.Get().(envelope)

		origin := ee.Request.Header.Get("Origin")

		spew.Dump("input, real", aomp.Origin, origin)

		if aomp.Origin != origin {
			err = errors.New("origin fail")
			return
		}

		r = &expr.Vars{}
		return
	}
}
