package apigw

import (
	"context"
	"fmt"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/davecgh/go-spew/spew"
)

type (
	verifierQueryParam struct{}
	verifierOrigin     struct{}
)

func (h verifierQueryParam) Meta(f *types.Function) functionMeta {
	return functionMeta{
		Step:   0,
		Name:   "verifierQueryParam",
		Label:  "Query parameters verifier",
		Kind:   "verifier",
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

func (h verifierOrigin) Meta(f *types.Function) functionMeta {
	return functionMeta{
		Step:   0,
		Name:   "verifierOrigin",
		Label:  "Origin verifier",
		Kind:   "verifier",
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

func (h verifierQueryParam) Handler() handlerFunc {
	return func(ctx context.Context, scope *scp, params map[string]interface{}, ff functionHandler) error {
		for k := range ff.params {

			v, ok := params[k]

			if !ok {
				spew.Dump("not in params", k)
				continue
			}

			vv := map[string]interface{}{}
			vals := scope.req.URL.Query()

			for k, v := range vals {
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
				// spew.Dump("ERR!", err)
				return err
			}

			b, err := tt.Test(ctx, out)

			if err != nil {
				// spew.Dump("ERR!", err)
				return err
			}

			spew.Dump("BBBB", b)

			if !b {
				return fmt.Errorf("failed on step %d, function %s", ff.step, ff.name)
			}
		}

		// testing
		scope.req.Header.Add(fmt.Sprintf("step_%d", ff.step), ff.name)

		return nil
	}
}

func (h verifierOrigin) Handler() handlerFunc {
	return func(ctx context.Context, scope *scp, params map[string]interface{}, ff functionHandler) error {
		for k := range ff.params {
			v, ok := params[k]

			if !ok {
				spew.Dump("not in params", k)
				continue
			}

			vv := map[string]interface{}{
				"origin": scope.req.Header.Get("Origin"),
			}

			// get the request data and put it into vars
			out, err := expr.NewVars(vv)

			if err != nil {
				spew.Dump("ERR!", err)
				return err
			}

			pp := expr.NewParser()
			tt, err := pp.Parse(v.(string))

			if err != nil {
				spew.Dump("ERR!", err)
				return err
			}

			b, err := tt.Test(ctx, out)

			if err != nil {
				spew.Dump("ERR!", err)
				return err
			}

			spew.Dump("BBBB", b)

			if !b {
				return fmt.Errorf("failed on step %d, function %s", ff.step, ff.name)
			}
		}

		// testing
		scope.req.Header.Add(fmt.Sprintf("step_%d", ff.step), ff.name)

		return nil
	}
}
