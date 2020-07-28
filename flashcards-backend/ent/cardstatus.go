// Code generated by entc, DO NOT EDIT.

package ent

import (
	"flashcards-backend/ent/cardstatus"
	"flashcards-backend/ent/user"
	"flashcards-backend/ent/word"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CardStatus is the model entity for the CardStatus schema.
type CardStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DoneLast holds the value of the "done_last" field.
	DoneLast time.Time `json:"done_last,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardStatusQuery when eager-loading is set.
	Edges              CardStatusEdges `json:"edges"`
	card_status_card   *int
	user_card_statuses *int
}

// CardStatusEdges holds the relations/edges for other nodes in the graph.
type CardStatusEdges struct {
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
func (e CardStatusEdges) UserOrErr() (*User, error) {
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
func (e CardStatusEdges) CardOrErr() (*Word, error) {
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
func (*CardStatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // done_last
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CardStatus) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // card_status_card
		&sql.NullInt64{}, // user_card_statuses
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CardStatus fields.
func (cs *CardStatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(cardstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	cs.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field done_last", values[0])
	} else if value.Valid {
		cs.DoneLast = value.Time
	}
	values = values[1:]
	if len(values) == len(cardstatus.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field card_status_card", value)
		} else if value.Valid {
			cs.card_status_card = new(int)
			*cs.card_status_card = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_card_statuses", value)
		} else if value.Valid {
			cs.user_card_statuses = new(int)
			*cs.user_card_statuses = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the CardStatus.
func (cs *CardStatus) QueryUser() *UserQuery {
	return (&CardStatusClient{config: cs.config}).QueryUser(cs)
}

// QueryCard queries the card edge of the CardStatus.
func (cs *CardStatus) QueryCard() *WordQuery {
	return (&CardStatusClient{config: cs.config}).QueryCard(cs)
}

// Update returns a builder for updating this CardStatus.
// Note that, you need to call CardStatus.Unwrap() before calling this method, if this CardStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CardStatus) Update() *CardStatusUpdateOne {
	return (&CardStatusClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (cs *CardStatus) Unwrap() *CardStatus {
	tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CardStatus is not a transactional entity")
	}
	cs.config.driver = tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CardStatus) String() string {
	var builder strings.Builder
	builder.WriteString("CardStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", cs.ID))
	builder.WriteString(", done_last=")
	builder.WriteString(cs.DoneLast.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CardStatusSlice is a parsable slice of CardStatus.
type CardStatusSlice []*CardStatus

func (cs CardStatusSlice) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
