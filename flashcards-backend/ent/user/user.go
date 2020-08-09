// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// EdgeCardLogs holds the string denoting the cardlogs edge name in mutations.
	EdgeCardLogs = "cardLogs"
	// EdgeCardSchedules holds the string denoting the cardschedules edge name in mutations.
	EdgeCardSchedules = "CardSchedules"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CardLogsTable is the table the holds the cardLogs relation/edge.
	CardLogsTable = "card_logs"
	// CardLogsInverseTable is the table name for the CardLog entity.
	// It exists in this package in order to avoid circular dependency with the "cardlog" package.
	CardLogsInverseTable = "card_logs"
	// CardLogsColumn is the table column denoting the cardLogs relation/edge.
	CardLogsColumn = "user_card_logs"
	// CardSchedulesTable is the table the holds the CardSchedules relation/edge.
	CardSchedulesTable = "card_schedules"
	// CardSchedulesInverseTable is the table name for the CardSchedule entity.
	// It exists in this package in order to avoid circular dependency with the "cardschedule" package.
	CardSchedulesInverseTable = "card_schedules"
	// CardSchedulesColumn is the table column denoting the CardSchedules relation/edge.
	CardSchedulesColumn = "user_card_schedules"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
