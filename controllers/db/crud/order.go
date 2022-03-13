package crud

import (
	"log"
	"online-supermarket/controllers/ent"
)

func (crud Crud) AddOrder(orders []*ent.OrderCreate) ([]*ent.Order, error) {
	if addedOrders, err := crud.Client.Order.CreateBulk(orders...).Save(crud.Ctx); err != nil {
		log.Println("on AddOrder() in controllers/db/crud/order.go: ", err)
		return nil, err
	} else {
		return addedOrders, nil
	}
}
