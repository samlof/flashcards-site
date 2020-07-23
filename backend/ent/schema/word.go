// Package schema contains ent
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("langData").NotEmpty().Comment("For example fi-en").MaxLen(15),
		field.String("word1").MaxLen(255).NotEmpty(),
		field.String("word2").MaxLen(255).NotEmpty(),
		field.Time("created_at").
			Default(func() time.Time { return time.Now().UTC() }).Immutable(),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return nil
}
