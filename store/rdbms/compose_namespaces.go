package rdbms

import (
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"strings"
)

func (s Store) convertComposeNamespaceFilter(f types.NamespaceFilter) (query squirrel.SelectBuilder, err error) {
	query = s.composeNamespacesSelectBuilder()

	query = filter.StateCondition(query, "cns.deleted_at", f.Deleted)

	if len(f.NamespaceID) > 0 {
		query = query.Where(squirrel.Eq{"cns.id": f.NamespaceID})
	}

	if len(f.LabeledIDs) > 0 {
		query = query.Where(squirrel.Eq{"cns.id": f.LabeledIDs})
	}

	if f.Query != "" {
		q := "%" + strings.ToLower(f.Query) + "%"
		query = query.Where(squirrel.Or{
			squirrel.Like{"LOWER(cns.name)": q},
			squirrel.Like{"LOWER(cns.slug)": q},
		})
	}

	if f.Name != "" {
		query = query.Where(squirrel.Eq{"LOWER(cns.name)": strings.ToLower(f.Name)})
	}

	if f.Slug != "" {
		query = query.Where(squirrel.Eq{"LOWER(cns.slug)": strings.ToLower(f.Slug)})
	}

	return
}
