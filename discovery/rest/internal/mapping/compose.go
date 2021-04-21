package mapping

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/service"
	"github.com/cortezaproject/corteza-server/compose/types"
	"strings"
)

type (
	namespaceSearcher interface {
		Find(ctx context.Context, filter types.NamespaceFilter) (types.NamespaceSet, types.NamespaceFilter, error)
	}

	moduleSearcher interface {
		Find(ctx context.Context, filter types.ModuleFilter) (types.ModuleSet, types.ModuleFilter, error)
	}

	composeAccessControl interface {
		CanReadRecordValue(ctx context.Context, r *types.ModuleField) bool
		CanSearchRecords(ctx context.Context, r *types.Module) bool
	}

	composeMapping struct {
		ac  composeAccessControl
		ns  namespaceSearcher
		mod moduleSearcher
	}
)

func ComposeMapping() *composeMapping {
	return &composeMapping{
		ac:  service.DefaultAccessControl,
		ns:  service.DefaultNamespace,
		mod: service.DefaultModule,
	}
}

func (m composeMapping) Namespaces(ctx context.Context) ([]*Mapping, error) {
	return []*Mapping{&Mapping{
		Index: "compose-namespaces",
		Mapping: map[string]*property{
			"resourceType": {Type: "keyword"},

			"namespaceID": {Type: "unsigned_long"},

			"name":    {Type: "keyword", Boost: 2},
			"handle":  {Type: "keyword", Boost: 2},
			"labels":  {Type: "keyword"},
			"enabled": {Type: "boolean"},

			"meta": {Type: "object", Properties: map[string]*property{
				"subtitle":    {Type: "keyword"},
				"description": {Type: "text"},
			}},

			"created": change(),
			"updated": change(),
			"deleted": change(),

			"security": security(),
		},
	}}, nil
}

func (m composeMapping) Modules(ctx context.Context) ([]*Mapping, error) {
	return []*Mapping{&Mapping{
		Index: fmt.Sprintf("compose-modules"),
		Mapping: map[string]*property{
			"resourceType": {Type: "keyword"},

			"moduleID": &property{Type: "unsigned_long"},
			"created":  change(),
			"updated":  change(),
			"deleted":  change(),

			"security": security(),

			"name":   {Type: "text", Boost: 2},
			"handle": {Type: "keyword", Boost: 2},
			"fields": {
				Type: "object",
				Properties: map[string]*property{
					"name":  {Type: "keyword"},
					"label": {Type: "text"},
				},
			},
		},
	}}, nil
}

func (m composeMapping) Records(ctx context.Context) ([]*Mapping, error) {
	var (
		nn, _, err = m.ns.Find(ctx, types.NamespaceFilter{})
		out        = make([]*Mapping, 0, len(nn)*50)
	)

	if err != nil {
		return nil, err
	}

	return out, nn.Walk(func(ns *types.Namespace) (err error) {
		// @todo do we index records in this namespace?
		var (
			mm types.ModuleSet
		)

		mm, _, err = m.mod.Find(ctx, types.ModuleFilter{NamespaceID: ns.ID})
		if err != nil {
			return err
		}

		// collect record mappings for all modules
		return mm.Walk(func(mod *types.Module) error {
			// @todo do we index records of module?

			if m := m.records(ctx, mod, mm); m != nil {
				out = append(out, m)
			}

			return nil
		})
	})
}

// Module mapping is always returned, even if it contains no properties
//
// If there are no properties indexer can assume
func (m composeMapping) records(ctx context.Context, mod *types.Module, mm types.ModuleSet) (mapping *Mapping) {
	//	@todo check if module can be Mappinged

	mapping = &Mapping{
		Index:   fmt.Sprintf("compose-records-%d-%d", mod.NamespaceID, mod.ID),
		Mapping: make(map[string]*property),
	}

	props := m.moduleFieldsMappingProperties(ctx, mod, mm)
	if len(props) == 0 {
		// no mapping
		return
	}
	mapping.Mapping["resourceType"] = &property{Type: "keyword"}

	mapping.Mapping["recordID"] = &property{Type: "unsigned_long"}
	mapping.Mapping["created"] = change()
	mapping.Mapping["updated"] = change()
	mapping.Mapping["deleted"] = change()
	mapping.Mapping["security"] = security()

	mapping.Mapping["values"] = &property{
		Type:       "nested",
		Properties: props,
	}

	// mapping for record is there, add fields for namespace & module
	mapping.Mapping["module"] = &property{
		Type: "nested",
		Properties: map[string]*property{
			"moduleID": &property{Type: "unsigned_long"},
			"name":     &property{Type: "text"},
			"handle":   &property{Type: "keyword"},
		},
	}
	mapping.Mapping["namespace"] = &property{
		Type: "nested",
		Properties: map[string]*property{
			"namespaceID": &property{Type: "unsigned_long"},
			"name":        &property{Type: "text"},
			"handle":      &property{Type: "keyword"},
		},
	}

	return
}

func (m composeMapping) moduleFieldsMappingProperties(ctx context.Context, mod *types.Module, mm types.ModuleSet) map[string]*property {
	mf := make(map[string]*property)

	for _, f := range mod.Fields {
		for k, m := range m.moduleFieldMappingProperties(ctx, f, mm) {
			mf[k] = m
		}
	}

	return mf
}

func (m composeMapping) moduleFieldMappingProperties(ctx context.Context, f *types.ModuleField, mm types.ModuleSet) map[string]*property {
	if !m.ac.CanReadRecordValue(ctx, f) {
		return nil
	}

	var (
		p    = &property{Type: "text"}
		name = f.Name
	)
	// @todo get indexing props from f.Options!

	if p.Type == "" {
		// Type not explicitly set, get it by converting
		switch strings.ToLower(f.Kind) {
		case "bool":
			p.Type = "boolean"

		case "datetime":
			p.Type = "date"
		// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
		// @todo format param!

		case "email":
			p.Type = "keyword"

		case "file":
		// @todo ref

		case "number":
			// @todo what number type to use?
			return nil

		case "record":
			//iopt := f.Options.Index()

			// @todo make sure we do not fall into recursive loop

			// if ifc.parents[]

			// @todo if single-value then object
			// https://www.elastic.co/guide/en/elasticsearch/reference/current/object.html
			// @todo if multi-value then nested
			// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html

			//var deepDive = 1
			//
			//if deepDive == 0 {
			//	// no
			//	return nil
			//}
			//
			//m := mm.FindByID(f.Options.Uint64("moduleID"))
			//if m == nil {
			//	return nil
			//}
			//
			//if opt.nestingLevel < 2 {
			//	return append(composeModuleFieldsMappingProperties(m, mm, Context{prefix: f.Name, nestingLevel: opt.nestingLevel + 1}), m)
			//}
			return nil

		case "select":
			p.Type = "keyword"

		case "string":
			p.Type = "text"

		case "url":
			p.Type = "keyword"

		case "user":
			// @todo ref
			return nil
		}
	}

	return map[string]*property{name: p}
}
