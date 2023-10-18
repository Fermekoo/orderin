package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(config *utils.Config, routes *gin.RouterGroup) {
	service := services.NewCategoryService(db.Connect(config))
	handler := handler.NewCategoryHandler(service)
	routes.GET("/categories", handler.GetAll)
}
