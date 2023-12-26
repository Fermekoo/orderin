package routes

import (
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.RouterGroup, handler *handler.UserHandler, middleware gin.HandlerFunc) {

	authRoutes := routes.Group("/auth")
	authRoutes.POST("/register", handler.Register)
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/refresh-token", handler.RefreshToken)

	userRoutes := routes.Group("/user").Use(middleware)
	userRoutes.GET("/profile", handler.Profile)
}
