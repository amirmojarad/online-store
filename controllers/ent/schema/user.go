package schema

import (
	"fmt"
	"net/mail"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func ValidateEmail() func(email string) error {
	return func(email string) error {
		if _, err := mail.ParseAddress(email); err != nil {
			return fmt.Errorf("%s is not a valid email", email)
		}
		return nil
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique().Validate(ValidateEmail()),
		field.String("password").NotEmpty().MinLen(6),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Customer.Type).Unique(),
	}
}
