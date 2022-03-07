// Code generated by entc, DO NOT EDIT.

package order

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmmount holds the string denoting the ammount field in the database.
	FieldAmmount = "ammount"
	// FieldShippingAddress holds the string denoting the shipping_address field in the database.
	FieldShippingAddress = "shipping_address"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the order in the database.
	Table = "orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldAmmount,
	FieldShippingAddress,
	FieldEmail,
	FieldDate,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AmmountValidator is a validator for the "ammount" field. It is called by the builders before save.
	AmmountValidator func(float64) error
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusInProgress Status = "in_progress"
	StatusDelivered  Status = "delivered"
	StatusReferred   Status = "referred"
	StatusCanceled   Status = "canceled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInProgress, StatusDelivered, StatusReferred, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}