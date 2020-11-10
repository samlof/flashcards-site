// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/word"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// WordCreate is the builder for creating a Word entity.
type WordCreate struct {
	config
	mutation *WordMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (wc *WordCreate) SetCreateTime(t time.Time) *WordCreate {
	wc.mutation.SetCreateTime(t)
	return wc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (wc *WordCreate) SetNillableCreateTime(t *time.Time) *WordCreate {
	if t != nil {
		wc.SetCreateTime(*t)
	}
	return wc
}

// SetUpdateTime sets the update_time field.
func (wc *WordCreate) SetUpdateTime(t time.Time) *WordCreate {
	wc.mutation.SetUpdateTime(t)
	return wc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (wc *WordCreate) SetNillableUpdateTime(t *time.Time) *WordCreate {
	if t != nil {
		wc.SetUpdateTime(*t)
	}
	return wc
}

// SetLang1 sets the lang1 field.
func (wc *WordCreate) SetLang1(s string) *WordCreate {
	wc.mutation.SetLang1(s)
	return wc
}

// SetLang2 sets the lang2 field.
func (wc *WordCreate) SetLang2(s string) *WordCreate {
	wc.mutation.SetLang2(s)
	return wc
}

// SetWord1 sets the word1 field.
func (wc *WordCreate) SetWord1(s string) *WordCreate {
	wc.mutation.SetWord1(s)
	return wc
}

// SetWord2 sets the word2 field.
func (wc *WordCreate) SetWord2(s string) *WordCreate {
	wc.mutation.SetWord2(s)
	return wc
}

// AddCardLogIDs adds the cardLogs edge to CardLog by ids.
func (wc *WordCreate) AddCardLogIDs(ids ...int) *WordCreate {
	wc.mutation.AddCardLogIDs(ids...)
	return wc
}

// AddCardLogs adds the cardLogs edges to CardLog.
func (wc *WordCreate) AddCardLogs(c ...*CardLog) *WordCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wc.AddCardLogIDs(ids...)
}

// AddCardScheduleIDs adds the cardSchedules edge to CardSchedule by ids.
func (wc *WordCreate) AddCardScheduleIDs(ids ...int) *WordCreate {
	wc.mutation.AddCardScheduleIDs(ids...)
	return wc
}

// AddCardSchedules adds the cardSchedules edges to CardSchedule.
func (wc *WordCreate) AddCardSchedules(c ...*CardSchedule) *WordCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wc.AddCardScheduleIDs(ids...)
}

// Mutation returns the WordMutation object of the builder.
func (wc *WordCreate) Mutation() *WordMutation {
	return wc.mutation
}

// Save creates the Word in the database.
func (wc *WordCreate) Save(ctx context.Context) (*Word, error) {
	var (
		err  error
		node *Word
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			node, err = wc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WordCreate) SaveX(ctx context.Context) *Word {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wc *WordCreate) defaults() {
	if _, ok := wc.mutation.CreateTime(); !ok {
		v := word.DefaultCreateTime()
		wc.mutation.SetCreateTime(v)
	}
	if _, ok := wc.mutation.UpdateTime(); !ok {
		v := word.DefaultUpdateTime()
		wc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WordCreate) check() error {
	if _, ok := wc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := wc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := wc.mutation.Lang1(); !ok {
		return &ValidationError{Name: "lang1", err: errors.New("ent: missing required field \"lang1\"")}
	}
	if v, ok := wc.mutation.Lang1(); ok {
		if err := word.Lang1Validator(v); err != nil {
			return &ValidationError{Name: "lang1", err: fmt.Errorf("ent: validator failed for field \"lang1\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Lang2(); !ok {
		return &ValidationError{Name: "lang2", err: errors.New("ent: missing required field \"lang2\"")}
	}
	if v, ok := wc.mutation.Lang2(); ok {
		if err := word.Lang2Validator(v); err != nil {
			return &ValidationError{Name: "lang2", err: fmt.Errorf("ent: validator failed for field \"lang2\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Word1(); !ok {
		return &ValidationError{Name: "word1", err: errors.New("ent: missing required field \"word1\"")}
	}
	if v, ok := wc.mutation.Word1(); ok {
		if err := word.Word1Validator(v); err != nil {
			return &ValidationError{Name: "word1", err: fmt.Errorf("ent: validator failed for field \"word1\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Word2(); !ok {
		return &ValidationError{Name: "word2", err: errors.New("ent: missing required field \"word2\"")}
	}
	if v, ok := wc.mutation.Word2(); ok {
		if err := word.Word2Validator(v); err != nil {
			return &ValidationError{Name: "word2", err: fmt.Errorf("ent: validator failed for field \"word2\": %w", err)}
		}
	}
	return nil
}

func (wc *WordCreate) sqlSave(ctx context.Context) (*Word, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *WordCreate) createSpec() (*Word, *sqlgraph.CreateSpec) {
	var (
		_node = &Word{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: word.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: word.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := wc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: word.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := wc.mutation.Lang1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldLang1,
		})
		_node.Lang1 = value
	}
	if value, ok := wc.mutation.Lang2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldLang2,
		})
		_node.Lang2 = value
	}
	if value, ok := wc.mutation.Word1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldWord1,
		})
		_node.Word1 = value
	}
	if value, ok := wc.mutation.Word2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldWord2,
		})
		_node.Word2 = value
	}
	if nodes := wc.mutation.CardLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   word.CardLogsTable,
			Columns: []string{word.CardLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cardlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CardSchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   word.CardSchedulesTable,
			Columns: []string{word.CardSchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cardschedule.FieldID,
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

// WordCreateBulk is the builder for creating a bulk of Word entities.
type WordCreateBulk struct {
	config
	builders []*WordCreate
}

// Save creates the Word entities in the database.
func (wcb *WordCreateBulk) Save(ctx context.Context) ([]*Word, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Word, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WordMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wcb *WordCreateBulk) SaveX(ctx context.Context) []*Word {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
