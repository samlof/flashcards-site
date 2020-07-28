package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CardStatus holds the schema definition for the CardStatus entity.
type CardStatus struct {
	ent.Schema
}

// Fields of the CardStatus.
func (CardStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Time("scheduled_for").Comment("Time when card is scheduled to be done next"),
	}
}

// Edges of the CardStatus.
func (CardStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("cardStatuses").Unique(),
		edge.To("card", Word.Type).Unique(),
	}
}
