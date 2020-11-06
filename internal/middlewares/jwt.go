package middlewares

import (
	"net/http"
	"strings"

	"github.com/clh97/ecs/internal/constants"
	"github.com/clh97/ecs/internal/utils"
	"github.com/gin-gonic/gin"
)

// JWT is the middleware responsible for verifying authentication
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(nil, "Missing Authorization header", ""))
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) <= 1 || !utils.ValidateToken(parts[1]) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(nil, "Invalid token", ""))
			return
		}

		c.Next()
	}
}
