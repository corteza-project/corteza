package apigw

import (
	"context"

	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	functionMetaList []*functionMeta

	Handler interface {
		Handler() handlerFunc
		Meta(f *types.Function) functionMeta
	}

	handlerFunc func(context.Context, *scp, map[string]interface{}, functionHandler) error

	functionMeta struct {
		Step   int                    `json:"step"`
		Weight int                    `json:"-"`
		Name   string                 `json:"name"`
		Label  string                 `json:"label"`
		Kind   string                 `json:"kind"`
		Params map[string]interface{} `json:"-"`
		Args   []*functionMetaArg     `json:"params,omitempty"`
	}

	functionMetaArg struct {
		Label   string                 `json:"label"`
		Type    string                 `json:"type"`
		Example string                 `json:"example"`
		Options map[string]interface{} `json:"options"`
	}

	functionHandler struct {
		step    int
		weight  int
		name    string
		label   string
		kind    string
		handler handlerFunc
		params  map[string]interface{}
	}
)

func (ff functionHandler) Exec(ctx context.Context, scope *scp, params map[string]interface{}) error {
	return ff.handler(ctx, scope, params, ff)
}

func (ff *functionHandler) SetHandler(h handlerFunc) {
	ff.handler = h
}

func (ff *functionHandler) Merge(ctx context.Context, p functionMeta) {
	ff.step = p.Step
	ff.kind = p.Kind
	ff.label = p.Label
	ff.name = p.Name
	ff.weight = p.Weight
	ff.params = p.Params
}

func (ff functionHandler) Weight() int {
	// if there's gonna be more than 1000 funcs
	// per step, we're doing something wrong
	return ff.step*1000 + ff.weight
}

func (fm functionMetaList) Filter(f func(*functionMeta) (bool, error)) (out functionMetaList, err error) {
	var ok bool
	out = functionMetaList{}
	for i := range fm {
		if ok, err = f(fm[i]); err != nil {
			return
		} else if ok {
			out = append(out, fm[i])
		}
	}

	return
}
