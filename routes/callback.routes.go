package routes

import (
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func CallbackRoutes(config *utils.Config, routes *gin.RouterGroup) {
	service := services.NewOrderService(config, db.Connect(config))
	handler := handler.NewPaymentCallbackHandler(service)

	callbackRoutes := routes.Group("/")
	callbackRoutes.POST("/payment-notification", handler.Callback)
}
