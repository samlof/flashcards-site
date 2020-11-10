// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/predicate"
	"flashcards-backend/ent/word"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// WordQuery is the builder for querying Word entities.
type WordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Word
	// eager-loading edges.
	withCardLogs      *CardLogQuery
	withCardSchedules *CardScheduleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wq *WordQuery) Where(ps ...predicate.Word) *WordQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WordQuery) Limit(limit int) *WordQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WordQuery) Offset(offset int) *WordQuery {
	wq.offset = &offset
	return wq
}

// Order adds an order step to the query.
func (wq *WordQuery) Order(o ...OrderFunc) *WordQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryCardLogs chains the current query on the cardLogs edge.
func (wq *WordQuery) QueryCardLogs() *CardLogQuery {
	query := &CardLogQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(cardlog.Table, cardlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, word.CardLogsTable, word.CardLogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCardSchedules chains the current query on the cardSchedules edge.
func (wq *WordQuery) QueryCardSchedules() *CardScheduleQuery {
	query := &CardScheduleQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(cardschedule.Table, cardschedule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, word.CardSchedulesTable, word.CardSchedulesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Word entity in the query. Returns *NotFoundError when no word was found.
func (wq *WordQuery) First(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{word.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WordQuery) FirstX(ctx context.Context) *Word {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Word id in the query. Returns *NotFoundError when no id was found.
func (wq *WordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{word.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WordQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Word entity in the query, returns an error if not exactly one entity was returned.
func (wq *WordQuery) Only(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{word.Label}
	default:
		return nil, &NotSingularError{word.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WordQuery) OnlyX(ctx context.Context) *Word {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Word id in the query, returns an error if not exactly one id was returned.
func (wq *WordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = &NotSingularError{word.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WordQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Words.
func (wq *WordQuery) All(ctx context.Context) ([]*Word, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WordQuery) AllX(ctx context.Context) []*Word {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Word ids.
func (wq *WordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wq.Select(word.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WordQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WordQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WordQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WordQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WordQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WordQuery) Clone() *WordQuery {
	if wq == nil {
		return nil
	}
	return &WordQuery{
		config:            wq.config,
		limit:             wq.limit,
		offset:            wq.offset,
		order:             append([]OrderFunc{}, wq.order...),
		unique:            append([]string{}, wq.unique...),
		predicates:        append([]predicate.Word{}, wq.predicates...),
		withCardLogs:      wq.withCardLogs.Clone(),
		withCardSchedules: wq.withCardSchedules.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

//  WithCardLogs tells the query-builder to eager-loads the nodes that are connected to
// the "cardLogs" edge. The optional arguments used to configure the query builder of the edge.
func (wq *WordQuery) WithCardLogs(opts ...func(*CardLogQuery)) *WordQuery {
	query := &CardLogQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withCardLogs = query
	return wq
}

//  WithCardSchedules tells the query-builder to eager-loads the nodes that are connected to
// the "cardSchedules" edge. The optional arguments used to configure the query builder of the edge.
func (wq *WordQuery) WithCardSchedules(opts ...func(*CardScheduleQuery)) *WordQuery {
	query := &CardScheduleQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withCardSchedules = query
	return wq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Word.Query().
//		GroupBy(word.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WordQuery) GroupBy(field string, fields ...string) *WordGroupBy {
	group := &WordGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Word.Query().
//		Select(word.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wq *WordQuery) Select(field string, fields ...string) *WordSelect {
	selector := &WordSelect{config: wq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(), nil
	}
	return selector
}

func (wq *WordQuery) prepareQuery(ctx context.Context) error {
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WordQuery) sqlAll(ctx context.Context) ([]*Word, error) {
	var (
		nodes       = []*Word{}
		_spec       = wq.querySpec()
		loadedTypes = [2]bool{
			wq.withCardLogs != nil,
			wq.withCardSchedules != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Word{config: wq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wq.withCardLogs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Word)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.CardLogs = []*CardLog{}
		}
		query.withFKs = true
		query.Where(predicate.CardLog(func(s *sql.Selector) {
			s.Where(sql.InValues(word.CardLogsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.card_log_card
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "card_log_card" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "card_log_card" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.CardLogs = append(node.Edges.CardLogs, n)
		}
	}

	if query := wq.withCardSchedules; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Word)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.CardSchedules = []*CardSchedule{}
		}
		query.withFKs = true
		query.Where(predicate.CardSchedule(func(s *sql.Selector) {
			s.Where(sql.InValues(word.CardSchedulesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.card_schedule_card
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "card_schedule_card" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "card_schedule_card" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.CardSchedules = append(node.Edges.CardSchedules, n)
		}
	}

	return nodes, nil
}

func (wq *WordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wq *WordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   word.Table,
			Columns: word.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, word.ValidColumn)
			}
		}
	}
	return _spec
}

func (wq *WordQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(word.Table)
	selector := builder.Select(t1.Columns(word.Columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(word.Columns...)...)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector, word.ValidColumn)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WordGroupBy is the builder for group-by Word entities.
type WordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WordGroupBy) Aggregate(fns ...AggregateFunc) *WordGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wgb *WordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *WordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *WordGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *WordGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *WordGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *WordGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *WordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *WordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *WordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wgb *WordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *WordGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *WordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !word.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WordGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql
	columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
	columns = append(columns, wgb.fields...)
	for _, fn := range wgb.fns {
		columns = append(columns, fn(selector, word.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(wgb.fields...)
}

// WordSelect is the builder for select fields of Word entities.
type WordSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ws *WordSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ws.path(ctx)
	if err != nil {
		return err
	}
	ws.sql = query
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *WordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *WordSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ws *WordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *WordSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *WordSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *WordSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *WordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *WordSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *WordSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ws *WordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *WordSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *WordSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ws.fields {
		if !word.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ws.sqlQuery().Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ws *WordSelect) sqlQuery() sql.Querier {
	selector := ws.sql
	selector.Select(selector.Columns(ws.fields...)...)
	return selector
}
