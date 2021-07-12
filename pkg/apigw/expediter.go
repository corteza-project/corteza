package apigw

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/davecgh/go-spew/spew"
)

type (
	expediterRedirection struct{}

	errorHandler struct {
		name   string
		args   []string
		weight int
		step   int
	}
)

func NewExpediterRedirection() expediterRedirection {
	return expediterRedirection{}
}

func (h expediterRedirection) Meta(f *types.ApigwFunction) functionMeta {
	return functionMeta{
		Step:   3,
		Name:   "expediterRedirection",
		Label:  "Redirection expediter",
		Kind:   "expediter",
		Weight: int(f.Weight),
		Params: f.Params,
	}
}

func (h expediterRedirection) Handler() handlerFunc {
	return func(ctx context.Context, scope *scp, params map[string]interface{}, ff functionHandler) error {
		scope.Writer().Header().Add(fmt.Sprintf("step_%d", ff.step), ff.name)
		http.Redirect(scope.Writer(), scope.Request(), params["location"].(string), http.StatusFound)

		return nil
	}
}

func (pp errorHandler) Exec(ctx context.Context, scope *scp, err error) {
	type (
		responseHelper struct {
			Msg string `json:"msg"`
		}
	)

	resp := responseHelper{
		Msg: err.Error(),
	}
	spew.Dump("ERR in expediter", err, resp)

	json.NewEncoder(scope.Writer()).Encode(resp)
}
