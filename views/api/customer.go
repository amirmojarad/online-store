package api

import (
	"log"
	"net/http"
	"online-supermarket/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api API) CustomerRouter() {
	category := api.Router.Group("/users")
	category.POST("/", api.postCustomer())
	category.POST("/:id/orders", postOrderToCustomer())
	category.POST("/:id/cart", api.postCartItemToCustomer())
	category.GET("/:id/products/all", api.getAllProducts())
	category.GET("/:id", api.getCustomer())
}

func (api API) getCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if user, err := api.Crud.GetUserByID(id); err != nil {
			log.Println("on getCustomer in view/api/customer.go: ", err)
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
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

func (api API) postCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := &models.User{}
		ctx.BindJSON(&u)
		if newCustomer, err := api.Crud.AddCustomer(u); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			relatedUser, _ := newCustomer.QueryUser().First(api.Crud.Ctx)
			ctx.IndentedJSON(http.StatusCreated, gin.H{
				"customer":     newCustomer,
				"related_user": relatedUser,
			})
			return
		}
	}
}

func postOrderToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (api API) postCartItemToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerID, _ := strconv.Atoi(ctx.Param("id"))
		productsIDs := &[]int{}
		if err := ctx.BindJSON(&productsIDs); err != nil {
			ctx.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
				"message": "UnprocessableEntity",
			})
			return
		}
		if products, err := api.Crud.AddProductsToCart(productsIDs, customerID); err != nil {
			log.Println("on postCartItemToCustomer() in view/api/customer.go: ", err)
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			ctx.IndentedJSON(http.StatusCreated, products)
		}
	}
}
