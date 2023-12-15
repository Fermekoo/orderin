package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/middleware"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(config *utils.Config, routes *gin.RouterGroup) {
	db := db.Connect(config)
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)
	service := services.NewOrderService(config, orderRepo, cartRepo)
	handler := handler.NewOrderHandler(service)

	orderRoutes := routes.Group("/").Use(middleware.JWTMiddleware(config))
	orderRoutes.POST("/orders", handler.Order)
}
