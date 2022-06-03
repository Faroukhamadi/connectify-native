// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Faroukhamadi/connectify-native/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDNotIn(ids ...int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDGT(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HeaderID applies equality check predicate on the "header_id" field. It's identical to HeaderIDEQ.
func HeaderID(v int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaderID), v))
	})
}

// IsFromSender applies equality check predicate on the "is_from_sender" field. It's identical to IsFromSenderEQ.
func IsFromSender(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFromSender), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Read applies equality check predicate on the "read" field. It's identical to ReadEQ.
func Read(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRead), v))
	})
}

// SentAt applies equality check predicate on the "sent_at" field. It's identical to SentAtEQ.
func SentAt(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentAt), v))
	})
}

// HeaderIDEQ applies the EQ predicate on the "header_id" field.
func HeaderIDEQ(v int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaderID), v))
	})
}

// HeaderIDNEQ applies the NEQ predicate on the "header_id" field.
func HeaderIDNEQ(v int32) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeaderID), v))
	})
}

// HeaderIDIn applies the In predicate on the "header_id" field.
func HeaderIDIn(vs ...int32) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHeaderID), v...))
	})
}

// HeaderIDNotIn applies the NotIn predicate on the "header_id" field.
func HeaderIDNotIn(vs ...int32) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHeaderID), v...))
	})
}

// HeaderIDIsNil applies the IsNil predicate on the "header_id" field.
func HeaderIDIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHeaderID)))
	})
}

// HeaderIDNotNil applies the NotNil predicate on the "header_id" field.
func HeaderIDNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHeaderID)))
	})
}

// IsFromSenderEQ applies the EQ predicate on the "is_from_sender" field.
func IsFromSenderEQ(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFromSender), v))
	})
}

// IsFromSenderNEQ applies the NEQ predicate on the "is_from_sender" field.
func IsFromSenderNEQ(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsFromSender), v))
	})
}

// IsFromSenderIsNil applies the IsNil predicate on the "is_from_sender" field.
func IsFromSenderIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsFromSender)))
	})
}

// IsFromSenderNotNil applies the NotNil predicate on the "is_from_sender" field.
func IsFromSenderNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsFromSender)))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ReadEQ applies the EQ predicate on the "read" field.
func ReadEQ(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRead), v))
	})
}

// ReadNEQ applies the NEQ predicate on the "read" field.
func ReadNEQ(v bool) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRead), v))
	})
}

// ReadIsNil applies the IsNil predicate on the "read" field.
func ReadIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRead)))
	})
}

// ReadNotNil applies the NotNil predicate on the "read" field.
func ReadNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRead)))
	})
}

// SentAtEQ applies the EQ predicate on the "sent_at" field.
func SentAtEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentAt), v))
	})
}

// SentAtNEQ applies the NEQ predicate on the "sent_at" field.
func SentAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSentAt), v))
	})
}

// SentAtIn applies the In predicate on the "sent_at" field.
func SentAtIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSentAt), v...))
	})
}

// SentAtNotIn applies the NotIn predicate on the "sent_at" field.
func SentAtNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSentAt), v...))
	})
}

// SentAtGT applies the GT predicate on the "sent_at" field.
func SentAtGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSentAt), v))
	})
}

// SentAtGTE applies the GTE predicate on the "sent_at" field.
func SentAtGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSentAt), v))
	})
}

// SentAtLT applies the LT predicate on the "sent_at" field.
func SentAtLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSentAt), v))
	})
}

// SentAtLTE applies the LTE predicate on the "sent_at" field.
func SentAtLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSentAt), v))
	})
}

// SentAtIsNil applies the IsNil predicate on the "sent_at" field.
func SentAtIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSentAt)))
	})
}

// SentAtNotNil applies the NotNil predicate on the "sent_at" field.
func SentAtNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSentAt)))
	})
}

// HasHeader applies the HasEdge predicate on the "header" edge.
func HasHeader() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHeaderWith applies the HasEdge predicate on the "header" edge with a given conditions (other predicates).
func HasHeaderWith(preds ...predicate.Header) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}