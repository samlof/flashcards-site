// Code generated by entc, DO NOT EDIT.

package cardschedule

import (
	"flashcards-backend/ent/predicate"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ScheduledFor applies equality check predicate on the "scheduled_for" field. It's identical to ScheduledForEQ.
func ScheduledFor(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScheduledFor), v))
	})
}

// Reviewed applies equality check predicate on the "reviewed" field. It's identical to ReviewedEQ.
func Reviewed(v bool) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewed), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ScheduledForEQ applies the EQ predicate on the "scheduled_for" field.
func ScheduledForEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScheduledFor), v))
	})
}

// ScheduledForNEQ applies the NEQ predicate on the "scheduled_for" field.
func ScheduledForNEQ(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScheduledFor), v))
	})
}

// ScheduledForIn applies the In predicate on the "scheduled_for" field.
func ScheduledForIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScheduledFor), v...))
	})
}

// ScheduledForNotIn applies the NotIn predicate on the "scheduled_for" field.
func ScheduledForNotIn(vs ...time.Time) predicate.CardSchedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSchedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScheduledFor), v...))
	})
}

// ScheduledForGT applies the GT predicate on the "scheduled_for" field.
func ScheduledForGT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScheduledFor), v))
	})
}

// ScheduledForGTE applies the GTE predicate on the "scheduled_for" field.
func ScheduledForGTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScheduledFor), v))
	})
}

// ScheduledForLT applies the LT predicate on the "scheduled_for" field.
func ScheduledForLT(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScheduledFor), v))
	})
}

// ScheduledForLTE applies the LTE predicate on the "scheduled_for" field.
func ScheduledForLTE(v time.Time) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScheduledFor), v))
	})
}

// ReviewedEQ applies the EQ predicate on the "reviewed" field.
func ReviewedEQ(v bool) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewed), v))
	})
}

// ReviewedNEQ applies the NEQ predicate on the "reviewed" field.
func ReviewedNEQ(v bool) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReviewed), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Word) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.CardSchedule) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.CardSchedule) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CardSchedule) predicate.CardSchedule {
	return predicate.CardSchedule(func(s *sql.Selector) {
		p(s.Not())
	})
}