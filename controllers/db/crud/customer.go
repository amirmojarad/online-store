package crud

import (
	"log"
	"online-supermarket/controllers/ent"
	"online-supermarket/controllers/ent/order"
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

func (crud Crud) GetOrdersProducts(customerID, orderID int) ([]*ent.Product, error) {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		return nil, err
	} else {
		if o, err := cust.QueryOrders().Where(order.ID(orderID)).First(crud.Ctx); err != nil {
			return nil, err
		} else {
			if products, err := o.QueryProducts().All(crud.Ctx); err != nil {
				return nil, err
			} else {
				return products, nil
			}
		}
	}
}

func (crud Crud) GetOrdersOfCustomer(id int) ([]*ent.Order, error) {
	if cust, err := crud.GetCustomer(id); err != nil {
		return nil, err
	} else {
		return cust.QueryOrders().All(crud.Ctx)
	}
}

func (crud Crud) GetOrderOfCustomer(customerID, orderID int) (*ent.Order, error) {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		return nil, err
	} else {
		return cust.QueryOrders().Where(order.ID(orderID)).First(crud.Ctx)
	}
}

func (crud Crud) AddProductsToCart(ids *[]int, customerID int) ([]*ent.Product, error) {
	cust, err := crud.GetCustomer(customerID)
	if err != nil {
		log.Println("on AddProductsToCart() in controllers/db/crud/order.go: ", err)
		return nil, err
	}
	cust, err = cust.Update().AddCartProductIDs(*ids...).Save(crud.Ctx)
	if err != nil {
		log.Println("on AddProductsToCart() in controllers/db/crud/order.go: ", err)
		return nil, err
	}
	return cust.QueryCartProducts().AllX(crud.Ctx), nil
}

func (crud Crud) GetAllProducts(customerID int) ([]*ent.Product, error) {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		log.Println("on GetAllProducts() in controllers/db/crud/order.go: ", err)
		return nil, err
	} else {
		return cust.QueryCartProducts().All(crud.Ctx)
	}
}

func (crud Crud) DeleteItems(cartItems *[]int, customerID int) error {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		log.Println("on DeleteItems() in controllers/db/crud/order.go: ", err)
		return err
	} else {
		cust.Update().RemoveCartProductIDs(*cartItems...).Exec(crud.Ctx)
		return nil
	}
}

func (crud Crud) DeleteOrder(orderIDs *[]int, customerID int) error {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		return nil
	} else {
		cust.Update().RemoveCartProductIDs(*orderIDs...).Save(crud.Ctx)
		return nil
	}
}

func (crud Crud) ChangeOrderStatus(customerID, orderID int, status order.Status) error {
	if o, err := crud.GetOrderOfCustomer(customerID, orderID); err != nil {
		return err
	} else {
		return o.Update().SetStatus(status).Exec(crud.Ctx)
	}
}
