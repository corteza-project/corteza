package rest

import (
	"context"
	"github.com/cortezaproject/corteza-server/discovery/rest/internal/mapping"
	"github.com/cortezaproject/corteza-server/discovery/rest/request"
)

type (
	mappings struct {
		sys interface {
			Users(context.Context) ([]*mapping.Mapping, error)
		}

		cmp interface {
			Namespaces(context.Context) ([]*mapping.Mapping, error)
			Modules(context.Context) ([]*mapping.Mapping, error)
			Records(context.Context) ([]*mapping.Mapping, error)
		}
	}
)

func Mappings() *mappings {
	return &mappings{
		sys: mapping.SystemMapping(),
		cmp: mapping.ComposeMapping(),
	}
}

func (ctrl mappings) List(ctx context.Context, r *request.MappingsList) (interface{}, error) {
	var (
		out = make([]*mapping.Mapping, 0, 100)
		// Collection of all mapping functions we support
		//
		// Each function is responsible to
		mappingFn = make([]func(ctx context.Context) ([]*mapping.Mapping, error), 4)
	)

	mappingFn = append(mappingFn, ctrl.sys.Users)
	mappingFn = append(mappingFn, ctrl.cmp.Namespaces)
	mappingFn = append(mappingFn, ctrl.cmp.Modules)
	mappingFn = append(mappingFn, ctrl.cmp.Records)

	return out, func() error {
		for _, fn := range mappingFn {
			if fn == nil {
				continue
			}

			mm, err := fn(ctx)
			if err != nil {
				return err
			}

			out = append(out, mm...)
		}

		return nil
	}()
}
