package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/services"
	"github.com/clh97/ecs/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateApp is the handler for the app creation endpoint
func CreateApp(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)

	payload := dtos.AppCreation{}

	// Binding
	err := c.ShouldBindJSON(&payload)

	// Validation
	err = validate.Struct(payload)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, fmt.Sprintf("%s is invalid", err.Field()), ""))
			return
		}
		log.Fatal(err)
		return
	}

	// Service
	result, svcErr := services.CreateApp(payload, userID)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// GetApps is the handler for the app list endpoint
func GetApps(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)

	result, svcErr := services.GetApps(userID)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}