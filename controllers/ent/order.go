// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"online-supermarket/controllers/ent/customer"
	"online-supermarket/controllers/ent/order"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Ammount holds the value of the "ammount" field.
	Ammount float64 `json:"ammount,omitempty"`
	// ShippingAddress holds the value of the "shipping_address" field.
	ShippingAddress string `json:"shipping_address,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Status holds the value of the "status" field.
	Status order.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges           OrderEdges `json:"edges"`
	customer_orders *int
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[1] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldAmmount:
			values[i] = new(sql.NullFloat64)
		case order.FieldID:
			values[i] = new(sql.NullInt64)
		case order.FieldShippingAddress, order.FieldEmail, order.FieldStatus:
			values[i] = new(sql.NullString)
		case order.FieldDate:
			values[i] = new(sql.NullTime)
		case order.ForeignKeys[0]: // customer_orders
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldAmmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ammount", values[i])
			} else if value.Valid {
				o.Ammount = value.Float64
			}
		case order.FieldShippingAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shipping_address", values[i])
			} else if value.Valid {
				o.ShippingAddress = value.String
			}
		case order.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				o.Email = value.String
			}
		case order.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				o.Date = value.Time
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = order.Status(value.String)
			}
		case order.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_orders", value)
			} else if value.Valid {
				o.customer_orders = new(int)
				*o.customer_orders = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProducts queries the "products" edge of the Order entity.
func (o *Order) QueryProducts() *ProductQuery {
	return (&OrderClient{config: o.config}).QueryProducts(o)
}

// QueryCustomer queries the "customer" edge of the Order entity.
func (o *Order) QueryCustomer() *CustomerQuery {
	return (&OrderClient{config: o.config}).QueryCustomer(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", ammount=")
	builder.WriteString(fmt.Sprintf("%v", o.Ammount))
	builder.WriteString(", shipping_address=")
	builder.WriteString(o.ShippingAddress)
	builder.WriteString(", email=")
	builder.WriteString(o.Email)
	builder.WriteString(", date=")
	builder.WriteString(o.Date.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
