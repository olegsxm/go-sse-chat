// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/olegsxm/go-sse-chat/ent/conversation"
	"github.com/olegsxm/go-sse-chat/ent/message"
	"github.com/olegsxm/go-sse-chat/ent/predicate"
	"github.com/olegsxm/go-sse-chat/ent/user"
)

// ConversationQuery is the builder for querying Conversation entities.
type ConversationQuery struct {
	config
	ctx          *QueryContext
	order        []conversation.OrderOption
	inters       []Interceptor
	predicates   []predicate.Conversation
	withUser     *UserQuery
	withMessages *MessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConversationQuery builder.
func (cq *ConversationQuery) Where(ps ...predicate.Conversation) *ConversationQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConversationQuery) Limit(limit int) *ConversationQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConversationQuery) Offset(offset int) *ConversationQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConversationQuery) Unique(unique bool) *ConversationQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConversationQuery) Order(o ...conversation.OrderOption) *ConversationQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryUser chains the current query on the "user" edge.
func (cq *ConversationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(conversation.Table, conversation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, conversation.UserTable, conversation.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMessages chains the current query on the "messages" edge.
func (cq *ConversationQuery) QueryMessages() *MessageQuery {
	query := (&MessageClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(conversation.Table, conversation.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, conversation.MessagesTable, conversation.MessagesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Conversation entity from the query.
// Returns a *NotFoundError when no Conversation was found.
func (cq *ConversationQuery) First(ctx context.Context) (*Conversation, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{conversation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConversationQuery) FirstX(ctx context.Context) *Conversation {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Conversation ID from the query.
// Returns a *NotFoundError when no Conversation ID was found.
func (cq *ConversationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{conversation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConversationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Conversation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Conversation entity is found.
// Returns a *NotFoundError when no Conversation entities are found.
func (cq *ConversationQuery) Only(ctx context.Context) (*Conversation, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{conversation.Label}
	default:
		return nil, &NotSingularError{conversation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConversationQuery) OnlyX(ctx context.Context) *Conversation {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Conversation ID in the query.
// Returns a *NotSingularError when more than one Conversation ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConversationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{conversation.Label}
	default:
		err = &NotSingularError{conversation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConversationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Conversations.
func (cq *ConversationQuery) All(ctx context.Context) ([]*Conversation, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Conversation, *ConversationQuery]()
	return withInterceptors[[]*Conversation](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConversationQuery) AllX(ctx context.Context) []*Conversation {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Conversation IDs.
func (cq *ConversationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(conversation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConversationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConversationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConversationQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConversationQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConversationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConversationQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConversationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConversationQuery) Clone() *ConversationQuery {
	if cq == nil {
		return nil
	}
	return &ConversationQuery{
		config:       cq.config,
		ctx:          cq.ctx.Clone(),
		order:        append([]conversation.OrderOption{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Conversation{}, cq.predicates...),
		withUser:     cq.withUser.Clone(),
		withMessages: cq.withMessages.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConversationQuery) WithUser(opts ...func(*UserQuery)) *ConversationQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withUser = query
	return cq
}

// WithMessages tells the query-builder to eager-load the nodes that are connected to
// the "messages" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConversationQuery) WithMessages(opts ...func(*MessageQuery)) *ConversationQuery {
	query := (&MessageClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withMessages = query
	return cq
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
//	client.Conversation.Query().
//		GroupBy(conversation.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConversationQuery) GroupBy(field string, fields ...string) *ConversationGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConversationGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = conversation.Label
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
//	client.Conversation.Query().
//		Select(conversation.FieldName).
//		Scan(ctx, &v)
func (cq *ConversationQuery) Select(fields ...string) *ConversationSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConversationSelect{ConversationQuery: cq}
	sbuild.label = conversation.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConversationSelect configured with the given aggregations.
func (cq *ConversationQuery) Aggregate(fns ...AggregateFunc) *ConversationSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConversationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !conversation.ValidColumn(f) {
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

func (cq *ConversationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Conversation, error) {
	var (
		nodes       = []*Conversation{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withUser != nil,
			cq.withMessages != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Conversation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Conversation{config: cq.config}
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
	if query := cq.withUser; query != nil {
		if err := cq.loadUser(ctx, query, nodes,
			func(n *Conversation) { n.Edges.User = []*User{} },
			func(n *Conversation, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withMessages; query != nil {
		if err := cq.loadMessages(ctx, query, nodes,
			func(n *Conversation) { n.Edges.Messages = []*Message{} },
			func(n *Conversation, e *Message) { n.Edges.Messages = append(n.Edges.Messages, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConversationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Conversation, init func(*Conversation), assign func(*Conversation, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Conversation)
	nids := make(map[uuid.UUID]map[*Conversation]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(conversation.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(conversation.UserPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(conversation.UserPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(conversation.UserPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Conversation]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *ConversationQuery) loadMessages(ctx context.Context, query *MessageQuery, nodes []*Conversation, init func(*Conversation), assign func(*Conversation, *Message)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Conversation)
	nids := make(map[uuid.UUID]map[*Conversation]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(conversation.MessagesTable)
		s.Join(joinT).On(s.C(message.FieldID), joinT.C(conversation.MessagesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(conversation.MessagesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(conversation.MessagesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Conversation]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Message](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "messages" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ConversationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConversationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(conversation.Table, conversation.Columns, sqlgraph.NewFieldSpec(conversation.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversation.FieldID)
		for i := range fields {
			if fields[i] != conversation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
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

func (cq *ConversationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(conversation.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = conversation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConversationGroupBy is the group-by builder for Conversation entities.
type ConversationGroupBy struct {
	selector
	build *ConversationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConversationGroupBy) Aggregate(fns ...AggregateFunc) *ConversationGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConversationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversationQuery, *ConversationGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConversationGroupBy) sqlScan(ctx context.Context, root *ConversationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConversationSelect is the builder for selecting fields of Conversation entities.
type ConversationSelect struct {
	*ConversationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConversationSelect) Aggregate(fns ...AggregateFunc) *ConversationSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConversationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConversationQuery, *ConversationSelect](ctx, cs.ConversationQuery, cs, cs.inters, v)
}

func (cs *ConversationSelect) sqlScan(ctx context.Context, root *ConversationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
