// Code generated by entc, DO NOT EDIT.

package ent

import (
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CardLog is the model entity for the CardLog schema.
type CardLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Result holds the value of the "result" field.
	Result cardlog.Result `json:"result,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardLogQuery when eager-loading is set.
	Edges          CardLogEdges `json:"edges"`
	card_log_card  *int
	user_card_logs *int
}

// CardLogEdges holds the relations/edges for other nodes in the graph.
type CardLogEdges struct {
	// User holds the value of the user edge.
	User *User
	// Card holds the value of the card edge.
	Card *Word
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardLogEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardLogEdges) CardOrErr() (*Word, error) {
	if e.loadedTypes[1] {
		if e.Card == nil {
			// The edge card was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: word.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CardLog) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullString{}, // result
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CardLog) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // card_log_card
		&sql.NullInt64{}, // user_card_logs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CardLog fields.
func (cl *CardLog) assignValues(values ...interface{}) error {
	if m, n := len(values), len(cardlog.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	cl.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		cl.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field result", values[1])
	} else if value.Valid {
		cl.Result = cardlog.Result(value.String)
	}
	values = values[2:]
	if len(values) == len(cardlog.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field card_log_card", value)
		} else if value.Valid {
			cl.card_log_card = new(int)
			*cl.card_log_card = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_card_logs", value)
		} else if value.Valid {
			cl.user_card_logs = new(int)
			*cl.user_card_logs = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the CardLog.
func (cl *CardLog) QueryUser() *UserQuery {
	return (&CardLogClient{config: cl.config}).QueryUser(cl)
}

// QueryCard queries the card edge of the CardLog.
func (cl *CardLog) QueryCard() *WordQuery {
	return (&CardLogClient{config: cl.config}).QueryCard(cl)
}

// Update returns a builder for updating this CardLog.
// Note that, you need to call CardLog.Unwrap() before calling this method, if this CardLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (cl *CardLog) Update() *CardLogUpdateOne {
	return (&CardLogClient{config: cl.config}).UpdateOne(cl)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (cl *CardLog) Unwrap() *CardLog {
	tx, ok := cl.config.driver.(*txDriver)
	if !ok {
		panic("ent: CardLog is not a transactional entity")
	}
	cl.config.driver = tx.drv
	return cl
}

// String implements the fmt.Stringer.
func (cl *CardLog) String() string {
	var builder strings.Builder
	builder.WriteString("CardLog(")
	builder.WriteString(fmt.Sprintf("id=%v", cl.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(cl.CreateTime.Format(time.ANSIC))
	builder.WriteString(", result=")
	builder.WriteString(fmt.Sprintf("%v", cl.Result))
	builder.WriteByte(')')
	return builder.String()
}

// CardLogs is a parsable slice of CardLog.
type CardLogs []*CardLog

func (cl CardLogs) config(cfg config) {
	for _i := range cl {
		cl[_i].config = cfg
	}
}
