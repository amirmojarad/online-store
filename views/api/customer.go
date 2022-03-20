package api

import (
	"fmt"
	"log"
	"net/http"
	"online-supermarket/controllers/ent"
	"online-supermarket/controllers/ent/order"
	"online-supermarket/models"
	"online-supermarket/views/middlewares/security"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api API) CustomerRouter() {
	category := api.Router.Group("/users", security.CheckAuth())
	// Customer Endpoints
	category.POST("/", api.postCustomer())
	category.GET("/:id", api.getCustomer())

	// Cart Endpoints
	category.POST("/:id/cart", api.postCartItemToCustomer())
	category.GET("/:id/cart", api.getCartItems())
	category.DELETE("/:id/cart", api.deleteCartItems())
	// Orders Endpoint
	customerOrdersPath := "/:id/orders"
	category.POST(customerOrdersPath, api.postOrderToCustomer())
	category.GET(customerOrdersPath, api.getCustomerOrders())
	category.DELETE(customerOrdersPath, api.deleteCustomerOrders())
	category.PATCH(customerOrdersPath+"/:orderID", api.patchCustomerOrder())

	// Get All Products
	category.GET("/:id/products/all", api.getAllProducts())
}

type OrderAndProducts struct {
	Order    *ent.Order
	Products []*ent.Product
}

func (api API) getCustomerOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := api.GetIdParameter(ctx, "id")
		if orders, err := api.Crud.GetOrdersOfCustomer(id); err != nil {
			log.Println("on getCustomerOrders() in view/api/customer.go: ", err)
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			result := make([]OrderAndProducts, len(orders))
			for i, item := range orders {
				result[i] = OrderAndProducts{Order: item, Products: item.QueryProducts().AllX(ctx)}
			}
			ctx.IndentedJSON(http.StatusOK, result)
		}
	}
}

func (api API) deleteCartItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
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

type OrderModel struct {
	Order        *ent.Order
	ProductItems *[]int
}

func (api API) postOrderToCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := api.GetIdParameter(ctx, "id")
		om := OrderModel{}
		log.Println("OM", om)
		ctx.BindJSON(&om)
		log.Println("OM", om)
		if o, err := api.Crud.AddOrder(om.Order, om.ProductItems, id); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"order": o,
			})
		}
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

func (api API) getCartItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if items, err := api.Crud.GetAllProducts(id); err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("found no cart item for customer with id: %d", id),
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, items)
			return
		}

	}
}

func (api API) patchCustomerOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerID, _ := api.GetIdParameter(ctx, "id")
		orderID, _ := api.GetIdParameter(ctx, "orderID")
		body := make(map[string]string)
		ctx.BindJSON(&body)
		if err := api.Crud.ChangeOrderStatus(customerID, orderID, order.Status(body["status"])); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, "changed")
		}
	}
}

func (api API) deleteCustomerOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := api.GetIdParameter(ctx, "id")
		productsIDs := &[]int{}
		ctx.BindJSON(&productsIDs)
		if err := api.Crud.DeleteItems(productsIDs, id); err != nil {
			log.Println("on deleteCartItems() in view/api/customer.go: ", err)
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("items with ids %+v deleted", productsIDs))
	}
}
