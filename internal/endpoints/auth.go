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
	"github.com/lib/pq"
)

var validate *validator.Validate

// Authenticate is the handler for the user authentication endpoint
func Authenticate(c *gin.Context) {
	payload := dtos.Login{}

	// Data validation
	if err := c.ShouldBind(&payload); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.JSONValidationError(fieldErr, fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag())))
			return
		}
	}

	c.JSON(http.StatusOK, constants.Success(nil, "Successfully authenticated"))
}

// Create is the handler for user creation endpoint
func Create(c *gin.Context) {
	payload := dtos.AccountCreation{}

	// Binding
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.JSONValidationError(err, err.Error()))
		return
	}

	// Validation
	err = validate.Struct(payload)

	if err != nil {
		// Iterates through every validation errors and returns first one
		for _, err := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, constants.JSONValidationError(err, fmt.Sprintf("%s is invalid", err.Field())))
			return
		}
		log.Fatal(err)
		return
	}

	// Service
	_, err = services.CreateUser(payload)

	// Service error check
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if constants.IsErrUniqueViolation(err) {
				c.AbortWithStatusJSON(http.StatusConflict, constants.ResourceConflictError(err, "Email already registered"))
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.InternalError)
		return
	}

	// Success
	c.JSON(http.StatusCreated, constants.ResourceCreated(nil, "Successfully created resource"))
}

func init() {
	validate = validator.New()
}
