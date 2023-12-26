package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(routes *gin.RouterGroup, handler *handler.ProductHandler) {
	routes.GET("/products", handler.GetAll)
	routes.GET("/products/:id", handler.Detail)
}
