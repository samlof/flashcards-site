// Code generated by entc, DO NOT EDIT.

package word

import (
	"time"
)

const (
	// Label holds the string label denoting the word type in the database.
	Label = "word"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldLang1 holds the string denoting the lang1 field in the database.
	FieldLang1 = "lang1"
	// FieldLang2 holds the string denoting the lang2 field in the database.
	FieldLang2 = "lang2"
	// FieldWord1 holds the string denoting the word1 field in the database.
	FieldWord1 = "word1"
	// FieldWord2 holds the string denoting the word2 field in the database.
	FieldWord2 = "word2"

	// EdgeCardLogs holds the string denoting the cardlogs edge name in mutations.
	EdgeCardLogs = "cardLogs"
	// EdgeCardSchedules holds the string denoting the cardschedules edge name in mutations.
	EdgeCardSchedules = "cardSchedules"

	// Table holds the table name of the word in the database.
	Table = "words"
	// CardLogsTable is the table the holds the cardLogs relation/edge.
	CardLogsTable = "card_logs"
	// CardLogsInverseTable is the table name for the CardLog entity.
	// It exists in this package in order to avoid circular dependency with the "cardlog" package.
	CardLogsInverseTable = "card_logs"
	// CardLogsColumn is the table column denoting the cardLogs relation/edge.
	CardLogsColumn = "card_log_card"
	// CardSchedulesTable is the table the holds the cardSchedules relation/edge.
	CardSchedulesTable = "card_schedules"
	// CardSchedulesInverseTable is the table name for the CardSchedule entity.
	// It exists in this package in order to avoid circular dependency with the "cardschedule" package.
	CardSchedulesInverseTable = "card_schedules"
	// CardSchedulesColumn is the table column denoting the cardSchedules relation/edge.
	CardSchedulesColumn = "card_schedule_card"
)

// Columns holds all SQL columns for word fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldLang1,
	FieldLang2,
	FieldWord1,
	FieldWord2,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// Lang1Validator is a validator for the "lang1" field. It is called by the builders before save.
	Lang1Validator func(string) error
	// Lang2Validator is a validator for the "lang2" field. It is called by the builders before save.
	Lang2Validator func(string) error
	// Word1Validator is a validator for the "word1" field. It is called by the builders before save.
	Word1Validator func(string) error
	// Word2Validator is a validator for the "word2" field. It is called by the builders before save.
	Word2Validator func(string) error
)
