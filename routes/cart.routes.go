package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/middleware"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func CartRoutes(config utils.Config, routes *gin.RouterGroup) {

	service := services.NewCartService(config, db.Connect(config))
	handler := handler.NewCartHandler(service)

	cartRoutes := routes.Group("/").Use(middleware.JWTMiddleware(config))
	cartRoutes.POST("/carts", handler.AddCart)
	cartRoutes.GET("/carts", handler.GetAll)
}
