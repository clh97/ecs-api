package internal

import (
	"github.com/clh97/ecs/internal/endpoints"
	"github.com/clh97/ecs/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRouter is where the API routes will be defined
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.RequestID)

	r.GET("/", endpoints.Home)

	return r
}
