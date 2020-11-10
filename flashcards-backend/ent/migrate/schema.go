// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// CardLogsColumns holds the columns for the "card_logs" table.
	CardLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "result", Type: field.TypeEnum, Enums: []string{"bad", "easy", "good", "retry"}},
		{Name: "card_log_card", Type: field.TypeInt, Nullable: true},
		{Name: "user_card_logs", Type: field.TypeInt, Nullable: true},
	}
	// CardLogsTable holds the schema information for the "card_logs" table.
	CardLogsTable = &schema.Table{
		Name:       "card_logs",
		Columns:    CardLogsColumns,
		PrimaryKey: []*schema.Column{CardLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "card_logs_words_card",
				Columns: []*schema.Column{CardLogsColumns[3]},

				RefColumns: []*schema.Column{WordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "card_logs_users_cardLogs",
				Columns: []*schema.Column{CardLogsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CardSchedulesColumns holds the columns for the "card_schedules" table.
	CardSchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "scheduled_for", Type: field.TypeTime},
		{Name: "reviewed", Type: field.TypeBool},
		{Name: "card_schedule_card", Type: field.TypeInt, Nullable: true},
		{Name: "user_card_schedules", Type: field.TypeInt, Nullable: true},
	}
	// CardSchedulesTable holds the schema information for the "card_schedules" table.
	CardSchedulesTable = &schema.Table{
		Name:       "card_schedules",
		Columns:    CardSchedulesColumns,
		PrimaryKey: []*schema.Column{CardSchedulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "card_schedules_words_card",
				Columns: []*schema.Column{CardSchedulesColumns[5]},

				RefColumns: []*schema.Column{WordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "card_schedules_users_CardSchedules",
				Columns: []*schema.Column{CardSchedulesColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "firebase_uid", Type: field.TypeString, Unique: true, Size: 255},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_firebase_uid",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// UserSettingsColumns holds the columns for the "user_settings" table.
	UserSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "new_cards_per_day", Type: field.TypeInt, Default: 10},
		{Name: "user_settings", Type: field.TypeInt, Nullable: true},
	}
	// UserSettingsTable holds the schema information for the "user_settings" table.
	UserSettingsTable = &schema.Table{
		Name:       "user_settings",
		Columns:    UserSettingsColumns,
		PrimaryKey: []*schema.Column{UserSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_settings_users_Settings",
				Columns: []*schema.Column{UserSettingsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WordsColumns holds the columns for the "words" table.
	WordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "lang1", Type: field.TypeString, Size: 15},
		{Name: "lang2", Type: field.TypeString, Size: 15},
		{Name: "word1", Type: field.TypeString, Size: 255},
		{Name: "word2", Type: field.TypeString, Size: 255},
	}
	// WordsTable holds the schema information for the "words" table.
	WordsTable = &schema.Table{
		Name:        "words",
		Columns:     WordsColumns,
		PrimaryKey:  []*schema.Column{WordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardLogsTable,
		CardSchedulesTable,
		UsersTable,
		UserSettingsTable,
		WordsTable,
	}
)

func init() {
	CardLogsTable.ForeignKeys[0].RefTable = WordsTable
	CardLogsTable.ForeignKeys[1].RefTable = UsersTable
	CardSchedulesTable.ForeignKeys[0].RefTable = WordsTable
	CardSchedulesTable.ForeignKeys[1].RefTable = UsersTable
	UserSettingsTable.ForeignKeys[0].RefTable = UsersTable
}
