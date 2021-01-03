package endpoints

import (
	"fmt"
	"strconv"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/services"
	"github.com/gin-gonic/gin"
)

/*
Page needs to be associated with some app
ECS <- App <- Page <- User comment
*/

/*
	Page is the handler for the page endpoint
	--
	That means it's able to deal with page creation if it doesn't exist
	and with page content listing if it does exist.
	---
	It works in this way to abstract the front end request into a single GET
	instead of a POST and GET
*/
func Page(c *gin.Context) {

	// payload := dtos.PageCreation{}

	// Binding
	// if err := c.ShouldBindJSON(&payload); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
	// 	return
	// }

	// // Validation
	// if err := validate.Struct(payload); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, constants.HTTPErrorResponse(err, "Validation/structure error", ""))
	// 	return
	// }

	appID, _ := c.Params.Get("app-id")
	pageIDStr, _ := c.Params.Get("page-id")
	fmt.Println("App id:", appID)
	fmt.Println("Page id:", pageIDStr)

	// IF THE PAGE EXISTS, RETURN CONTENT
	// IF NOT, CREATE PAGE

	pageID, _ := strconv.Atoi(pageIDStr)

	pageGetPayload := dtos.PageGet{AppURLID: appID, PageID: pageID}

	// Service
	result, svcErr := services.GetPage(pageGetPayload)

	if svcErr != nil {
		// This means the resource simply doesn't exist.
		// We can now create it and keep the user happy
		if errNotFound := svcErr.HTTPErrorResponse.Error == constants.ErrNotFound; errNotFound {
			pageCreatePayload := dtos.PageCreation{}

			// TODO: page creation stuff
		} else {
			c.AbortWithStatusJSON(svcErr.HTTPStatus, svcErr.HTTPErrorResponse)
			return
		}
	}

	c.JSON(result.HTTPStatus, result.HTTPResponse)
}
