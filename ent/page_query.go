// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/page"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PageQuery is the builder for querying Page entities.
type PageQuery struct {
	config
	ctx        *QueryContext
	order      []page.OrderOption
	inters     []Interceptor
	predicates []predicate.Page
	withHub    *HubQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PageQuery builder.
func (pq *PageQuery) Where(ps ...predicate.Page) *PageQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PageQuery) Limit(limit int) *PageQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PageQuery) Offset(offset int) *PageQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PageQuery) Unique(unique bool) *PageQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PageQuery) Order(o ...page.OrderOption) *PageQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryHub chains the current query on the "hub" edge.
func (pq *PageQuery) QueryHub() *HubQuery {
	query := (&HubClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(page.Table, page.FieldID, selector),
			sqlgraph.To(hub.Table, hub.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, page.HubTable, page.HubColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Page entity from the query.
// Returns a *NotFoundError when no Page was found.
func (pq *PageQuery) First(ctx context.Context) (*Page, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{page.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PageQuery) FirstX(ctx context.Context) *Page {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Page ID from the query.
// Returns a *NotFoundError when no Page ID was found.
func (pq *PageQuery) FirstID(ctx context.Context) (id mixin.ID, err error) {
	var ids []mixin.ID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{page.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PageQuery) FirstIDX(ctx context.Context) mixin.ID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Page entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Page entity is found.
// Returns a *NotFoundError when no Page entities are found.
func (pq *PageQuery) Only(ctx context.Context) (*Page, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{page.Label}
	default:
		return nil, &NotSingularError{page.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PageQuery) OnlyX(ctx context.Context) *Page {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Page ID in the query.
// Returns a *NotSingularError when more than one Page ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PageQuery) OnlyID(ctx context.Context) (id mixin.ID, err error) {
	var ids []mixin.ID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{page.Label}
	default:
		err = &NotSingularError{page.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PageQuery) OnlyIDX(ctx context.Context) mixin.ID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pages.
func (pq *PageQuery) All(ctx context.Context) ([]*Page, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Page, *PageQuery]()
	return withInterceptors[[]*Page](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PageQuery) AllX(ctx context.Context) []*Page {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Page IDs.
func (pq *PageQuery) IDs(ctx context.Context) (ids []mixin.ID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(page.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PageQuery) IDsX(ctx context.Context) []mixin.ID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PageQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PageQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PageQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PageQuery) Clone() *PageQuery {
	if pq == nil {
		return nil
	}
	return &PageQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]page.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Page{}, pq.predicates...),
		withHub:    pq.withHub.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithHub tells the query-builder to eager-load the nodes that are connected to
// the "hub" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PageQuery) WithHub(opts ...func(*HubQuery)) *PageQuery {
	query := (&HubClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withHub = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Page.Query().
//		GroupBy(page.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PageQuery) GroupBy(field string, fields ...string) *PageGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PageGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = page.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Page.Query().
//		Select(page.FieldName).
//		Scan(ctx, &v)
func (pq *PageQuery) Select(fields ...string) *PageSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PageSelect{PageQuery: pq}
	sbuild.label = page.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PageSelect configured with the given aggregations.
func (pq *PageQuery) Aggregate(fns ...AggregateFunc) *PageSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !page.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Page, error) {
	var (
		nodes       = []*Page{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withHub != nil,
		}
	)
	if pq.withHub != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, page.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Page).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Page{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withHub; query != nil {
		if err := pq.loadHub(ctx, query, nodes, nil,
			func(n *Page, e *Hub) { n.Edges.Hub = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PageQuery) loadHub(ctx context.Context, query *HubQuery, nodes []*Page, init func(*Page), assign func(*Page, *Hub)) error {
	ids := make([]mixin.ID, 0, len(nodes))
	nodeids := make(map[mixin.ID][]*Page)
	for i := range nodes {
		if nodes[i].hub_fk == nil {
			continue
		}
		fk := *nodes[i].hub_fk
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(hub.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "hub_fk" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for i := range fields {
			if fields[i] != page.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(page.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = page.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PageGroupBy is the group-by builder for Page entities.
type PageGroupBy struct {
	selector
	build *PageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PageGroupBy) Aggregate(fns ...AggregateFunc) *PageGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PageQuery, *PageGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PageGroupBy) sqlScan(ctx context.Context, root *PageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PageSelect is the builder for selecting fields of Page entities.
type PageSelect struct {
	*PageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PageSelect) Aggregate(fns ...AggregateFunc) *PageSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PageQuery, *PageSelect](ctx, ps.PageQuery, ps, ps.inters, v)
}

func (ps *PageSelect) sqlScan(ctx context.Context, root *PageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
