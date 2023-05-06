package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fermekoo/orderin-api/routes"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	router *gin.Engine
	config utils.Config
}

func NewApiServer(config utils.Config) (*ApiServer, error) {
	server := &ApiServer{
		config: config,
	}
	server.setupRouter()
	return server, nil
}

func (server *ApiServer) setupRouter() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "orderin-api api",
			"version": "1.0.0",
		})
	})

	v1 := router.Group("/v1")
	routes.UserRoutes(server.config, v1)
	routes.CategoryRoutes(server.config, v1)
	routes.ProductRoutes(server.config, v1)

	server.router = router
}

func (server *ApiServer) Start(address string, ctx context.Context) {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	server_err := make(chan error, 1)
	go func() {
		server_err <- srv.ListenAndServe()

	}()

	shutdown_channel := make(chan os.Signal, 1)
	signal.Notify(shutdown_channel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown_channel:
		log.Println("shutdown signal", sig)
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
		}
	case err := <-server_err:
		if err != nil {
			log.Fatalf("server : %v", err)
		}
	}
	log.Printf("http server run on port %s", address)
}
