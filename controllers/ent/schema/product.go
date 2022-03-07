package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		// SKU stands for “stock keeping unit”
		// it is a number (usually eight alphanumeric digits)
		field.String("sku"),
		field.String("name"),
		field.Float("price"),
		field.Float("weight"),
		field.String("descriptions"),
		field.String("thumbnail"),
		field.Time("create_date").Default(time.Now()),
		field.Int("stock").Default(0),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}