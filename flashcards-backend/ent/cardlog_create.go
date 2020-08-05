// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CardLogCreate is the builder for creating a CardLog entity.
type CardLogCreate struct {
	config
	mutation *CardLogMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (clc *CardLogCreate) SetCreateTime(t time.Time) *CardLogCreate {
	clc.mutation.SetCreateTime(t)
	return clc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (clc *CardLogCreate) SetNillableCreateTime(t *time.Time) *CardLogCreate {
	if t != nil {
		clc.SetCreateTime(*t)
	}
	return clc
}

// SetResult sets the result field.
func (clc *CardLogCreate) SetResult(c cardlog.Result) *CardLogCreate {
	clc.mutation.SetResult(c)
	return clc
}

// SetScheduledFor sets the scheduled_for field.
func (clc *CardLogCreate) SetScheduledFor(t time.Time) *CardLogCreate {
	clc.mutation.SetScheduledFor(t)
	return clc
}

// SetUserID sets the user edge to User by id.
func (clc *CardLogCreate) SetUserID(id int) *CardLogCreate {
	clc.mutation.SetUserID(id)
	return clc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (clc *CardLogCreate) SetNillableUserID(id *int) *CardLogCreate {
	if id != nil {
		clc = clc.SetUserID(*id)
	}
	return clc
}

// SetUser sets the user edge to User.
func (clc *CardLogCreate) SetUser(u *User) *CardLogCreate {
	return clc.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (clc *CardLogCreate) SetCardID(id int) *CardLogCreate {
	clc.mutation.SetCardID(id)
	return clc
}

// SetCard sets the card edge to Word.
func (clc *CardLogCreate) SetCard(w *Word) *CardLogCreate {
	return clc.SetCardID(w.ID)
}

// Mutation returns the CardLogMutation object of the builder.
func (clc *CardLogCreate) Mutation() *CardLogMutation {
	return clc.mutation
}

// Save creates the CardLog in the database.
func (clc *CardLogCreate) Save(ctx context.Context) (*CardLog, error) {
	if err := clc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *CardLog
	)
	if len(clc.hooks) == 0 {
		node, err = clc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			clc.mutation = mutation
			node, err = clc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(clc.hooks) - 1; i >= 0; i-- {
			mut = clc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (clc *CardLogCreate) SaveX(ctx context.Context) *CardLog {
	v, err := clc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (clc *CardLogCreate) preSave() error {
	if _, ok := clc.mutation.CreateTime(); !ok {
		v := cardlog.DefaultCreateTime()
		clc.mutation.SetCreateTime(v)
	}
	if _, ok := clc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New("ent: missing required field \"result\"")}
	}
	if v, ok := clc.mutation.Result(); ok {
		if err := cardlog.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	if _, ok := clc.mutation.ScheduledFor(); !ok {
		return &ValidationError{Name: "scheduled_for", err: errors.New("ent: missing required field \"scheduled_for\"")}
	}
	if _, ok := clc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card", err: errors.New("ent: missing required edge \"card\"")}
	}
	return nil
}

func (clc *CardLogCreate) sqlSave(ctx context.Context) (*CardLog, error) {
	cl, _spec := clc.createSpec()
	if err := sqlgraph.CreateNode(ctx, clc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	cl.ID = int(id)
	return cl, nil
}

func (clc *CardLogCreate) createSpec() (*CardLog, *sqlgraph.CreateSpec) {
	var (
		cl    = &CardLog{config: clc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cardlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardlog.FieldID,
			},
		}
	)
	if value, ok := clc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardlog.FieldCreateTime,
		})
		cl.CreateTime = value
	}
	if value, ok := clc.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: cardlog.FieldResult,
		})
		cl.Result = value
	}
	if value, ok := clc.mutation.ScheduledFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardlog.FieldScheduledFor,
		})
		cl.ScheduledFor = value
	}
	if nodes := clc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clc.mutation.CardIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return cl, _spec
}

// CardLogCreateBulk is the builder for creating a bulk of CardLog entities.
type CardLogCreateBulk struct {
	config
	builders []*CardLogCreate
}

// Save creates the CardLog entities in the database.
func (clcb *CardLogCreateBulk) Save(ctx context.Context) ([]*CardLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(clcb.builders))
	nodes := make([]*CardLog, len(clcb.builders))
	mutators := make([]Mutator, len(clcb.builders))
	for i := range clcb.builders {
		func(i int, root context.Context) {
			builder := clcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*CardLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, clcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, clcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, clcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (clcb *CardLogCreateBulk) SaveX(ctx context.Context) []*CardLog {
	v, err := clcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
