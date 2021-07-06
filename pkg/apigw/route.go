package apigw

import (
	"context"
	"net/http"
)

type (
	route struct {
		ID       uint64
		endpoint string
		method   string

		pipe *pl
	}
)

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var (
		ctx   = context.Background()
		scope = scp{
			req:    req,
			writer: w,
		}
	)

	err := r.pipe.Exec(ctx, &scope)

	if err != nil {
		// log error
	}
}
