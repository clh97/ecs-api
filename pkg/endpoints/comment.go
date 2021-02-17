package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateComment is the handler for the comment creation endpoint
func CreateComment(c *gin.Context) {
	var result *constants.ServiceResult
	var svcErr *constants.ServiceError

	userID, svcErr := services.GetUserIDFromContext(c)

	// if svcErr != nil {
	// 	c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
	// 	return
	// }

	urlID := c.Param("app-url-id")
	pageID := c.Param("page-id")

	payload := dtos.CommentCreation{
		AppURLID: urlID,
		PageID:   pageID,
		UserID:   userID,
	}

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

	// setting as nil to be sure it wont inherit from last definition
	svcErr = nil

	// Service
	if userID != 0 {
		fmt.Println("100% authenticated comment coming!")
	} else {
		fmt.Println("Non-authenticated comment coming!")
		result, svcErr = services.CreatePublicComment(payload)
	}

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// GetComments is the handler for the comment listing endpoint
func GetComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome Comment! =D"})
}
