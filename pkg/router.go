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
	public.GET("/page/:app-url-id/:page-id", endpoints.GetPage)
	public.GET("/page/:app-url-id", endpoints.GetPages)
	public.POST("/comment/:app-url-id/:page-id", endpoints.CreateComment)
	public.GET("/comment/:app-url-id/:page-id", endpoints.GetComments)

	/* Private routes */
	private := r.Group("/api/v1")
	private.Use(middlewares.JWT())
	{
		private.POST("/app", endpoints.CreateApp)
		private.GET("/app", endpoints.GetApps)
		private.DELETE("/app", endpoints.DeleteApp)
		private.POST("/page/:app-url-id", endpoints.CreatePage)
	}

	return r
}
