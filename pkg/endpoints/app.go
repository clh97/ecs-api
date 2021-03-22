package endpoints

import (
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
	userID, err := utils.GetUserIDFromContext(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, constants.HTTPErrorResponse(err, "Unauthorized", ""))
		return
	}

	payload := dtos.AppCreation{}

	// Binding
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	// Validation
	if err := validate.Struct(payload); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
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

// DeleteApp is the handler for the app removal endpoint
func DeleteApp(c *gin.Context) {
	payload := dtos.AppDelete{}

	// Binding
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	// Validation
	if err := validate.Struct(payload); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
			return
		}
		log.Fatal(err)
		return
	}

	// Service
	result, svcErr := services.DeleteApp(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// GetApps is the handler for the app listing endpoint
func GetApps(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, constants.HTTPErrorResponse(err, "Unauthorized", ""))
		return
	}

	result, svcErr := services.GetApps(userID)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}
