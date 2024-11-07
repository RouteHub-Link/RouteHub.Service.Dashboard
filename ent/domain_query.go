// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	entdomain "RouteHub.Service.Dashboard/ent/domain"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/predicate"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainQuery is the builder for querying Domain entities.
type DomainQuery struct {
	config
	ctx              *QueryContext
	order            []entdomain.OrderOption
	inters           []Interceptor
	predicates       []predicate.Domain
	withOrganization *OrganizationQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DomainQuery builder.
func (dq *DomainQuery) Where(ps ...predicate.Domain) *DomainQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DomainQuery) Limit(limit int) *DomainQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DomainQuery) Offset(offset int) *DomainQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DomainQuery) Unique(unique bool) *DomainQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DomainQuery) Order(o ...entdomain.OrderOption) *DomainQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryOrganization chains the current query on the "organization" edge.
func (dq *DomainQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entdomain.Table, entdomain.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entdomain.OrganizationTable, entdomain.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Domain entity from the query.
// Returns a *NotFoundError when no Domain was found.
func (dq *DomainQuery) First(ctx context.Context) (*Domain, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entdomain.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DomainQuery) FirstX(ctx context.Context) *Domain {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Domain ID from the query.
// Returns a *NotFoundError when no Domain ID was found.
func (dq *DomainQuery) FirstID(ctx context.Context) (id mixin.ID, err error) {
	var ids []mixin.ID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entdomain.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DomainQuery) FirstIDX(ctx context.Context) mixin.ID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Domain entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Domain entity is found.
// Returns a *NotFoundError when no Domain entities are found.
func (dq *DomainQuery) Only(ctx context.Context) (*Domain, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entdomain.Label}
	default:
		return nil, &NotSingularError{entdomain.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DomainQuery) OnlyX(ctx context.Context) *Domain {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Domain ID in the query.
// Returns a *NotSingularError when more than one Domain ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DomainQuery) OnlyID(ctx context.Context) (id mixin.ID, err error) {
	var ids []mixin.ID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entdomain.Label}
	default:
		err = &NotSingularError{entdomain.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DomainQuery) OnlyIDX(ctx context.Context) mixin.ID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Domains.
func (dq *DomainQuery) All(ctx context.Context) ([]*Domain, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Domain, *DomainQuery]()
	return withInterceptors[[]*Domain](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DomainQuery) AllX(ctx context.Context) []*Domain {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Domain IDs.
func (dq *DomainQuery) IDs(ctx context.Context) (ids []mixin.ID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(entdomain.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DomainQuery) IDsX(ctx context.Context) []mixin.ID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DomainQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DomainQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DomainQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DomainQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DomainQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DomainQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DomainQuery) Clone() *DomainQuery {
	if dq == nil {
		return nil
	}
	return &DomainQuery{
		config:           dq.config,
		ctx:              dq.ctx.Clone(),
		order:            append([]entdomain.OrderOption{}, dq.order...),
		inters:           append([]Interceptor{}, dq.inters...),
		predicates:       append([]predicate.Domain{}, dq.predicates...),
		withOrganization: dq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DomainQuery) WithOrganization(opts ...func(*OrganizationQuery)) *DomainQuery {
	query := (&OrganizationClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withOrganization = query
	return dq
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
//	client.Domain.Query().
//		GroupBy(entdomain.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DomainQuery) GroupBy(field string, fields ...string) *DomainGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DomainGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = entdomain.Label
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
//	client.Domain.Query().
//		Select(entdomain.FieldName).
//		Scan(ctx, &v)
func (dq *DomainQuery) Select(fields ...string) *DomainSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DomainSelect{DomainQuery: dq}
	sbuild.label = entdomain.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DomainSelect configured with the given aggregations.
func (dq *DomainQuery) Aggregate(fns ...AggregateFunc) *DomainSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DomainQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !entdomain.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DomainQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Domain, error) {
	var (
		nodes       = []*Domain{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withOrganization != nil,
		}
	)
	if dq.withOrganization != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, entdomain.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Domain).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Domain{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withOrganization; query != nil {
		if err := dq.loadOrganization(ctx, query, nodes, nil,
			func(n *Domain, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DomainQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Domain, init func(*Domain), assign func(*Domain, *Organization)) error {
	ids := make([]mixin.ID, 0, len(nodes))
	nodeids := make(map[mixin.ID][]*Domain)
	for i := range nodes {
		if nodes[i].organization_id == nil {
			continue
		}
		fk := *nodes[i].organization_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DomainQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DomainQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entdomain.Table, entdomain.Columns, sqlgraph.NewFieldSpec(entdomain.FieldID, field.TypeString))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entdomain.FieldID)
		for i := range fields {
			if fields[i] != entdomain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DomainQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(entdomain.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = entdomain.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DomainGroupBy is the group-by builder for Domain entities.
type DomainGroupBy struct {
	selector
	build *DomainQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DomainGroupBy) Aggregate(fns ...AggregateFunc) *DomainGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DomainGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DomainQuery, *DomainGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DomainGroupBy) sqlScan(ctx context.Context, root *DomainQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DomainSelect is the builder for selecting fields of Domain entities.
type DomainSelect struct {
	*DomainQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DomainSelect) Aggregate(fns ...AggregateFunc) *DomainSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DomainSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DomainQuery, *DomainSelect](ctx, ds.DomainQuery, ds, ds.inters, v)
}

func (ds *DomainSelect) sqlScan(ctx context.Context, root *DomainQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
