// Code generated by entc, DO NOT EDIT.

package cardlog

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the cardlog type in the database.
	Label = "card_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"

	// Table holds the table name of the cardlog in the database.
	Table = "card_logs"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "card_logs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_card_logs"
	// CardTable is the table the holds the card relation/edge.
	CardTable = "card_logs"
	// CardInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	CardInverseTable = "words"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "card_log_card"
)

// Columns holds all SQL columns for cardlog fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldResult,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CardLog type.
var ForeignKeys = []string{
	"card_log_card",
	"user_card_logs",
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
)

// Result defines the type for the result enum field.
type Result string

// Result values.
const (
	ResultGood    Result = "good"
	ResultAverage Result = "average"
	ResultBad     Result = "bad"
)

func (r Result) String() string {
	return string(r)
}

// ResultValidator is a validator for the "r" field enum values. It is called by the builders before save.
func ResultValidator(r Result) error {
	switch r {
	case ResultGood, ResultAverage, ResultBad:
		return nil
	default:
		return fmt.Errorf("cardlog: invalid enum value for result field: %q", r)
	}
}

// ResultAll returns all values of Result enum
func ResultAll() []Result {
	return []Result{
		ResultGood,
		ResultAverage,
		ResultBad,
	}
}
