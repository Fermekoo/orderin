package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

func Start(config *utils.Config) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(config.TimeoutContext)*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "orderin-api api",
			"version": "1.0.0",
		})
	})

	return router
}
