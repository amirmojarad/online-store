package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Optional(),
		field.String("billing_address").Optional(),
		field.String("country").Optional(),
		field.String("phone"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("customer").Unique().Required(),
		edge.To("cart_products", Product.Type),
		edge.To("orders", Order.Type),
	}
}
