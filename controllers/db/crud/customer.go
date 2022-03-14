package crud

import (
	"log"
	"online-supermarket/controllers/ent"
	"online-supermarket/models"
)

func (crud Crud) AddCustomer(user *models.User) (*ent.Customer, error) {
	u, err := crud.GetUserByEmail(&user.Email)
	if err != nil {
		log.Println("on AddCustomer in controllers/db/crud/customer.go: ", err)
		return nil, err
	}
	newCustomer, err := crud.Client.Customer.Create().
		SetPhone(user.Phone).
		SetBillingAddress(user.BillingAddress).
		SetCountry(user.Country).
		SetUser(u).
		Save(crud.Ctx)
	if err != nil {
		log.Println("on AddCustomer in controllers/db/crud/customer.go: ", err)
		return nil, err
	}
	return newCustomer, nil
}

func (crud Crud) GetCustomer(id int) (*ent.Customer, error) {
	if fetchedCustomer, err := crud.Client.Customer.Get(crud.Ctx, id); err != nil {
		log.Println("on GetCustomer in controllers/db/crud/customer.go: ", err)
		return nil, err
	} else {
		return fetchedCustomer, nil
	}
}

func (crud Crud) GetOrdersOfUser(id int) ([]*ent.Order, error) {
	if orders, err := crud.Client.Customer.Query().QueryOrders().All(crud.Ctx); err != nil {
		log.Println("on AddOrder() in controllers/db/crud/order.go: ", err)
		return nil, err
	} else {
		return orders, nil
	}
}

func (crud Crud) GetPurchasedProducts(id int) ([]*ent.Product, error) {
	if fetchedCustomer, err := crud.GetCustomer(id); err != nil {
		log.Println("on GetPurchasedProducts() in controllers/db/crud/order.go: ", err)
		return nil, err
	} else {
		if products, err := fetchedCustomer.QueryPurchasedProducts().All(crud.Ctx); err != nil {
			log.Println("on GetPurchasedProducts() in controllers/db/crud/order.go: ", err)
			return nil, err
		} else {
			return products, err
		}
	}
}
