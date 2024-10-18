package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int("mana").
			Default(100).
			Comment("Мана игрока"),
		field.Int("hp").
			Default(100).
			Comment("HP игрока"),
		field.Float("position_x").
			Default(0.0).
			Comment("Позиция X"),
		field.Float("position_y").
			Default(0.0).
			Comment("Позиция Y"),
		field.Float("position_z").
			Default(0.0).
			Comment("Позиция Z"),
		field.JSON("inventory", []string{}).
			Optional().
			Comment("Инвентарь игрока"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("player").
			Unique().
			Required().
			Comment("Связь с User"),
	}
}
