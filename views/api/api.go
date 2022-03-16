package api

import (
	"context"
	"online-supermarket/controllers/db/crud"
	"online-supermarket/controllers/ent"
	"strconv"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	Crud   *crud.Crud
}

func (api API) GetIdParameter(ctx *gin.Context, param string) (int, error) {
	return strconv.Atoi(ctx.Param(param))
}

func RunAPI(ctx context.Context, client *ent.Client) {
	api := API{Router: gin.New(), Crud: &crud.Crud{Ctx: ctx, Client: client}}

	api.Router.Use(gin.Logger())
	api.Router.Use(gin.Recovery())

	api.AuthRouter()
	api.CustomerRouter()

	api.Router.Run("localhost:8080")
}
