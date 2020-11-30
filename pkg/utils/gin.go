package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext gets user from Gin context and returns as interface
func GetUserIDFromContext(c *gin.Context) int {
	claimsInterface, _ := c.Get("Claims")
	claims := claimsInterface.(jwt.MapClaims)

	userIDInterface := claims["user_id"]
	userIDFloat := userIDInterface.(float64)
	userID := int(userIDFloat)

	return userID
}
