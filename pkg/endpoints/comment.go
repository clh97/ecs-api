package endpoints

import (
	"log"
	"net/http"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreatePublicComment is the handler for the public comment creation endpoint
func CreatePublicComment(c *gin.Context) {
	urlID := c.Param("app-url-id")
	pageID := c.Param("page-id")

	payload := dtos.CommentCreation{
		AppURLID: urlID,
		PageID:   pageID,
		UserID:   0,
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

	// Service
	result, svcErr := services.CreatePublicComment(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// CreatePrivateComment is the handler for the private comment creation endpoint
func CreatePrivateComment(c *gin.Context) {
	userID, err := services.GetUserIDFromContext(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, constants.HTTPErrorResponse(err, "Unauthorized", ""))
		return
	}

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

	// Service
	result, svcErr := services.CreatePrivateComment(payload)

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
