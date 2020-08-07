// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/predicate"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CardLogUpdate is the builder for updating CardLog entities.
type CardLogUpdate struct {
	config
	hooks      []Hook
	mutation   *CardLogMutation
	predicates []predicate.CardLog
}

// Where adds a new predicate for the builder.
func (clu *CardLogUpdate) Where(ps ...predicate.CardLog) *CardLogUpdate {
	clu.predicates = append(clu.predicates, ps...)
	return clu
}

// SetReviewed sets the reviewed field.
func (clu *CardLogUpdate) SetReviewed(b bool) *CardLogUpdate {
	clu.mutation.SetReviewed(b)
	return clu
}

// SetNillableReviewed sets the reviewed field if the given value is not nil.
func (clu *CardLogUpdate) SetNillableReviewed(b *bool) *CardLogUpdate {
	if b != nil {
		clu.SetReviewed(*b)
	}
	return clu
}

// SetUserID sets the user edge to User by id.
func (clu *CardLogUpdate) SetUserID(id int) *CardLogUpdate {
	clu.mutation.SetUserID(id)
	return clu
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (clu *CardLogUpdate) SetNillableUserID(id *int) *CardLogUpdate {
	if id != nil {
		clu = clu.SetUserID(*id)
	}
	return clu
}

// SetUser sets the user edge to User.
func (clu *CardLogUpdate) SetUser(u *User) *CardLogUpdate {
	return clu.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (clu *CardLogUpdate) SetCardID(id int) *CardLogUpdate {
	clu.mutation.SetCardID(id)
	return clu
}

// SetCard sets the card edge to Word.
func (clu *CardLogUpdate) SetCard(w *Word) *CardLogUpdate {
	return clu.SetCardID(w.ID)
}

// Mutation returns the CardLogMutation object of the builder.
func (clu *CardLogUpdate) Mutation() *CardLogMutation {
	return clu.mutation
}

// ClearUser clears the user edge to User.
func (clu *CardLogUpdate) ClearUser() *CardLogUpdate {
	clu.mutation.ClearUser()
	return clu
}

// ClearCard clears the card edge to Word.
func (clu *CardLogUpdate) ClearCard() *CardLogUpdate {
	clu.mutation.ClearCard()
	return clu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (clu *CardLogUpdate) Save(ctx context.Context) (int, error) {

	if _, ok := clu.mutation.CardID(); clu.mutation.CardCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"card\"")
	}
	var (
		err      error
		affected int
	)
	if len(clu.hooks) == 0 {
		affected, err = clu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			clu.mutation = mutation
			affected, err = clu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(clu.hooks) - 1; i >= 0; i-- {
			mut = clu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (clu *CardLogUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *CardLogUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *CardLogUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (clu *CardLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardlog.Table,
			Columns: cardlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardlog.FieldID,
			},
		},
	}
	if ps := clu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.Reviewed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cardlog.FieldReviewed,
		})
	}
	if clu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardlog.UserTable,
			Columns: []string{cardlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardlog.UserTable,
			Columns: []string{cardlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if clu.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardlog.CardTable,
			Columns: []string{cardlog.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardlog.CardTable,
			Columns: []string{cardlog.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CardLogUpdateOne is the builder for updating a single CardLog entity.
type CardLogUpdateOne struct {
	config
	hooks    []Hook
	mutation *CardLogMutation
}

// SetReviewed sets the reviewed field.
func (cluo *CardLogUpdateOne) SetReviewed(b bool) *CardLogUpdateOne {
	cluo.mutation.SetReviewed(b)
	return cluo
}

// SetNillableReviewed sets the reviewed field if the given value is not nil.
func (cluo *CardLogUpdateOne) SetNillableReviewed(b *bool) *CardLogUpdateOne {
	if b != nil {
		cluo.SetReviewed(*b)
	}
	return cluo
}

// SetUserID sets the user edge to User by id.
func (cluo *CardLogUpdateOne) SetUserID(id int) *CardLogUpdateOne {
	cluo.mutation.SetUserID(id)
	return cluo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (cluo *CardLogUpdateOne) SetNillableUserID(id *int) *CardLogUpdateOne {
	if id != nil {
		cluo = cluo.SetUserID(*id)
	}
	return cluo
}

// SetUser sets the user edge to User.
func (cluo *CardLogUpdateOne) SetUser(u *User) *CardLogUpdateOne {
	return cluo.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (cluo *CardLogUpdateOne) SetCardID(id int) *CardLogUpdateOne {
	cluo.mutation.SetCardID(id)
	return cluo
}

// SetCard sets the card edge to Word.
func (cluo *CardLogUpdateOne) SetCard(w *Word) *CardLogUpdateOne {
	return cluo.SetCardID(w.ID)
}

// Mutation returns the CardLogMutation object of the builder.
func (cluo *CardLogUpdateOne) Mutation() *CardLogMutation {
	return cluo.mutation
}

// ClearUser clears the user edge to User.
func (cluo *CardLogUpdateOne) ClearUser() *CardLogUpdateOne {
	cluo.mutation.ClearUser()
	return cluo
}

// ClearCard clears the card edge to Word.
func (cluo *CardLogUpdateOne) ClearCard() *CardLogUpdateOne {
	cluo.mutation.ClearCard()
	return cluo
}

// Save executes the query and returns the updated entity.
func (cluo *CardLogUpdateOne) Save(ctx context.Context) (*CardLog, error) {

	if _, ok := cluo.mutation.CardID(); cluo.mutation.CardCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"card\"")
	}
	var (
		err  error
		node *CardLog
	)
	if len(cluo.hooks) == 0 {
		node, err = cluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cluo.mutation = mutation
			node, err = cluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cluo.hooks) - 1; i >= 0; i-- {
			mut = cluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *CardLogUpdateOne) SaveX(ctx context.Context) *CardLog {
	cl, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cl
}

// Exec executes the query on the entity.
func (cluo *CardLogUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *CardLogUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cluo *CardLogUpdateOne) sqlSave(ctx context.Context) (cl *CardLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardlog.Table,
			Columns: cardlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardlog.FieldID,
			},
		},
	}
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CardLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cluo.mutation.Reviewed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cardlog.FieldReviewed,
		})
	}
	if cluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardlog.UserTable,
			Columns: []string{cardlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardlog.UserTable,
			Columns: []string{cardlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cluo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardlog.CardTable,
			Columns: []string{cardlog.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardlog.CardTable,
			Columns: []string{cardlog.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	cl = &CardLog{config: cluo.config}
	_spec.Assign = cl.assignValues
	_spec.ScanValues = cl.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardlog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cl, nil
}
