// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/predicate"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CardScheduleUpdate is the builder for updating CardSchedule entities.
type CardScheduleUpdate struct {
	config
	hooks      []Hook
	mutation   *CardScheduleMutation
	predicates []predicate.CardSchedule
}

// Where adds a new predicate for the builder.
func (csu *CardScheduleUpdate) Where(ps ...predicate.CardSchedule) *CardScheduleUpdate {
	csu.predicates = append(csu.predicates, ps...)
	return csu
}

// SetReviewed sets the reviewed field.
func (csu *CardScheduleUpdate) SetReviewed(b bool) *CardScheduleUpdate {
	csu.mutation.SetReviewed(b)
	return csu
}

// SetNillableReviewed sets the reviewed field if the given value is not nil.
func (csu *CardScheduleUpdate) SetNillableReviewed(b *bool) *CardScheduleUpdate {
	if b != nil {
		csu.SetReviewed(*b)
	}
	return csu
}

// SetUserID sets the user edge to User by id.
func (csu *CardScheduleUpdate) SetUserID(id int) *CardScheduleUpdate {
	csu.mutation.SetUserID(id)
	return csu
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (csu *CardScheduleUpdate) SetNillableUserID(id *int) *CardScheduleUpdate {
	if id != nil {
		csu = csu.SetUserID(*id)
	}
	return csu
}

// SetUser sets the user edge to User.
func (csu *CardScheduleUpdate) SetUser(u *User) *CardScheduleUpdate {
	return csu.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (csu *CardScheduleUpdate) SetCardID(id int) *CardScheduleUpdate {
	csu.mutation.SetCardID(id)
	return csu
}

// SetCard sets the card edge to Word.
func (csu *CardScheduleUpdate) SetCard(w *Word) *CardScheduleUpdate {
	return csu.SetCardID(w.ID)
}

// Mutation returns the CardScheduleMutation object of the builder.
func (csu *CardScheduleUpdate) Mutation() *CardScheduleMutation {
	return csu.mutation
}

// ClearUser clears the user edge to User.
func (csu *CardScheduleUpdate) ClearUser() *CardScheduleUpdate {
	csu.mutation.ClearUser()
	return csu
}

// ClearCard clears the card edge to Word.
func (csu *CardScheduleUpdate) ClearCard() *CardScheduleUpdate {
	csu.mutation.ClearCard()
	return csu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (csu *CardScheduleUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := csu.mutation.UpdateTime(); !ok {
		v := cardschedule.UpdateDefaultUpdateTime()
		csu.mutation.SetUpdateTime(v)
	}

	if _, ok := csu.mutation.CardID(); csu.mutation.CardCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"card\"")
	}
	var (
		err      error
		affected int
	)
	if len(csu.hooks) == 0 {
		affected, err = csu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csu.mutation = mutation
			affected, err = csu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csu.hooks) - 1; i >= 0; i-- {
			mut = csu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (csu *CardScheduleUpdate) SaveX(ctx context.Context) int {
	affected, err := csu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (csu *CardScheduleUpdate) Exec(ctx context.Context) error {
	_, err := csu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csu *CardScheduleUpdate) ExecX(ctx context.Context) {
	if err := csu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csu *CardScheduleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardschedule.Table,
			Columns: cardschedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardschedule.FieldID,
			},
		},
	}
	if ps := csu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := csu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardschedule.FieldUpdateTime,
		})
	}
	if value, ok := csu.mutation.Reviewed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cardschedule.FieldReviewed,
		})
	}
	if csu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardschedule.UserTable,
			Columns: []string{cardschedule.UserColumn},
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
	if nodes := csu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardschedule.UserTable,
			Columns: []string{cardschedule.UserColumn},
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
	if csu.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardschedule.CardTable,
			Columns: []string{cardschedule.CardColumn},
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
	if nodes := csu.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardschedule.CardTable,
			Columns: []string{cardschedule.CardColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, csu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardschedule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CardScheduleUpdateOne is the builder for updating a single CardSchedule entity.
type CardScheduleUpdateOne struct {
	config
	hooks    []Hook
	mutation *CardScheduleMutation
}

// SetReviewed sets the reviewed field.
func (csuo *CardScheduleUpdateOne) SetReviewed(b bool) *CardScheduleUpdateOne {
	csuo.mutation.SetReviewed(b)
	return csuo
}

// SetNillableReviewed sets the reviewed field if the given value is not nil.
func (csuo *CardScheduleUpdateOne) SetNillableReviewed(b *bool) *CardScheduleUpdateOne {
	if b != nil {
		csuo.SetReviewed(*b)
	}
	return csuo
}

// SetUserID sets the user edge to User by id.
func (csuo *CardScheduleUpdateOne) SetUserID(id int) *CardScheduleUpdateOne {
	csuo.mutation.SetUserID(id)
	return csuo
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (csuo *CardScheduleUpdateOne) SetNillableUserID(id *int) *CardScheduleUpdateOne {
	if id != nil {
		csuo = csuo.SetUserID(*id)
	}
	return csuo
}

// SetUser sets the user edge to User.
func (csuo *CardScheduleUpdateOne) SetUser(u *User) *CardScheduleUpdateOne {
	return csuo.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (csuo *CardScheduleUpdateOne) SetCardID(id int) *CardScheduleUpdateOne {
	csuo.mutation.SetCardID(id)
	return csuo
}

// SetCard sets the card edge to Word.
func (csuo *CardScheduleUpdateOne) SetCard(w *Word) *CardScheduleUpdateOne {
	return csuo.SetCardID(w.ID)
}

// Mutation returns the CardScheduleMutation object of the builder.
func (csuo *CardScheduleUpdateOne) Mutation() *CardScheduleMutation {
	return csuo.mutation
}

// ClearUser clears the user edge to User.
func (csuo *CardScheduleUpdateOne) ClearUser() *CardScheduleUpdateOne {
	csuo.mutation.ClearUser()
	return csuo
}

// ClearCard clears the card edge to Word.
func (csuo *CardScheduleUpdateOne) ClearCard() *CardScheduleUpdateOne {
	csuo.mutation.ClearCard()
	return csuo
}

// Save executes the query and returns the updated entity.
func (csuo *CardScheduleUpdateOne) Save(ctx context.Context) (*CardSchedule, error) {
	if _, ok := csuo.mutation.UpdateTime(); !ok {
		v := cardschedule.UpdateDefaultUpdateTime()
		csuo.mutation.SetUpdateTime(v)
	}

	if _, ok := csuo.mutation.CardID(); csuo.mutation.CardCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"card\"")
	}
	var (
		err  error
		node *CardSchedule
	)
	if len(csuo.hooks) == 0 {
		node, err = csuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csuo.mutation = mutation
			node, err = csuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csuo.hooks) - 1; i >= 0; i-- {
			mut = csuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (csuo *CardScheduleUpdateOne) SaveX(ctx context.Context) *CardSchedule {
	cs, err := csuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return cs
}

// Exec executes the query on the entity.
func (csuo *CardScheduleUpdateOne) Exec(ctx context.Context) error {
	_, err := csuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csuo *CardScheduleUpdateOne) ExecX(ctx context.Context) {
	if err := csuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (csuo *CardScheduleUpdateOne) sqlSave(ctx context.Context) (cs *CardSchedule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardschedule.Table,
			Columns: cardschedule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardschedule.FieldID,
			},
		},
	}
	id, ok := csuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CardSchedule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := csuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardschedule.FieldUpdateTime,
		})
	}
	if value, ok := csuo.mutation.Reviewed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cardschedule.FieldReviewed,
		})
	}
	if csuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardschedule.UserTable,
			Columns: []string{cardschedule.UserColumn},
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
	if nodes := csuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cardschedule.UserTable,
			Columns: []string{cardschedule.UserColumn},
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
	if csuo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardschedule.CardTable,
			Columns: []string{cardschedule.CardColumn},
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
	if nodes := csuo.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cardschedule.CardTable,
			Columns: []string{cardschedule.CardColumn},
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
	cs = &CardSchedule{config: csuo.config}
	_spec.Assign = cs.assignValues
	_spec.ScanValues = cs.scanValues()
	if err = sqlgraph.UpdateNode(ctx, csuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardschedule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return cs, nil
}