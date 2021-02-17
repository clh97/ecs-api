package endpoints

import (
	"net/http"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/services"
	"github.com/gin-gonic/gin"
)

/*
Page needs to be associated with some app
ECS <- App <- Page <- User comment
*/

// CreatePage creates a page in a app, identified by urlid
func CreatePage(c *gin.Context) {
	payload := dtos.PageCreation{}

	// Binding
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	// Validation
	if err := validate.Struct(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	result, svcErr := services.CreatePage(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// GetPage returns a single page by its urlid and pageid
func GetPage(c *gin.Context) {
	urlID := c.Param("app-url-id")
	pageID := c.Param("page-id")

	payload := dtos.PageGet{AppURLID: urlID, PageID: pageID}

	// Validation
	if err := validate.Struct(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	// Service
	result, svcErr := services.GetPage(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}

// GetPages returns a list o pages by its url id
func GetPages(c *gin.Context) {
	urlID := c.Param("app-url-id")

	payload := dtos.PagesGet{AppURLID: urlID}

	// Validation
	if err := validate.Struct(payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
		return
	}

	// Service
	result, svcErr := services.GetPages(payload)

	if svcErr != nil {
		c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
		return
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}
