package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(config *utils.Config, routes *gin.RouterGroup) {
	productRepo := repositories.NewProductRepo(db.Connect(config))
	service := services.NewProductService(productRepo)
	handler := handler.NewProductHandler(service)
	routes.GET("/products", handler.GetAll)
	routes.GET("/products/:id", handler.Detail)
}
