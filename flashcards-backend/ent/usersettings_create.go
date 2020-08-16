// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/usersettings"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserSettingsCreate is the builder for creating a UserSettings entity.
type UserSettingsCreate struct {
	config
	mutation *UserSettingsMutation
	hooks    []Hook
}

// SetCreateTime sets the create_time field.
func (usc *UserSettingsCreate) SetCreateTime(t time.Time) *UserSettingsCreate {
	usc.mutation.SetCreateTime(t)
	return usc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (usc *UserSettingsCreate) SetNillableCreateTime(t *time.Time) *UserSettingsCreate {
	if t != nil {
		usc.SetCreateTime(*t)
	}
	return usc
}

// SetUpdateTime sets the update_time field.
func (usc *UserSettingsCreate) SetUpdateTime(t time.Time) *UserSettingsCreate {
	usc.mutation.SetUpdateTime(t)
	return usc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (usc *UserSettingsCreate) SetNillableUpdateTime(t *time.Time) *UserSettingsCreate {
	if t != nil {
		usc.SetUpdateTime(*t)
	}
	return usc
}

// SetNewCardsPerDay sets the newCardsPerDay field.
func (usc *UserSettingsCreate) SetNewCardsPerDay(i int) *UserSettingsCreate {
	usc.mutation.SetNewCardsPerDay(i)
	return usc
}

// SetNillableNewCardsPerDay sets the newCardsPerDay field if the given value is not nil.
func (usc *UserSettingsCreate) SetNillableNewCardsPerDay(i *int) *UserSettingsCreate {
	if i != nil {
		usc.SetNewCardsPerDay(*i)
	}
	return usc
}

// SetUserID sets the user edge to User by id.
func (usc *UserSettingsCreate) SetUserID(id int) *UserSettingsCreate {
	usc.mutation.SetUserID(id)
	return usc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (usc *UserSettingsCreate) SetNillableUserID(id *int) *UserSettingsCreate {
	if id != nil {
		usc = usc.SetUserID(*id)
	}
	return usc
}

// SetUser sets the user edge to User.
func (usc *UserSettingsCreate) SetUser(u *User) *UserSettingsCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSettingsMutation object of the builder.
func (usc *UserSettingsCreate) Mutation() *UserSettingsMutation {
	return usc.mutation
}

// Save creates the UserSettings in the database.
func (usc *UserSettingsCreate) Save(ctx context.Context) (*UserSettings, error) {
	if err := usc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *UserSettings
	)
	if len(usc.hooks) == 0 {
		node, err = usc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSettingsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			usc.mutation = mutation
			node, err = usc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usc.hooks) - 1; i >= 0; i-- {
			mut = usc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSettingsCreate) SaveX(ctx context.Context) *UserSettings {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (usc *UserSettingsCreate) preSave() error {
	if _, ok := usc.mutation.CreateTime(); !ok {
		v := usersettings.DefaultCreateTime()
		usc.mutation.SetCreateTime(v)
	}
	if _, ok := usc.mutation.UpdateTime(); !ok {
		v := usersettings.DefaultUpdateTime()
		usc.mutation.SetUpdateTime(v)
	}
	if _, ok := usc.mutation.NewCardsPerDay(); !ok {
		v := usersettings.DefaultNewCardsPerDay
		usc.mutation.SetNewCardsPerDay(v)
	}
	return nil
}

func (usc *UserSettingsCreate) sqlSave(ctx context.Context) (*UserSettings, error) {
	us, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	us.ID = int(id)
	return us, nil
}

func (usc *UserSettingsCreate) createSpec() (*UserSettings, *sqlgraph.CreateSpec) {
	var (
		us    = &UserSettings{config: usc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usersettings.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersettings.FieldID,
			},
		}
	)
	if value, ok := usc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usersettings.FieldCreateTime,
		})
		us.CreateTime = value
	}
	if value, ok := usc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usersettings.FieldUpdateTime,
		})
		us.UpdateTime = value
	}
	if value, ok := usc.mutation.NewCardsPerDay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: usersettings.FieldNewCardsPerDay,
		})
		us.NewCardsPerDay = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersettings.UserTable,
			Columns: []string{usersettings.UserColumn},
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
	return us, _spec
}

// UserSettingsCreateBulk is the builder for creating a bulk of UserSettings entities.
type UserSettingsCreateBulk struct {
	config
	builders []*UserSettingsCreate
}

// Save creates the UserSettings entities in the database.
func (uscb *UserSettingsCreateBulk) Save(ctx context.Context) ([]*UserSettings, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSettings, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*UserSettingsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (uscb *UserSettingsCreateBulk) SaveX(ctx context.Context) []*UserSettings {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}