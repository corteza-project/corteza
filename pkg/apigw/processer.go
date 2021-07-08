package apigw

import (
	"context"
	"fmt"

	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	WfExecer interface {
		Exec(ctx context.Context, workflowID uint64, p atypes.WorkflowExecParams) (*expr.Vars, atypes.Stacktrace, error)
	}

	processerWorkflow struct {
		d WfExecer
	}
)

func NewProcesserWorkflow(wf WfExecer) processerWorkflow {
	return processerWorkflow{
		d: wf,
	}
}

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
		var (
			wfID int64
			ok   bool
			err  error
		)

		// validate workflow param
		if _, ok = params["workflow"]; !ok {
			return fmt.Errorf("invalid param workflow")
		}

		wfID, err = expr.CastToInteger(params["workflow"])

		if err != nil {
			return err
		}

		// setup scope for workflow
		vv := map[string]interface{}{
			"request": scope.Request(),
		}

		// get the request data and put it into vars
		in, err := expr.NewVars(vv)

		if err != nil {
			return err
		}

		wp := atypes.WorkflowExecParams{
			Trace: false,
			// todo depending on settings per-route
			Async: false,
			// todo depending on settings per-route
			Wait:  true,
			Input: in,
		}

		out, _, err := h.d.Exec(ctx, uint64(wfID), wp)

		if err != nil {
			return err
		}

		// merge out with scope
		merged, err := in.Merge(out)

		if err != nil {
			return err
		}

		mm, err := expr.CastToVars(merged)

		for k, v := range mm {
			scope.Set(k, v)
		}

		return err
	}
}
