// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/goods"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// GoodsQuery is the builder for querying Goods entities.
type GoodsQuery struct {
	config
	ctx        *QueryContext
	order      []goods.OrderOption
	inters     []Interceptor
	predicates []predicate.Goods
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodsQuery builder.
func (_q *GoodsQuery) Where(ps ...predicate.Goods) *GoodsQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *GoodsQuery) Limit(limit int) *GoodsQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *GoodsQuery) Offset(offset int) *GoodsQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *GoodsQuery) Unique(unique bool) *GoodsQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *GoodsQuery) Order(o ...goods.OrderOption) *GoodsQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first Goods entity from the query.
// Returns a *NotFoundError when no Goods was found.
func (_q *GoodsQuery) First(ctx context.Context) (*Goods, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goods.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *GoodsQuery) FirstX(ctx context.Context) *Goods {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Goods ID from the query.
// Returns a *NotFoundError when no Goods ID was found.
func (_q *GoodsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goods.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *GoodsQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Goods entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Goods entity is found.
// Returns a *NotFoundError when no Goods entities are found.
func (_q *GoodsQuery) Only(ctx context.Context) (*Goods, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goods.Label}
	default:
		return nil, &NotSingularError{goods.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *GoodsQuery) OnlyX(ctx context.Context) *Goods {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Goods ID in the query.
// Returns a *NotSingularError when more than one Goods ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *GoodsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = &NotSingularError{goods.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *GoodsQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodsSlice.
func (_q *GoodsQuery) All(ctx context.Context) ([]*Goods, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Goods, *GoodsQuery]()
	return withInterceptors[[]*Goods](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *GoodsQuery) AllX(ctx context.Context) []*Goods {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Goods IDs.
func (_q *GoodsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(goods.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *GoodsQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *GoodsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*GoodsQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *GoodsQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *GoodsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *GoodsQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *GoodsQuery) Clone() *GoodsQuery {
	if _q == nil {
		return nil
	}
	return &GoodsQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]goods.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Goods{}, _q.predicates...),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (_q *GoodsQuery) GroupBy(field string, fields ...string) *GoodsGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoodsGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = goods.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (_q *GoodsQuery) Select(fields ...string) *GoodsSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &GoodsSelect{GoodsQuery: _q}
	sbuild.label = goods.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoodsSelect configured with the given aggregations.
func (_q *GoodsQuery) Aggregate(fns ...AggregateFunc) *GoodsSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *GoodsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !goods.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *GoodsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Goods, error) {
	var (
		nodes = []*Goods{}
		_spec = _q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Goods).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Goods{config: _q.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (_q *GoodsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *GoodsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(goods.Table, goods.Columns, sqlgraph.NewFieldSpec(goods.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goods.FieldID)
		for i := range fields {
			if fields[i] != goods.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *GoodsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(goods.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = goods.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range _q.modifiers {
		m(selector)
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (_q *GoodsQuery) ForUpdate(opts ...sql.LockOption) *GoodsQuery {
	if _q.driver.Dialect() == dialect.Postgres {
		_q.Unique(false)
	}
	_q.modifiers = append(_q.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return _q
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (_q *GoodsQuery) ForShare(opts ...sql.LockOption) *GoodsQuery {
	if _q.driver.Dialect() == dialect.Postgres {
		_q.Unique(false)
	}
	_q.modifiers = append(_q.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return _q
}

// Modify adds a query modifier for attaching custom logic to queries.
func (_q *GoodsQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodsSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// GoodsGroupBy is the group-by builder for Goods entities.
type GoodsGroupBy struct {
	selector
	build *GoodsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GoodsGroupBy) Aggregate(fns ...AggregateFunc) *GoodsGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GoodsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, ent.OpQueryGroupBy)
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodsQuery, *GoodsGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GoodsGroupBy) sqlScan(ctx context.Context, root *GoodsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoodsSelect is the builder for selecting fields of Goods entities.
type GoodsSelect struct {
	*GoodsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GoodsSelect) Aggregate(fns ...AggregateFunc) *GoodsSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GoodsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, ent.OpQuerySelect)
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodsQuery, *GoodsSelect](ctx, gs.GoodsQuery, gs, gs.inters, v)
}

func (gs *GoodsSelect) sqlScan(ctx context.Context, root *GoodsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GoodsSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodsSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}
