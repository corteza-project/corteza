package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/apigw_function.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/store/rdbms/builders"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ = errors.Is

// SearchApigwFunctions returns all matching rows
//
// This function calls convertApigwFunctionFilter with the given
// types.FunctionFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchApigwFunctions(ctx context.Context, f types.FunctionFilter) (types.FunctionSet, types.FunctionFilter, error) {
	var (
		err error
		set []*types.Function
		q   squirrel.SelectBuilder
	)

	return set, f, func() error {
		q, err = s.convertApigwFunctionFilter(f)
		if err != nil {
			return err
		}

		// Paging enabled
		// {search: {enablePaging:true}}
		// Cleanup unwanted cursor values (only relevant is f.PageCursor, next&prev are reset and returned)
		f.PrevPage, f.NextPage = nil, nil

		if f.PageCursor != nil {
			// Page cursor exists so we need to validate it against used sort
			// To cover the case when paging cursor is set but sorting is empty, we collect the sorting instructions
			// from the cursor.
			// This (extracted sorting info) is then returned as part of response
			if f.Sort, err = f.PageCursor.Sort(f.Sort); err != nil {
				return err
			}
		}

		// Make sure results are always sorted at least by primary keys
		if f.Sort.Get("id") == nil {
			f.Sort = append(f.Sort, &filter.SortExpr{
				Column:     "id",
				Descending: f.Sort.LastDescending(),
			})
		}

		// Cloned sorting instructions for the actual sorting
		// Original are passed to the fetchFullPageOfUsers fn used for cursor creation so it MUST keep the initial
		// direction information
		sort := f.Sort.Clone()

		// When cursor for a previous page is used it's marked as reversed
		// This tells us to flip the descending flag on all used sort keys
		if f.PageCursor != nil && f.PageCursor.ROrder {
			sort.Reverse()
		}

		// Apply sorting expr from filter to query
		if q, err = setOrderBy(q, sort, s.sortableApigwFunctionColumns()); err != nil {
			return err
		}

		set, f.PrevPage, f.NextPage, err = s.fetchFullPageOfApigwFunctions(
			ctx,
			q, f.Sort, f.PageCursor,
			f.Limit,
			nil,
			func(cur *filter.PagingCursor) squirrel.Sqlizer {
				return builders.CursorCondition(cur, nil)
			},
		)

		if err != nil {
			return err
		}

		f.PageCursor = nil
		return nil
	}()
}

// fetchFullPageOfApigwFunctions collects all requested results.
//
// Function applies:
//  - cursor conditions (where ...)
//  - limit
//
// Main responsibility of this function is to perform additional sequential queries in case when not enough results
// are collected due to failed check on a specific row (by check fn).
//
// Function then moves cursor to the last item fetched
func (s Store) fetchFullPageOfApigwFunctions(
	ctx context.Context,
	q squirrel.SelectBuilder,
	sort filter.SortExprSet,
	cursor *filter.PagingCursor,
	reqItems uint,
	check func(*types.Function) (bool, error),
	cursorCond func(*filter.PagingCursor) squirrel.Sqlizer,
) (set []*types.Function, prev, next *filter.PagingCursor, err error) {
	var (
		aux []*types.Function

		// When cursor for a previous page is used it's marked as reversed
		// This tells us to flip the descending flag on all used sort keys
		reversedOrder = cursor != nil && cursor.ROrder

		// copy of the select builder
		tryQuery squirrel.SelectBuilder

		// Copy no. of required items to limit
		// Limit will change when doing subsequent queries to fill
		// the set with all required items
		limit = reqItems

		// cursor to prev. page is only calculated when cursor is used
		hasPrev = cursor != nil

		// next cursor is calculated when there are more pages to come
		hasNext bool
	)

	set = make([]*types.Function, 0, DefaultSliceCapacity)

	for try := 0; try < MaxRefetches; try++ {
		if cursor != nil {
			tryQuery = q.Where(cursorCond(cursor))
		} else {
			tryQuery = q
		}

		if limit > 0 {
			// fetching + 1 so we know if there are more items
			// we can fetch (next-page cursor)
			tryQuery = tryQuery.Limit(uint64(limit + 1))
		}

		if aux, err = s.QueryApigwFunctions(ctx, tryQuery, check); err != nil {
			return nil, nil, nil, err
		}

		if len(aux) == 0 {
			// nothing fetched
			break
		}

		// append fetched items
		set = append(set, aux...)

		if reqItems == 0 {
			// no max requested items specified, break out
			break
		}

		collected := uint(len(set))

		if reqItems > collected {
			// not enough items fetched, try again with adjusted limit
			limit = reqItems - collected

			if limit < MinEnsureFetchLimit {
				// In case limit is set very low and we've missed records in the first fetch,
				// make sure next fetch limit is a bit higher
				limit = MinEnsureFetchLimit
			}

			// Update cursor so that it points to the last item fetched
			cursor = s.collectApigwFunctionCursorValues(set[collected-1], sort...)

			// Copy reverse flag from sorting
			cursor.LThen = sort.Reversed()
			continue
		}

		if reqItems < collected {
			set = set[:reqItems]
			hasNext = true
		}

		break
	}

	collected := len(set)

	if collected == 0 {
		return nil, nil, nil, nil
	}

	if reversedOrder {
		// Fetched set needs to be reversed because we've forced a descending order to get the previous page
		for i, j := 0, collected-1; i < j; i, j = i+1, j-1 {
			set[i], set[j] = set[j], set[i]
		}

		// when in reverse-order rules on what cursor to return change
		hasPrev, hasNext = hasNext, hasPrev
	}

	if hasPrev {
		prev = s.collectApigwFunctionCursorValues(set[0], sort...)
		prev.ROrder = true
		prev.LThen = !sort.Reversed()
	}

	if hasNext {
		next = s.collectApigwFunctionCursorValues(set[collected-1], sort...)
		next.LThen = sort.Reversed()
	}

	return set, prev, next, nil
}

// QueryApigwFunctions queries the database, converts and checks each row and
// returns collected set
//
// Fn also returns total number of fetched items and last fetched item so that the caller can construct cursor
// for next page of results
func (s Store) QueryApigwFunctions(
	ctx context.Context,
	q squirrel.Sqlizer,
	check func(*types.Function) (bool, error),
) ([]*types.Function, error) {
	var (
		set = make([]*types.Function, 0, DefaultSliceCapacity)
		res *types.Function

		// Query rows with
		rows, err = s.Query(ctx, q)
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Err(); err == nil {
			res, err = s.internalApigwFunctionRowScanner(rows)
		}

		if err != nil {
			return nil, err
		}

		set = append(set, res)
	}

	return set, rows.Err()
}

// LookupApigwFunctionByID searches for function by ID
func (s Store) LookupApigwFunctionByID(ctx context.Context, id uint64) (*types.Function, error) {
	return s.execLookupApigwFunction(ctx, squirrel.Eq{
		s.preprocessColumn("af.id", ""): store.PreprocessValue(id, ""),
	})
}

// LookupApigwFunctionByRoute searches for function by route
func (s Store) LookupApigwFunctionByRoute(ctx context.Context, route string) (*types.Function, error) {
	return s.execLookupApigwFunction(ctx, squirrel.Eq{
		s.preprocessColumn("af.rel_route", ""): store.PreprocessValue(route, ""),
	})
}

// CreateApigwFunction creates one or more rows in apigw_functions table
func (s Store) CreateApigwFunction(ctx context.Context, rr ...*types.Function) (err error) {
	for _, res := range rr {
		err = s.checkApigwFunctionConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execCreateApigwFunctions(ctx, s.internalApigwFunctionEncoder(res))
		if err != nil {
			return err
		}
	}

	return
}

// UpdateApigwFunction updates one or more existing rows in apigw_functions
func (s Store) UpdateApigwFunction(ctx context.Context, rr ...*types.Function) error {
	return s.partialApigwFunctionUpdate(ctx, nil, rr...)
}

// partialApigwFunctionUpdate updates one or more existing rows in apigw_functions
func (s Store) partialApigwFunctionUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Function) (err error) {
	for _, res := range rr {
		err = s.checkApigwFunctionConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execUpdateApigwFunctions(
			ctx,
			squirrel.Eq{
				s.preprocessColumn("af.id", ""): store.PreprocessValue(res.ID, ""),
			},
			s.internalApigwFunctionEncoder(res).Skip("id").Only(onlyColumns...))
		if err != nil {
			return err
		}
	}

	return
}

// DeleteApigwFunction Deletes one or more rows from apigw_functions table
func (s Store) DeleteApigwFunction(ctx context.Context, rr ...*types.Function) (err error) {
	for _, res := range rr {

		err = s.execDeleteApigwFunctions(ctx, squirrel.Eq{
			s.preprocessColumn("af.id", ""): store.PreprocessValue(res.ID, ""),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteApigwFunctionByID Deletes row from the apigw_functions table
func (s Store) DeleteApigwFunctionByID(ctx context.Context, ID uint64) error {
	return s.execDeleteApigwFunctions(ctx, squirrel.Eq{
		s.preprocessColumn("af.id", ""): store.PreprocessValue(ID, ""),
	})
}

// TruncateApigwFunctions Deletes all rows from the apigw_functions table
func (s Store) TruncateApigwFunctions(ctx context.Context) error {
	return s.Truncate(ctx, s.apigwFunctionTable())
}

// execLookupApigwFunction prepares ApigwFunction query and executes it,
// returning types.Function (or error)
func (s Store) execLookupApigwFunction(ctx context.Context, cnd squirrel.Sqlizer) (res *types.Function, err error) {
	var (
		row rowScanner
	)

	row, err = s.QueryRow(ctx, s.apigwFunctionsSelectBuilder().Where(cnd))
	if err != nil {
		return
	}

	res, err = s.internalApigwFunctionRowScanner(row)
	if err != nil {
		return
	}

	return res, nil
}

// execCreateApigwFunctions updates all matched (by cnd) rows in apigw_functions with given data
func (s Store) execCreateApigwFunctions(ctx context.Context, payload store.Payload) error {
	return s.Exec(ctx, s.InsertBuilder(s.apigwFunctionTable()).SetMap(payload))
}

// execUpdateApigwFunctions updates all matched (by cnd) rows in apigw_functions with given data
func (s Store) execUpdateApigwFunctions(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return s.Exec(ctx, s.UpdateBuilder(s.apigwFunctionTable("af")).Where(cnd).SetMap(set))
}

// execDeleteApigwFunctions Deletes all matched (by cnd) rows in apigw_functions with given data
func (s Store) execDeleteApigwFunctions(ctx context.Context, cnd squirrel.Sqlizer) error {
	return s.Exec(ctx, s.DeleteBuilder(s.apigwFunctionTable("af")).Where(cnd))
}

func (s Store) internalApigwFunctionRowScanner(row rowScanner) (res *types.Function, err error) {
	res = &types.Function{}

	if _, has := s.config.RowScanners["apigwFunction"]; has {
		scanner := s.config.RowScanners["apigwFunction"].(func(_ rowScanner, _ *types.Function) error)
		err = scanner(row, res)
	} else {
		err = row.Scan(
			&res.ID,
			&res.Route,
			&res.Weight,
			&res.Kind,
			&res.Ref,
			&res.Params,
			&res.CreatedBy,
			&res.UpdatedBy,
			&res.DeletedBy,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound.Stack(1)
	}

	if err != nil {
		return nil, errors.Store("could not scan apigwFunction db row: %s", err).Wrap(err)
	} else {
		return res, nil
	}
}

// QueryApigwFunctions returns squirrel.SelectBuilder with set table and all columns
func (s Store) apigwFunctionsSelectBuilder() squirrel.SelectBuilder {
	return s.SelectBuilder(s.apigwFunctionTable("af"), s.apigwFunctionColumns("af")...)
}

// apigwFunctionTable name of the db table
func (Store) apigwFunctionTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "apigw_functions" + alias
}

// ApigwFunctionColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) apigwFunctionColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "rel_route",
		alias + "weight",
		alias + "kind",
		alias + "ref",
		alias + "params",
		alias + "created_by",
		alias + "updated_by",
		alias + "deleted_by",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// {true true false true true false}

// sortableApigwFunctionColumns returns all ApigwFunction columns flagged as sortable
//
// With optional string arg, all columns are returned aliased
func (Store) sortableApigwFunctionColumns() map[string]string {
	return map[string]string{
		"id": "id",
	}
}

// internalApigwFunctionEncoder encodes fields from types.Function to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeApigwFunction
// func when rdbms.customEncoder=true
func (s Store) internalApigwFunctionEncoder(res *types.Function) store.Payload {
	return store.Payload{
		"id":         res.ID,
		"rel_route":  res.Route,
		"weight":     res.Weight,
		"kind":       res.Kind,
		"ref":        res.Ref,
		"params":     res.Params,
		"created_by": res.CreatedBy,
		"updated_by": res.UpdatedBy,
		"deleted_by": res.DeletedBy,
		"created_at": res.CreatedAt,
		"updated_at": res.UpdatedAt,
		"deleted_at": res.DeletedAt,
	}
}

// collectApigwFunctionCursorValues collects values from the given resource that and sets them to the cursor
// to be used for pagination
//
// Values that are collected must come from sortable, unique or primary columns/fields
// At least one of the collected columns must be flagged as unique, otherwise fn appends primary keys at the end
//
// Known issue:
//   when collecting cursor values for query that sorts by unique column with partial index (ie: unique handle on
//   undeleted items)
func (s Store) collectApigwFunctionCursorValues(res *types.Function, cc ...*filter.SortExpr) *filter.PagingCursor {
	var (
		cursor = &filter.PagingCursor{LThen: filter.SortExprSet(cc).Reversed()}

		hasUnique bool

		// All known primary key columns

		pkId bool

		collect = func(cc ...*filter.SortExpr) {
			for _, c := range cc {
				switch c.Column {
				case "id":
					cursor.Set(c.Column, res.ID, c.Descending)

					pkId = true

				}
			}
		}
	)

	collect(cc...)
	if !hasUnique || !(pkId && true) {
		collect(&filter.SortExpr{Column: "id", Descending: false})
	}

	return cursor
}

// checkApigwFunctionConstraints performs lookups (on valid) resource to check if any of the values on unique fields
// already exists in the store
//
// Using built-in constraint checking would be more performant but unfortunately we cannot rely
// on the full support (MySQL does not support conditional indexes)
func (s *Store) checkApigwFunctionConstraints(ctx context.Context, res *types.Function) error {
	// Consider resource valid when all fields in unique constraint check lookups
	// have valid (non-empty) value
	//
	// Only string and uint64 are supported for now
	// feel free to add additional types if needed
	var valid = true

	if !valid {
		return nil
	}

	return nil
}
