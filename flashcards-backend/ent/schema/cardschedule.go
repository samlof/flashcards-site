package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/mixin"
)

// CardSchedule holds the schema definition for the CardSchedule entity.
type CardSchedule struct {
	ent.Schema
}

// Mixin of the CardSchedule.
func (CardSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the CardSchedule.
func (CardSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.Time("scheduled_for").
			Immutable(),
		field.Bool("reviewed").
			Default(false),
	}
}

// Edges of the CardSchedule.
func (CardSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("CardSchedules").Unique(),
		edge.To("card", Word.Type).Unique().Required(),
	}
}
