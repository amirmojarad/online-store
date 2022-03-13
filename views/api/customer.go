package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) CustomerRouter() {
	category := api.Router.Group("/users")
	category.POST("", postCustomer())
	category.POST("/:id/purchase", postPurchaseToCustomer())
	category.POST("/:id/orders", postOrderToCustomer())
	category.POST("/:id/cart", postCartItemToCustomer())
	category.GET("/:id/products/all", api.getAllProducts())
}

func (api API) getAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if products, err := api.Crud.GetProducts(); err != nil {
			log.Println("on getAllProducts in view/api/customer.go: ", err)
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, products)
		}
	}
}

func postCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func postOrderToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func postPurchaseToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func postCartItemToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
