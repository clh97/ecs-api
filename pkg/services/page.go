package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/types"
	"github.com/clh97/ecs/store"
	"github.com/gin-gonic/gin"
)

// CreatePage implements page creation functionality
func CreatePage(payload dtos.PageCreation) (*constants.ServiceResult, *constants.ServiceError) {
	db, err := store.CreateDBInstance()

	defer db.Close()

	if err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to create db instance"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError

		return nil, svcError
	}

	page := types.Page{
		Title:    payload.Title,
		URL:      payload.URL,
		AppURLID: payload.AppURLID,
	}

	result, err := db.NamedExec("INSERT INTO ecs_page (title, url, app_id) VALUES (:title, :url, :appurlid)", page)

	fmt.Println(result, page.Title)

	if err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to execute sql statement"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError

		return nil, svcError
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data:      nil,
		Message:   "Successfully created page",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated

	return svcResult, nil
}

// GetPage implements page retrieval functionality
func GetPage(payload dtos.PageGet) (*constants.ServiceResult, *constants.ServiceError) {
	db, err := store.CreateDBInstance()

	defer db.Close()

	if err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to create db instance"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError

		return nil, svcError
	}

	page := types.Page{}

	db.QueryRowx("SELECT * FROM ecs_page WHERE page_id = $1 AND app_id = $2", payload.PageID, payload.AppURLID).Scan(&page)

	fmt.Println(page)

	if page.IsEmpty() {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     constants.ErrNotFound,
			Message:   "Page not found",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusNotFound

		return nil, svcError
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data: gin.H{
			"Page": page,
		},
		Message:   "Successfully retrieved page",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusOK

	return svcResult, nil
}
