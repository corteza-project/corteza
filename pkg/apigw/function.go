package apigw

import (
	"context"

	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Handler interface {
		Handler() handlerFunc
		Meta(f *types.Function) functionMeta
	}

	handlerFunc func(context.Context, *scp, map[string]interface{}, functionHandler) error

	functionMeta struct {
		step   int
		weight int
		name   string
		label  string
		kind   string
		params map[string]interface{}
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
	ff.step = p.step
	ff.kind = p.kind
	ff.label = p.label
	ff.name = p.name
	ff.weight = p.weight
	ff.params = p.params
}

func (ff functionHandler) Weight() int {
	// if there's gonna be more than 1000 funcs
	// per step, we're doing something wrong
	return ff.step*1000 + ff.weight
}
