package api

import "github.com/gin-gonic/gin"

func (api API) CustomerRouter() {
	category := api.Router.Group("/category")
	category.POST("/users", postCustomer())
	category.POST("/users/:id/purchase", postPurchaseToCustomer())
	category.POST("/users/:id/orders", postOrderToCustomer())
	category.POST("/users/:id/cart", postCartItemToCustomer())
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
