package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(routes *gin.RouterGroup, handler *handler.OrderHandler, jwtMid gin.HandlerFunc) {
	orderRoutes := routes.Group("/").Use(jwtMid)
	orderRoutes.POST("/orders", handler.Order)
}
