package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/middleware"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(config utils.Config, routes *gin.RouterGroup) {
	service := services.NewOrderService(config, db.Connect(config))
	handler := handler.NewOrderHandler(service)

	orderRoutes := routes.Group("/").Use(middleware.JWTMiddleware(config))
	orderRoutes.POST("/orders", handler.Order)
}
