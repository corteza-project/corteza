package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/system/types"
)

func (s Store) convertApigwFunctionFilter(f types.FunctionFilter) (query squirrel.SelectBuilder, err error) {
	query = s.apigwFunctionsSelectBuilder()
	query = filter.StateCondition(query, "af.deleted_at", f.Deleted)
	return
}
