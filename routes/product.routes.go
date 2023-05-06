package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(config utils.Config, routes *gin.RouterGroup) {
	service := services.NewProductService(db.Connect(config))
	handler := handler.NewProductHandler(service)
	routes.GET("/products", handler.GetAll)
}
