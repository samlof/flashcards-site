// Package schema contains ent schema
package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/mixin"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Mixin of the Word.
func (Word) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("lang1").NotEmpty().Comment("For example fi").MaxLen(15).Immutable(),
		field.String("lang2").NotEmpty().Comment("For example en").MaxLen(15).Immutable(),
		field.String("word1").MaxLen(255).NotEmpty(),
		field.String("word2").MaxLen(255).NotEmpty(),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cardLogs", CardLog.Type).Ref("card"),
		edge.From("cardSchedules", CardSchedule.Type).Ref("card"),
	}
}
