package apigw

import (
	"context"
	"errors"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/davecgh/go-spew/spew"
)

type (
	contentLengthValidatorArgs struct {
		Length int
	}
)

func contentLengthValidator(c context.Context, params *expr.Vars) wfHandler {
	var (
		clv = contentLengthValidatorArgs{}
	)

	return func(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {

		params.Decode(&clv)
		spew.Dump("body size validator fn()", clv)
		e := er.Scope.GetValue()["envelope"]
		ee := e.Get().(envelope)

		cl := ee.Request.ContentLength

		if clv.Length < int(cl) {
			err = errors.New("content length overriden")
			return
		}

		r = &expr.Vars{}
		return
	}
}
