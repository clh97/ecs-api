package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/utils"
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

		setUserContext(c, parts[1])

		if len(parts) <= 1 || !utils.ValidateToken(parts[1]) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(nil, "Invalid token", ""))
			return
		}

		c.Next()
	}
}

// setUserContext sets up identification as header in context
func setUserContext(c *gin.Context, token string) {
	// c.Header("UserId", "1")
	claims, err := utils.DecodeToken(token)

	if err != nil {
		fmt.Printf("Could not set user context: %v\n", err)
	}

	c.Set("Claims", claims)
}
