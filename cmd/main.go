package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fermekoo/orderin-api/api"
	"github.com/Fermekoo/orderin-api/db"
	"github.com/Fermekoo/orderin-api/handler"
	"github.com/Fermekoo/orderin-api/middleware"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/routes"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
)

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("failed to setup config %w", err)
	}

	ginEngine := api.Start(&config)

	v1 := ginEngine.Group("/v1")
	v1.Static("/assets", "./assets")
	db := db.Connect(&config)
	tokenMaker, err := token.NewJWTMaker()
	if err != nil {
		log.Fatal("failed to setup token maker %w", err)
	}
	userRepo := repositories.NewUserRepo(db)
	sessionRepo := repositories.NewSessionRepo(db)
	userService := services.NewUserService(&config, tokenMaker, userRepo, sessionRepo)
	userHandler := handler.NewUserHandler(userService)
	jwtMid := middleware.JWTMiddleware(&config)
	routes.UserRoutes(v1, userHandler, jwtMid)

	categoriesRepo := repositories.NewCategoriesRepo(db)
	categoryService := services.NewCategoryService(categoriesRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	routes.CategoryRoutes(v1, categoryHandler)

	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	routes.ProductRoutes(v1, productHandler)

	cartRepo := repositories.NewCartRepo(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)
	routes.CartRoutes(v1, cartHandler, jwtMid)

	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(&config, orderRepo, cartRepo)
	orderHandler := handler.NewOrderHandler(orderService)
	routes.OrderRoutes(v1, orderHandler, jwtMid)

	callbakHandler := handler.NewPaymentCallbackHandler(orderService)
	routes.CallbackRoutes(v1, callbakHandler)

	srv := &http.Server{
		Addr:    config.HTTPServerPort,
		Handler: ginEngine,
	}

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- srv.ListenAndServe()
	}()
	ctx := context.Background()
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdownChannel:
		log.Println("shutdown signal", sig)
		ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(config.TimeoutContext)*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctxWithTimeout); err != nil {
			srv.Close()
		}
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server : %v", err)
		}
	}
	log.Printf("run on port %s", srv.Addr)
}
