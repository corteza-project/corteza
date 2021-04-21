package rest

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/discovery/rest/request"
)

type (
	feed struct {
	}
)

func Feed() *feed {
	return &feed{}
}

func (ctrl feed) Changes(ctx context.Context, r *request.FeedChanges) (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}
