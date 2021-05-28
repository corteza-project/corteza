package apigw

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/davecgh/go-spew/spew"
)

func formDataProcesserFn(c context.Context, er *wfexec.ExecRequest) (r wfexec.ExecResponse, err error) {
	type (
		formDataProcesserResponse struct {
			Name string `json:"name"`
		}
	)

	spew.Dump("step processer fn()")

	e := er.Scope.GetValue()["envelope"]
	ee := e.Get()

	// ee.(envelope).Writer.WriteHeader(int(id))
	ee.(envelope).Writer.Write([]byte(`{"test":"foobar"}`))

	e.Assign(ee)

	// req := values.Get()
	// ww := wr.Get()
	// writer := ww.(http.ResponseWriter)

	// formValue := req.(*http.Request).PostFormValue("name")

	// resp := formDataProcesserResponse{
	// 	// Name: fmt.Sprintf("AA %s AA", formValue),
	// 	Name: "formValue",
	// }

	// encoder := json.NewEncoder(writer)
	// encoder.Encode(resp)

	// writer.(*httptest.ResponseRecorder).Header()["Content-Type"] = []string{"application/json"}
	// writer.Header().Set("Content-Type", "application/json3")

	// spew.Dump(writer.(*httptest.ResponseRecorder).Header())
	// a, b := expr.NewKV(writer)
	// spew.Dump("Aaaaaaaaaaa", a)

	vv := &expr.Vars{}
	// vv.Set("writer", writer)

	r = vv

	return
}
