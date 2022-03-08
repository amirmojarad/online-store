// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "full_name", Type: field.TypeString, Nullable: true},
		{Name: "billing_address", Type: field.TypeString, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ammount", Type: field.TypeFloat64},
		{Name: "shipping_address", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"in_progress", "delivered", "referred", "canceled"}},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sku", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "weight", Type: field.TypeFloat64},
		{Name: "descriptions", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "stock", Type: field.TypeInt, Default: 0},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		CustomersTable,
		OrdersTable,
		ProductsTable,
	}
)

func init() {
}
