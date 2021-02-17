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

	// because it's a public comment, we need to set anonymous as true
	payload.Anonymous = true

	// var lastInsertedURLID string
	// err = db.QueryRowx("INSERT INTO ecs_app (name, url, owner_id) VALUES ($1, $2, $3) RETURNING url_id", payload.Name, payload.URL, userID).Scan(&lastInsertedURLID)

	result, err := db.NamedExec("INSERT INTO ecs_comment (app_id, page_id, content, content_format, anon_username, anon) VALUES (:appurlid, :pageid, :body, :format, :username, :anonymous)", payload)
	fmt.Println(result)

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Message:   "Successfully created comment",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated

	return svcResult, nil
}
