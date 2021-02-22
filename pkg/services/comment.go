package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/store"
)

// CreatePublicComment implements comment creation functionality
func CreatePublicComment(payload dtos.CommentCreation) (*constants.ServiceResult, *constants.ServiceError) {
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

	// it's a public comment so we need to set anonymous as true
	payload.Anonymous = true
	payload.UserID = 1
	payload.Username = "<anonymous>"

	_, err = db.NamedExec("INSERT INTO ecs_comment (app_id, page_id, user_id, content, content_format, anon_username, anon) VALUES (:appurlid, :pageid, :userid, :body, :format, :username, :anonymous)", payload)

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
		Message:   "Successfully created unauthenticated comment",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated

	return svcResult, nil
}

// CreatePrivateComment implements comment creation functionality
func CreatePrivateComment(payload dtos.CommentCreation) (*constants.ServiceResult, *constants.ServiceError) {
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

	var username string
	err = db.QueryRow("SELECT username from ecs_user WHERE id = $1", payload.UserID).Scan(&username)

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

	// it's a private comment so we need to set anonymous as false
	payload.Anonymous = false
	payload.Username = username

	fmt.Println(payload)

	// check up if uses anon_username or username
	_, err = db.NamedExec("INSERT INTO ecs_comment (app_id, page_id, user_id, content, content_format, username, anon) VALUES (:appurlid, :pageid, :userid, :body, :format, :username, :anonymous)", payload)

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
		Message:   "Successfully created authenticated comment",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated

	return svcResult, nil
}
