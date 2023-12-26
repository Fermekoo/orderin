package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func CallbackRoutes(routes *gin.RouterGroup, handler *handler.PaymentCallbackhandler) {
	routes.POST("/payment-notification", handler.Callback)
}
