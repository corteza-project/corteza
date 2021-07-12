package apigw

// match auth headers
// which header
// value of the header

import (
	"context"
	"fmt"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/davecgh/go-spew/spew"
)

type (
	validatorHeader struct{}
)

func NewValidatorHeader() validatorHeader {
	return validatorHeader{}
}

func (h validatorHeader) Meta(f *types.ApigwFunction) functionMeta {
	return functionMeta{
		Step:   1,
		Name:   "validatorHeader",
		Label:  "Header validator",
		Kind:   "validator",
		Weight: int(f.Weight),
		Params: f.Params,
		Args: []*functionMetaArg{
			{
				Type:    "expr",
				Label:   "expr",
				Options: map[string]interface{}{},
			},
		},
	}
}

func (h validatorHeader) Handler() handlerFunc {
	return func(ctx context.Context, scope *scp, params map[string]interface{}, ff functionHandler) error {
		for k := range ff.params {
			v, ok := params[k]

			if !ok {
				spew.Dump("not in params", k)
				continue
			}

			vv := map[string]interface{}{}
			headers := scope.Request().Header

			for k, v := range headers {
				vv[k] = v[0]
			}

			// get the request data and put it into vars
			out, err := expr.NewVars(vv)

			if err != nil {
				// spew.Dump("ERR!", err)
				return err
			}

			pp := expr.NewParser()
			tt, err := pp.Parse(v.(string))

			if err != nil {
				return err
			}

			b, err := tt.Test(ctx, out)

			if err != nil {
				return fmt.Errorf("could not validate headers: %s", err)
			}

			spew.Dump("BBBB", b)

			if !b {
				return fmt.Errorf("could not validate headers, failed on step %d, function %s", ff.step, ff.name)
			}
		}

		return nil
	}
}
