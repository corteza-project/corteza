package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/system/types"
)

func (s Store) convertApigwFunctionFilter(f types.ApigwFunctionFilter) (query squirrel.SelectBuilder, err error) {
	query = s.apigwFunctionsSelectBuilder()

	query = filter.StateCondition(query, "af.deleted_at", f.Deleted)

	if f.RouteID > 0 {
		query = query.Where(squirrel.Eq{"af.rel_route": f.RouteID})
	}

	return
}
