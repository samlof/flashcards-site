// Code generated by entc, DO NOT EDIT.

package cardschedule

import (
	"time"
)

const (
	// Label holds the string label denoting the cardschedule type in the database.
	Label = "card_schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldScheduledFor holds the string denoting the scheduled_for field in the database.
	FieldScheduledFor = "scheduled_for"
	// FieldReviewed holds the string denoting the reviewed field in the database.
	FieldReviewed = "reviewed"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"

	// Table holds the table name of the cardschedule in the database.
	Table = "card_schedules"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "card_schedules"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_card_schedules"
	// CardTable is the table the holds the card relation/edge.
	CardTable = "card_schedules"
	// CardInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	CardInverseTable = "words"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "card_schedule_card"
)

// Columns holds all SQL columns for cardschedule fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldScheduledFor,
	FieldReviewed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CardSchedule type.
var ForeignKeys = []string{
	"card_schedule_card",
	"user_card_schedules",
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultReviewed holds the default value on creation for the reviewed field.
	DefaultReviewed bool
)
