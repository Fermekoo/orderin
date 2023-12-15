package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func CallbackRoutes(config *utils.Config, routes *gin.RouterGroup) {
	db := db.Connect(config)
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)
	service := services.NewOrderService(config, orderRepo, cartRepo)
	handler := handler.NewPaymentCallbackHandler(service)

	callbackRoutes := routes.Group("/")
	callbackRoutes.POST("/payment-notification", handler.Callback)
}
