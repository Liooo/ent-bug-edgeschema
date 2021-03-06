// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/cell"
	"entgo.io/bug/ent/datafield"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/record"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// CellQuery is the builder for querying Cell entities.
type CellQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Cell
	// eager-loading edges.
	withRecord    *RecordQuery
	withDataField *DataFieldQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CellQuery builder.
func (cq *CellQuery) Where(ps ...predicate.Cell) *CellQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CellQuery) Limit(limit int) *CellQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CellQuery) Offset(offset int) *CellQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CellQuery) Unique(unique bool) *CellQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CellQuery) Order(o ...OrderFunc) *CellQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryRecord chains the current query on the "record" edge.
func (cq *CellQuery) QueryRecord() *RecordQuery {
	query := &RecordQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cell.Table, cell.RecordColumn, selector),
			sqlgraph.To(record.Table, record.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cell.RecordTable, cell.RecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDataField chains the current query on the "data_field" edge.
func (cq *CellQuery) QueryDataField() *DataFieldQuery {
	query := &DataFieldQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cell.Table, cell.DataFieldColumn, selector),
			sqlgraph.To(datafield.Table, datafield.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cell.DataFieldTable, cell.DataFieldColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cell entity from the query.
// Returns a *NotFoundError when no Cell was found.
func (cq *CellQuery) First(ctx context.Context) (*Cell, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cell.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CellQuery) FirstX(ctx context.Context) *Cell {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Cell entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cell entity is found.
// Returns a *NotFoundError when no Cell entities are found.
func (cq *CellQuery) Only(ctx context.Context) (*Cell, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cell.Label}
	default:
		return nil, &NotSingularError{cell.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CellQuery) OnlyX(ctx context.Context) *Cell {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Cells.
func (cq *CellQuery) All(ctx context.Context) ([]*Cell, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CellQuery) AllX(ctx context.Context) []*Cell {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (cq *CellQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CellQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CellQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CellQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CellQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CellQuery) Clone() *CellQuery {
	if cq == nil {
		return nil
	}
	return &CellQuery{
		config:        cq.config,
		limit:         cq.limit,
		offset:        cq.offset,
		order:         append([]OrderFunc{}, cq.order...),
		predicates:    append([]predicate.Cell{}, cq.predicates...),
		withRecord:    cq.withRecord.Clone(),
		withDataField: cq.withDataField.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithRecord tells the query-builder to eager-load the nodes that are connected to
// the "record" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CellQuery) WithRecord(opts ...func(*RecordQuery)) *CellQuery {
	query := &RecordQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withRecord = query
	return cq
}

// WithDataField tells the query-builder to eager-load the nodes that are connected to
// the "data_field" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CellQuery) WithDataField(opts ...func(*DataFieldQuery)) *CellQuery {
	query := &DataFieldQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withDataField = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DataFieldID uuid.UUID `json:"data_field_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Cell.Query().
//		GroupBy(cell.FieldDataFieldID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CellQuery) GroupBy(field string, fields ...string) *CellGroupBy {
	grbuild := &CellGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = cell.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DataFieldID uuid.UUID `json:"data_field_id,omitempty"`
//	}
//
//	client.Cell.Query().
//		Select(cell.FieldDataFieldID).
//		Scan(ctx, &v)
//
func (cq *CellQuery) Select(fields ...string) *CellSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &CellSelect{CellQuery: cq}
	selbuild.label = cell.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *CellQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !cell.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CellQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cell, error) {
	var (
		nodes       = []*Cell{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withRecord != nil,
			cq.withDataField != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Cell).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Cell{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cq.withRecord; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Cell)
		for i := range nodes {
			fk := nodes[i].RecordID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(record.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "record_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Record = n
			}
		}
	}

	if query := cq.withDataField; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Cell)
		for i := range nodes {
			fk := nodes[i].DataFieldID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(datafield.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "data_field_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.DataField = n
			}
		}
	}

	return nodes, nil
}

func (cq *CellQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CellQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cq *CellQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cell.Table,
			Columns: cell.Columns,
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CellQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cell.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = cell.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CellGroupBy is the group-by builder for Cell entities.
type CellGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CellGroupBy) Aggregate(fns ...AggregateFunc) *CellGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CellGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *CellGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !cell.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CellGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// CellSelect is the builder for selecting fields of Cell entities.
type CellSelect struct {
	*CellQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CellSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CellQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *CellSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
