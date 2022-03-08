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
	api := API{Router: gin.Default(), Crud: &crud.Crud{Ctx: ctx, Client: client}}

	api.POSTCategory("/category")
	api.GETCategories("/category")
	api.DELETECategory("/category/:id")

	api.Router.Run("localhost:8080")
}
