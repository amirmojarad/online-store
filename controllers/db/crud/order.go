package crud

import (
	"log"
	"online-supermarket/controllers/ent"
)

// func (crud Crud) AddOrder(orders []*ent.OrderCreate) ([]*ent.Order, error) {
// 	if addedOrders, err := crud.Client.Order.CreateBulk(orders...).Save(crud.Ctx); err != nil {
// 		log.Println("on AddOrder() in controllers/db/crud/order.go: ", err)
// 		return nil, err
// 	} else {
// 		return addedOrders, nil
// 	}
// }

func (crud Crud) AddOrder(o *ent.Order, productsItems *[]int, customerID int) (*ent.Order, error) {
	if cust, err := crud.GetCustomer(customerID); err != nil {
		log.Println("on AddOrder() in controllers/db/crud/order.go: ", err)
		return nil, err
	} else {
		if newOrder, err := crud.Client.Order.Create().
			SetAmmount(o.Ammount).
			SetDate(o.Date).
			SetEmail(o.Email).SetShippingAddress(o.ShippingAddress).
			SetStatus(o.Status).SetCustomer(cust).AddProductIDs(*productsItems...).Save(crud.Ctx); err != nil {
			log.Println("on AddOrder() in controllers/db/crud/order.go: ", err)
			return nil, err
		} else {
			return newOrder, err
		}

	}
}
