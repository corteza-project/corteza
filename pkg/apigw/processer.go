package apigw

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	dispatcher interface {
		Dispatch(ctx context.Context, ev eventbus.Event)
	}

	processerWorkflow struct {
		d dispatcher
	}
)

func (h processerWorkflow) Meta(f *types.Function) functionMeta {
	return functionMeta{
		Step:   2,
		Name:   "processerWorkflow",
		Label:  "Workflow processer",
		Kind:   "processer",
		Weight: int(f.Weight),
		Params: f.Params,
		Args: []*functionMetaArg{
			{
				Type:    "workflow",
				Label:   "workflow",
				Options: map[string]interface{}{},
			},
		},
	}
}

func (h processerWorkflow) Handler() handlerFunc {
	return func(ctx context.Context, scope *scp, params map[string]interface{}, ff functionHandler) error {
		// h.d.Dispatch(c, event.ApiOnProcess(&envlp))

		return nil
	}
}

// func formDataProcesserFn(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {
// 	type (
// 		formDataProcesserResponse struct {
// 			Name string `json:"name"`
// 		}
// 	)

// 	spew.Dump("step processer fn()")

// 	e := er.Scope.GetValue()["envelope"]
// 	ee := e.Get()

// 	// ee.(envelope).Writer.WriteHeader(int(id))
// 	ee.(envelope).Writer.Write([]byte(`{"test":"foobar"}`))

// 	e.Assign(ee)

// 	// req := values.Get()
// 	// ww := wr.Get()
// 	// writer := ww.(http.ResponseWriter)

// 	// formValue := req.(*http.Request).PostFormValue("name")

// 	// resp := formDataProcesserResponse{
// 	// 	// Name: fmt.Sprintf("AA %s AA", formValue),
// 	// 	Name: "formValue",
// 	// }

// 	// encoder := json.NewEncoder(writer)
// 	// encoder.Encode(resp)

// 	// writer.(*httptest.ResponseRecorder).Header()["Content-Type"] = []string{"application/json"}
// 	// writer.Header().Set("Content-Type", "application/json3")

// 	// spew.Dump(writer.(*httptest.ResponseRecorder).Header())
// 	// a, b := expr.NewKV(writer)
// 	// spew.Dump("Aaaaaaaaaaa", a)

// 	vv := &expr.Vars{}
// 	// vv.Set("writer", writer)

// 	r = vv

// 	return
// }
