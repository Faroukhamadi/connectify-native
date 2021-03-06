// Code generated by entc, DO NOT EDIT.

package header

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Faroukhamadi/connectify-native/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
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
func IDNotIn(ids ...int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
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
func IDGT(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// ReceiverId applies equality check predicate on the "receiverId" field. It's identical to ReceiverIdEQ.
func ReceiverId(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiverId), v))
	})
}

// SenderId applies equality check predicate on the "senderId" field. It's identical to SenderIdEQ.
func SenderId(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSenderId), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// MessageId applies equality check predicate on the "messageId" field. It's identical to MessageIdEQ.
func MessageId(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageId), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// ReceiverIdEQ applies the EQ predicate on the "receiverId" field.
func ReceiverIdEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReceiverId), v))
	})
}

// ReceiverIdNEQ applies the NEQ predicate on the "receiverId" field.
func ReceiverIdNEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReceiverId), v))
	})
}

// ReceiverIdIn applies the In predicate on the "receiverId" field.
func ReceiverIdIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReceiverId), v...))
	})
}

// ReceiverIdNotIn applies the NotIn predicate on the "receiverId" field.
func ReceiverIdNotIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReceiverId), v...))
	})
}

// ReceiverIdIsNil applies the IsNil predicate on the "receiverId" field.
func ReceiverIdIsNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReceiverId)))
	})
}

// ReceiverIdNotNil applies the NotNil predicate on the "receiverId" field.
func ReceiverIdNotNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReceiverId)))
	})
}

// SenderIdEQ applies the EQ predicate on the "senderId" field.
func SenderIdEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSenderId), v))
	})
}

// SenderIdNEQ applies the NEQ predicate on the "senderId" field.
func SenderIdNEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSenderId), v))
	})
}

// SenderIdIn applies the In predicate on the "senderId" field.
func SenderIdIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSenderId), v...))
	})
}

// SenderIdNotIn applies the NotIn predicate on the "senderId" field.
func SenderIdNotIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSenderId), v...))
	})
}

// SenderIdGT applies the GT predicate on the "senderId" field.
func SenderIdGT(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSenderId), v))
	})
}

// SenderIdGTE applies the GTE predicate on the "senderId" field.
func SenderIdGTE(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSenderId), v))
	})
}

// SenderIdLT applies the LT predicate on the "senderId" field.
func SenderIdLT(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSenderId), v))
	})
}

// SenderIdLTE applies the LTE predicate on the "senderId" field.
func SenderIdLTE(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSenderId), v))
	})
}

// SenderIdIsNil applies the IsNil predicate on the "senderId" field.
func SenderIdIsNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSenderId)))
	})
}

// SenderIdNotNil applies the NotNil predicate on the "senderId" field.
func SenderIdNotNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSenderId)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// MessageIdEQ applies the EQ predicate on the "messageId" field.
func MessageIdEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageId), v))
	})
}

// MessageIdNEQ applies the NEQ predicate on the "messageId" field.
func MessageIdNEQ(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessageId), v))
	})
}

// MessageIdIn applies the In predicate on the "messageId" field.
func MessageIdIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessageId), v...))
	})
}

// MessageIdNotIn applies the NotIn predicate on the "messageId" field.
func MessageIdNotIn(vs ...int32) predicate.Header {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Header(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessageId), v...))
	})
}

// MessageIdGT applies the GT predicate on the "messageId" field.
func MessageIdGT(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessageId), v))
	})
}

// MessageIdGTE applies the GTE predicate on the "messageId" field.
func MessageIdGTE(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessageId), v))
	})
}

// MessageIdLT applies the LT predicate on the "messageId" field.
func MessageIdLT(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessageId), v))
	})
}

// MessageIdLTE applies the LTE predicate on the "messageId" field.
func MessageIdLTE(v int32) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessageId), v))
	})
}

// MessageIdIsNil applies the IsNil predicate on the "messageId" field.
func MessageIdIsNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessageId)))
	})
}

// MessageIdNotNil applies the NotNil predicate on the "messageId" field.
func MessageIdNotNil() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessageId)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
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

// HasMessage applies the HasEdge predicate on the "message" edge.
func HasMessage() predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessageTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MessageTable, MessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessageWith applies the HasEdge predicate on the "message" edge with a given conditions (other predicates).
func HasMessageWith(preds ...predicate.Message) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MessageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, MessageTable, MessageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Header) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Header) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
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
func Not(p predicate.Header) predicate.Header {
	return predicate.Header(func(s *sql.Selector) {
		p(s.Not())
	})
}
