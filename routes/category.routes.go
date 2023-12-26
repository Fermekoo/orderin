package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(routes *gin.RouterGroup, handler *handler.Categoryhandler) {

	routes.GET("/categories", handler.GetAll)
}
