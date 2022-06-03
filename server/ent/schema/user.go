package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("username").Unique(),
		field.String("password").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_users", User.Type),
		edge.From("parent_users", User.Type).Ref("child_users"),
		edge.To("headers", Header.Type),
		// edge.To("headers", Header.Type),
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "user"}}
}
