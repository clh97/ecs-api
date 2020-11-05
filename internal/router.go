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
	r.Use(middlewares.RequestID())

	/* Public routes */
	public := r.Group("/api/v1")
	public.POST("/auth", endpoints.Auth)

	/* Private routes */
	private := r.Group("/api/v1")
	private.Use(middlewares.JWT())
	{
		private.GET("/", endpoints.Home)
	}

	return r
}
