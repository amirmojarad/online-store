package crud

import (
	"log"
	"online-supermarket/controllers/ent"
	"online-supermarket/controllers/ent/product"
)

func (crud Crud) GetProducts() ([]*ent.Product, error) {
	if products, err := crud.Client.Product.Query().Order(ent.Asc(product.FieldID)).All(crud.Ctx); err != nil {
		log.Println("on GetProducts in controllers/db/crud/product.go: ", err)
		return nil, err
	} else {
		return products, nil
	}

}
