package api

import (
	"log"
	"net/http"
	"online-supermarket/controllers/ent"
	"online-supermarket/views/middlewares/security"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) POSTCategory(path string) {
	api.Router.POST(path, security.CheckAuth(), func(ctx *gin.Context) {
		categoryModel := ent.Category{}
		if err := ctx.BindJSON(&categoryModel); err != nil {
			log.Println("error occured while binding json to model: ", err)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error_message":     err,
				"valid_json_schema": categoryModel,
			})
			return
		}
		log.Println(categoryModel)
		if addedCategory, err := api.Crud.AddCategory(&categoryModel); err != nil {
			log.Println("error occured when adding category to database: ", err)
			ctx.IndentedJSON(http.StatusServiceUnavailable, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, addedCategory)
			return
		}
	})
}

func (api *API) GETCategories(path string) {
	api.Router.GET(path, security.CheckAuth(), func(ctx *gin.Context) {
		if categories, err := api.Crud.GetAllCategories(); err != nil {
			log.Println("error occured when fetching categories to database: ", err)
			ctx.IndentedJSON(http.StatusServiceUnavailable, err.Error())
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, categories)
		}
	})
}

func (api *API) DELETECategory(path string) {
	api.Router.DELETE(path, func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		if result, err := api.Crud.DeleteCategory(id); err != nil {
			log.Println("error occured when fetching categories to database: ", err)
			ctx.IndentedJSON(http.StatusConflict, gin.H{
				"message": err.Error(),
				"result":  result,
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, result)
			return
		}
	})
}
