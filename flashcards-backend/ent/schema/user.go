package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
	"github.com/facebookincubator/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Optional().Comment("User email").MaxLen(255),
		field.String("firebaseUid").MinLen(5).Unique().Immutable().MaxLen(255),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cardLogs", CardLog.Type),
		edge.To("CardSchedules", CardSchedule.Type),
		edge.To("Settings", UserSettings.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("firebaseUid").Unique(),
	}
}
