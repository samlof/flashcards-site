package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// CardLog holds the schema definition for the CardLog entity.
type CardLog struct {
	ent.Schema
}

// Mixin of the CardLog.
func (CardLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Fields of the CardLog.
func (CardLog) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("result").
			Values("easy", "good", "bad", "retry").
			Immutable(),
	}
}

// Edges of the CardLog.
func (CardLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("cardLogs").Unique(),
		edge.To("card", Word.Type).Unique().Required(),
	}
}
