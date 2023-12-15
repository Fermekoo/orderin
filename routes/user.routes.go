package routes

import (
	"log"

	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/middleware"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
)

func UserRoutes(config *utils.Config, routes *gin.RouterGroup) {
	tokenMaker, err := token.NewJWTMaker()
	if err != nil {
		log.Fatal("failed to setup token maker %w", err)
	}
	db := db.Connect(config)
	userRepo := repositories.NewUserRepo(db)
	sessionRepo := repositories.NewSessionRepo(db)
	service := services.NewUserService(config, tokenMaker, userRepo, sessionRepo)
	handler := handler.NewUserHandler(service)
	authRoutes := routes.Group("/auth")
	authRoutes.POST("/register", handler.Register)
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/refresh-token", handler.RefreshToken)

	userRoutes := routes.Group("/user").Use(middleware.JWTMiddleware(config))
	userRoutes.GET("/profile", handler.Profile)
}
