package pkg

import (
	"github.com/clh97/ecs/pkg/endpoints"
	"github.com/clh97/ecs/pkg/middlewares"
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
	public.POST("/sign-up", endpoints.Create)
	public.POST("/sign-in", endpoints.Authenticate)

	/* Private routes */
	private := r.Group("/api/v1")
	private.Use(middlewares.JWT())
	{
		private.GET("/", endpoints.Home)
	}

	return r
}