package router

import (
	"purchase-price-calculation/internal/config"
	"purchase-price-calculation/internal/http/middleware"
	"purchase-price-calculation/internal/http/rest"

	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS(cfg))

	// static files
	r.Static("/static", "./static")
	// hello world
	r.GET("/hello", rest.Hello)
	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	return r
}
