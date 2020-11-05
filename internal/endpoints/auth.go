package endpoints

import (
	"fmt"
	"net/http"

	"github.com/clh97/ecs/internal/dtos"
	"github.com/clh97/ecs/internal/network"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Auth is the handler for the user endpoint
func Auth(c *gin.Context) {
	var payload dtos.Login

	if err := c.ShouldBind(&payload); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, network.JSONValidationError(fieldErr, fmt.Sprintf("%s %s", fieldErr.Field(), fieldErr.Tag())))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": &payload})
}

func init() {
	validate = validator.New()
}
