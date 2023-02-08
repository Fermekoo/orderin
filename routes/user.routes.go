package routes

import (
	"log"

	"github.com/Fermekoo/go-kapster/db"
	"github.com/Fermekoo/go-kapster/handler"
	"github.com/Fermekoo/go-kapster/middleware"
	"github.com/Fermekoo/go-kapster/services"
	"github.com/Fermekoo/go-kapster/utils"
	"github.com/Fermekoo/go-kapster/utils/token"
	"github.com/gin-gonic/gin"
)

func UserRoutes(config utils.Config, routes *gin.RouterGroup) {
	tokenMaker, err := token.NewJWTMaker()
	if err != nil {
		log.Fatal("failed to setup token maker %w", err)
	}
	service := services.NewUserService(config, db.Connect(), tokenMaker)
	handler := handler.NewUserHandler(service)
	authRoutes := routes.Group("/auth")
	authRoutes.POST("/register", handler.Register)
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/refresh-token", handler.RefreshToken)

	userRoutes := routes.Group("/user").Use(middleware.JWTMiddleware(config, tokenMaker))
	userRoutes.GET("/profile", handler.Profile)
}
