package api

import (
	"context"
	"online-supermarket/controllers/db/crud"
	"online-supermarket/controllers/ent"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	Crud   *crud.Crud
}

func RunAPI(ctx context.Context, client *ent.Client) {
	api := API{Router: gin.New(), Crud: &crud.Crud{Ctx: ctx, Client: client}}

	api.Router.Use(gin.Logger())
	api.Router.Use(gin.Recovery())

	api.LoginUser("/users/login")
	api.SignUpUser("/users/signup")

	// api.POSTCustomer("/users")
	// api.GETCustomer("/users")
	// api.PATCHCustomer("/users")
	// api.GETPurchasedProducts("/users/producst")
	// api.POSTPurchasedProducts("/users/producst")

	// api.GETCustomerOrders("/users/orders")

	// api.POSTProduct("/products")
	// api.GETProducts("/products")

	api.POSTCategory("/category")
	api.GETCategories("/category")
	api.DELETECategory("/category/:id")

	api.Router.Run("localhost:8080")
}
