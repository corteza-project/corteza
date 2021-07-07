package apigw

import (
	"fmt"

	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	registry struct {
		h map[string]Handler
	}
)

func NewRegistry() *registry {
	return &registry{
		h: map[string]Handler{},
	}
}

func (r *registry) Add(n string, h Handler) {
	r.h[n] = h
}

func (r *registry) Get(identifier string) (Handler, error) {
	var (
		ok bool
		f  Handler
	)

	if f, ok = r.h[identifier]; !ok {
		return nil, fmt.Errorf("could not get element from registry: %s", identifier)
	}

	return f, nil
}

func (r *registry) All() (list functionMetaList) {
	for _, handler := range r.h {
		m := handler.Meta(&types.Function{})
		list = append(list, &m)
	}

	return
}

func (r *registry) Preload() {
	r.Add("verifierQueryParam", verifierQueryParam{})
	r.Add("verifierOrigin", verifierOrigin{})
	r.Add("expediterRedirection", expediterRedirection{})
	r.Add("processerWorkflow", processerWorkflow{})
}
