package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Float("ammount").Min(0.0),
		field.String("shipping_address"),
		field.String("email"),
		field.Time("date").Default(time.Now()),
		field.Enum("status").Values("in_progress", "delivered", "referred", "canceled"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
