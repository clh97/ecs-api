package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/types"
	"github.com/clh97/ecs/store"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// CreateApp implements app creation service
func CreateApp(payload dtos.AppCreation, userID int) (*constants.ServiceResult, *constants.ServiceError) {
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

	var lastInsertedURLID string
	err = db.QueryRowx("INSERT INTO ecs_app (name, url, owner_id) VALUES ($1, $2, $3) RETURNING url_id", payload.Name, payload.URL, userID).Scan(&lastInsertedURLID)

	if err != nil {
		svcError := new(constants.ServiceError)

		if err, ok := err.(*pq.Error); ok {
			if constants.IsErrUniqueViolation(err) {
				svcError.HTTPStatus = http.StatusConflict
				svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
					Error:     errors.New("Unable to execute sql transaction"),
					Message:   "Email already registered",
					Timestamp: time.Now(),
					Success:   false,
				}
			}
		} else {
			svcError.HTTPStatus = http.StatusInternalServerError
			svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
				Error:     errors.New("Unable to execute sql transaction"),
				Message:   "Internal server error",
				Timestamp: time.Now(),
				Success:   false,
			}
		}
		return nil, svcError
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data: gin.H{
			"UrlID": lastInsertedURLID,
		},
		Message:   "Successfully created app",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated

	return svcResult, nil
}

// DeleteApp implements app delete functionality
func DeleteApp(payload dtos.AppDelete) (*constants.ServiceResult, *constants.ServiceError) {
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

	app := types.App{}

	db.QueryRowx("SELECT * FROM ecs_app WHERE url_id = $1", payload.URLID).StructScan(&app)

	if app.IsEmpty() {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Could not find app to delete"),
			Message:   "Could not find the specified app",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusNotFound

		return nil, svcError
	}

	_, err = db.Queryx("DELETE FROM ecs_app WHERE url_id = $1", payload.URLID)

	if err != nil {
		svcError := new(constants.ServiceError)

		if err, ok := err.(*pq.Error); ok {
			svcError.HTTPStatus = http.StatusInternalServerError
			svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
				Error:     err,
				Message:   "Internal server error",
				Timestamp: time.Now(),
				Success:   false,
			}
			return nil, svcError
		}
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data:      nil,
		Message:   "Successfully removed app",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusOK

	return svcResult, nil
}

// GetApps implements app listing functionality
func GetApps(userID int) (*constants.ServiceResult, *constants.ServiceError) {
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

	apps := []types.App{}

	err = db.Select(&apps, "SELECT name, url, url_id, created_at FROM ecs_app WHERE owner_id = $1", userID)

	if err != nil {
		svcError := new(constants.ServiceError)

		if err, ok := err.(*pq.Error); ok {
			if constants.IsErrUniqueViolation(err) {
				svcError.HTTPStatus = http.StatusConflict
				svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
					Error:     errors.New("Unable to execute sql transaction"),
					Message:   "Email already registered",
					Timestamp: time.Now(),
					Success:   false,
				}
			}
		} else {
			svcError.HTTPStatus = http.StatusInternalServerError
			svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
				Error:     errors.New("Unable to execute sql transaction"),
				Message:   "Internal server error",
				Timestamp: time.Now(),
				Success:   false,
			}
		}
		return nil, svcError
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data:      apps,
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusOK

	return svcResult, nil
}
