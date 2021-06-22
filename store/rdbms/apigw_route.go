package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/system/types"
)

func (s Store) convertApigwRouteFilter(f types.RouteFilter) (query squirrel.SelectBuilder, err error) {
	query = s.apigwRoutesSelectBuilder()
	query = filter.StateCondition(query, "ar.deleted_at", f.Deleted)
	return
}
