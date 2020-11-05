package middlewares

import (
	"net/http"

	"github.com/clh97/ecs/internal/network"
	"github.com/gin-gonic/gin"
)

// JWT is the middleware responsible for verifying authentication
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, network.MissingAuthHeaderError)
		}

		c.Next()
	}
}
