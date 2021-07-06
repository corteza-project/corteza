package apigw

import (
	"fmt"
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

func (r *registry) All() map[string]Handler {
	return r.h
}

func (r *registry) Preload() {
	r.Add("verifierQueryParam", verifierQueryParam{})
	r.Add("verifierOrigin", verifierOrigin{})
	r.Add("expediterRedirection", expediterRedirection{})
	r.Add("processerWorkflow", processerWorkflow{})
}
