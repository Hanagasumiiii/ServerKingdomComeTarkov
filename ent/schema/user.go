package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			NotEmpty().
			Unique().
			Comment("Уникальное имя пользователя"),
		field.String("password").
			NotEmpty().
			Sensitive().
			Comment("Хэшированный пароль пользователя"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("player", Player.Type).
			Unique().
			Comment("Один к одному к Player"),
	}
}
