package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext gets user from Gin context and returns as interface
func GetUserIDFromContext(c *gin.Context) (int, *constants.ServiceError) {
	claimsInterface, exists := c.Get("Claims")

	if !exists {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("User is not authenticated"),
			Message:   "Unauthorized",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusUnauthorized

		return 0, svcError
	}

	claims := claimsInterface.(jwt.MapClaims)

	userIDInterface := claims["user_id"]
	userIDFloat := userIDInterface.(float64)
	userID := int(userIDFloat)

	return userID, nil
}
