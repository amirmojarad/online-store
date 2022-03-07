package schema

import (
	"fmt"
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func ValidateEmail() func(email string) error {
	return func(email string) error {
		if _, err := mail.ParseAddress(email); err != nil {
			return fmt.Errorf("%s is not a valid email", email)
		}
		return nil
	}
}

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Validate(ValidateEmail()),
		field.String("password"),
		field.String("full_name").Optional(),
		field.String("billing_address").Optional(),
		field.String("country").Optional(),
		field.String("phone"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
