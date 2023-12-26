package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func CartRoutes(routes *gin.RouterGroup, handler *handler.CartHandler, jwtMid gin.HandlerFunc) {

	cartRoutes := routes.Group("/").Use(jwtMid)
	cartRoutes.POST("/carts", handler.AddCart)
	cartRoutes.GET("/carts", handler.GetAll)
	cartRoutes.PUT("carts/:id", handler.UpdateQty)
	cartRoutes.DELETE("/carts/:id", handler.Delete)
}
