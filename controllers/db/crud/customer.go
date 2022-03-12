package crud

import (
	"log"
	"online-supermarket/controllers/ent"
	"online-supermarket/controllers/ent/customer"
	"online-supermarket/utils"
)

func (crud *Crud) AddUser(customer *ent.Customer) (*ent.Customer, error) {
	hashedPassword, _ := utils.HashPassword(customer.Password)
	if createdCustomer, err := crud.Client.Customer.Create().
		SetBillingAddress(customer.BillingAddress).
		SetCountry(customer.Country).
		SetEmail(customer.Email).
		SetFullName(customer.FullName).
		SetPassword(hashedPassword).
		SetPhone(customer.Phone).
		Save(crud.Ctx); err != nil {
		log.Println("on AddUser Function in controllers/crud/Customer.go: ", err)
		return nil, err
	} else {
		return createdCustomer, nil
	}
}

func (crud *Crud) GetUserByEmail(email string) (*ent.Customer, error) {
	if fetchedUser, err := crud.Client.Customer.Query().
		Select(customer.FieldEmail, customer.FieldPassword, customer.FieldID).
		First(crud.Ctx); err != nil {
		log.Println("on GetUserByEmail Function in controllers/db/crud/customer.go: ", err)
		return nil, err
	} else {
		return fetchedUser, nil
	}
}
