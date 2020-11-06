package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clh97/ecs/internal/constants"
	"github.com/clh97/ecs/internal/dtos"
	"github.com/clh97/ecs/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Authenticate is the handler for the user authentication endpoint
func Authenticate(c *gin.Context) {
	payload := dtos.Login{}

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
	result, svcErr := services.AuthenticateUser(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// Create is the handler for user creation endpoint
func Create(c *gin.Context) {
	payload := dtos.AccountCreation{}

	// Binding
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Bad request", ""))
		return
	}

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
	result, svcErr := services.CreateUser(payload)

	// Service error check
	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	// Success
	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

func init() {
	validate = validator.New()
}
