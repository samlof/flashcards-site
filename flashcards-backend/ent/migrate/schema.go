// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CardLogsColumns holds the columns for the "card_logs" table.
	CardLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "result", Type: field.TypeEnum, Enums: []string{"good", "average", "bad"}},
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
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Size: 255},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
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
		UsersTable,
		WordsTable,
	}
)

func init() {
	CardLogsTable.ForeignKeys[0].RefTable = WordsTable
	CardLogsTable.ForeignKeys[1].RefTable = UsersTable
}
