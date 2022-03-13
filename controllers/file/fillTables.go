package file

import (
	"context"
	"log"
	"online-supermarket/controllers/ent"
)

func fetchProductsFromJsonFile() *[]ent.Product {
	file := CreateFile("controllers/file/resources/products.json")
	pl := []ent.Product{}
	if err := file.ReadJson(&pl); err != nil {
		log.Println("on FillProducts an error occured: ", err)
		return nil
	}
	return &pl
}

func FillProducts(client *ent.Client, ctx *context.Context) {
	products := fetchProductsFromJsonFile()
	bulk := make([]*ent.ProductCreate, len(*products))
	for i, item := range *products {
		bulk[i] = client.Product.Create().
			SetName(item.Name).
			SetDescriptions(item.Descriptions).
			SetPrice(item.Price).
			SetSku(item.Sku).
			SetStock(item.Stock).
			SetThumbnail(item.Thumbnail).
			SetWeight(item.Weight)
	}
	client.Product.CreateBulk(bulk...).Save(*ctx)
}
