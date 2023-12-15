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

func CartRoutes(config *utils.Config, routes *gin.RouterGroup) {
	cartRepo := repositories.NewCartRepo(db.Connect(config))
	service := services.NewCartService(cartRepo)
	handler := handler.NewCartHandler(service)

	cartRoutes := routes.Group("/").Use(middleware.JWTMiddleware(config))
	cartRoutes.POST("/carts", handler.AddCart)
	cartRoutes.GET("/carts", handler.GetAll)
	cartRoutes.PUT("carts/:id", handler.UpdateQty)
	cartRoutes.DELETE("/carts/:id", handler.Delete)
}
