package apigw

import (
	"context"
	"net/http"
	"sort"
)

type (
	Execer interface {
		Exec(context.Context, *scp, map[string]interface{}) error
	}

	Sorter interface {
		Weight() int
	}

	ErrorHandler interface {
		Exec(context.Context, *scp, error)
	}

	Payload struct {
		params map[string]interface{}
		worker Worker
	}

	Worker interface {
		Execer
		Sorter
	}

	workers []Payload

	pl struct {
		w   workers
		err ErrorHandler
	}

	scp map[string]interface{}
)

func (s scp) Request() *http.Request {
	if _, ok := s["request"]; ok {
		return s["request"].(*http.Request)
	}

	return nil
}

func (s scp) Writer() http.ResponseWriter {
	if _, ok := s["writer"]; ok {
		return s["writer"].(http.ResponseWriter)
	}

	return nil
}

func (s scp) Set(k string, v interface{}) {
	s[k] = v
}

// Exec takes care of error handling and main
// functionality that takes place in worker
func (pp *pl) Exec(ctx context.Context, scope *scp) (err error) {
	for _, w := range pp.w {
		err = w.worker.Exec(ctx, scope, w.params)

		if err != nil {
			// call the error handler
			pp.err.Exec(ctx, scope, err)
			return
		}
	}

	return
}

// Add registers a new worker with parameters
// fethed from store
func (pp *pl) Add(ff Worker, p map[string]interface{}) {
	pp.w = append(pp.w, Payload{worker: ff, params: p})
	sort.Sort(pp.w)
}

// add error handler
func (pp *pl) ErrorHandler(ff ErrorHandler) {
	pp.err = ff
}

func (a workers) Len() int { return len(a) }
func (a workers) Less(i, j int) bool {
	return a[i].worker.Weight() < a[j].worker.Weight()
}
func (a workers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
