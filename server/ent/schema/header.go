package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Header struct {
	ent.Schema
}

func (Header) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("status").Optional(),
		field.Int32("receiverId").Optional(),
		field.Int32("senderId").Optional(),
		field.Time("createdAt"), field.Int32("messageId").Optional(),
	}
}
func (Header) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("headers").Unique().Field("senderId").Field("receiverId"),
		edge.To("message", Message.Type).Unique(),
	}
}
func (Header) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "header"}}
}
