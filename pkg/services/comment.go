package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/store"
)

// CreateComment implements comment creation functionality
func CreateComment(payload dtos.CommentCreation) (*constants.ServiceResult, *constants.ServiceError) {
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

	return nil, nil
}
