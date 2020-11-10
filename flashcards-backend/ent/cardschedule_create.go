// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// CardScheduleCreate is the builder for creating a CardSchedule entity.
type CardScheduleCreate struct {
	config
	mutation *CardScheduleMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (csc *CardScheduleCreate) SetCreateTime(t time.Time) *CardScheduleCreate {
	csc.mutation.SetCreateTime(t)
	return csc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (csc *CardScheduleCreate) SetNillableCreateTime(t *time.Time) *CardScheduleCreate {
	if t != nil {
		csc.SetCreateTime(*t)
	}
	return csc
}

// SetUpdateTime sets the update_time field.
func (csc *CardScheduleCreate) SetUpdateTime(t time.Time) *CardScheduleCreate {
	csc.mutation.SetUpdateTime(t)
	return csc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (csc *CardScheduleCreate) SetNillableUpdateTime(t *time.Time) *CardScheduleCreate {
	if t != nil {
		csc.SetUpdateTime(*t)
	}
	return csc
}

// SetScheduledFor sets the scheduled_for field.
func (csc *CardScheduleCreate) SetScheduledFor(t time.Time) *CardScheduleCreate {
	csc.mutation.SetScheduledFor(t)
	return csc
}

// SetReviewed sets the reviewed field.
func (csc *CardScheduleCreate) SetReviewed(b bool) *CardScheduleCreate {
	csc.mutation.SetReviewed(b)
	return csc
}

// SetNillableReviewed sets the reviewed field if the given value is not nil.
func (csc *CardScheduleCreate) SetNillableReviewed(b *bool) *CardScheduleCreate {
	if b != nil {
		csc.SetReviewed(*b)
	}
	return csc
}

// SetUserID sets the user edge to User by id.
func (csc *CardScheduleCreate) SetUserID(id int) *CardScheduleCreate {
	csc.mutation.SetUserID(id)
	return csc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (csc *CardScheduleCreate) SetNillableUserID(id *int) *CardScheduleCreate {
	if id != nil {
		csc = csc.SetUserID(*id)
	}
	return csc
}

// SetUser sets the user edge to User.
func (csc *CardScheduleCreate) SetUser(u *User) *CardScheduleCreate {
	return csc.SetUserID(u.ID)
}

// SetCardID sets the card edge to Word by id.
func (csc *CardScheduleCreate) SetCardID(id int) *CardScheduleCreate {
	csc.mutation.SetCardID(id)
	return csc
}

// SetCard sets the card edge to Word.
func (csc *CardScheduleCreate) SetCard(w *Word) *CardScheduleCreate {
	return csc.SetCardID(w.ID)
}

// Mutation returns the CardScheduleMutation object of the builder.
func (csc *CardScheduleCreate) Mutation() *CardScheduleMutation {
	return csc.mutation
}

// Save creates the CardSchedule in the database.
func (csc *CardScheduleCreate) Save(ctx context.Context) (*CardSchedule, error) {
	var (
		err  error
		node *CardSchedule
	)
	csc.defaults()
	if len(csc.hooks) == 0 {
		if err = csc.check(); err != nil {
			return nil, err
		}
		node, err = csc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = csc.check(); err != nil {
				return nil, err
			}
			csc.mutation = mutation
			node, err = csc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(csc.hooks) - 1; i >= 0; i-- {
			mut = csc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (csc *CardScheduleCreate) SaveX(ctx context.Context) *CardSchedule {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (csc *CardScheduleCreate) defaults() {
	if _, ok := csc.mutation.CreateTime(); !ok {
		v := cardschedule.DefaultCreateTime()
		csc.mutation.SetCreateTime(v)
	}
	if _, ok := csc.mutation.UpdateTime(); !ok {
		v := cardschedule.DefaultUpdateTime()
		csc.mutation.SetUpdateTime(v)
	}
	if _, ok := csc.mutation.Reviewed(); !ok {
		v := cardschedule.DefaultReviewed
		csc.mutation.SetReviewed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csc *CardScheduleCreate) check() error {
	if _, ok := csc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := csc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := csc.mutation.ScheduledFor(); !ok {
		return &ValidationError{Name: "scheduled_for", err: errors.New("ent: missing required field \"scheduled_for\"")}
	}
	if _, ok := csc.mutation.Reviewed(); !ok {
		return &ValidationError{Name: "reviewed", err: errors.New("ent: missing required field \"reviewed\"")}
	}
	if _, ok := csc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card", err: errors.New("ent: missing required edge \"card\"")}
	}
	return nil
}

func (csc *CardScheduleCreate) sqlSave(ctx context.Context) (*CardSchedule, error) {
	_node, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (csc *CardScheduleCreate) createSpec() (*CardSchedule, *sqlgraph.CreateSpec) {
	var (
		_node = &CardSchedule{config: csc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cardschedule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardschedule.FieldID,
			},
		}
	)
	if value, ok := csc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardschedule.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := csc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardschedule.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := csc.mutation.ScheduledFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cardschedule.FieldScheduledFor,
		})
		_node.ScheduledFor = value
	}
	if value, ok := csc.mutation.Reviewed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cardschedule.FieldReviewed,
		})
		_node.Reviewed = value
	}
	if nodes := csc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csc.mutation.CardIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardScheduleCreateBulk is the builder for creating a bulk of CardSchedule entities.
type CardScheduleCreateBulk struct {
	config
	builders []*CardScheduleCreate
}

// Save creates the CardSchedule entities in the database.
func (cscb *CardScheduleCreateBulk) Save(ctx context.Context) ([]*CardSchedule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cscb.builders))
	nodes := make([]*CardSchedule, len(cscb.builders))
	mutators := make([]Mutator, len(cscb.builders))
	for i := range cscb.builders {
		func(i int, root context.Context) {
			builder := cscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardScheduleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (cscb *CardScheduleCreateBulk) SaveX(ctx context.Context) []*CardSchedule {
	v, err := cscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
