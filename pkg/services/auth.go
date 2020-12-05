package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/clh97/ecs/pkg/constants"
	"github.com/clh97/ecs/pkg/dtos"
	"github.com/clh97/ecs/pkg/types"
	"github.com/clh97/ecs/pkg/utils"
	"github.com/clh97/ecs/store"
	"github.com/lib/pq"
)

// CreateUser implements user creation service
func CreateUser(payload dtos.AccountCreation) (*constants.ServiceResult, *constants.ServiceError) {
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

	hashed, err := utils.GenerateFromPassword(payload.Password)

	if err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to generate hash from password"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError

		return nil, svcError
	}

	tx := db.MustBegin()

	query := "INSERT INTO ecs_user (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)"

	if _, err := tx.Exec(query, payload.FirstName, payload.LastName, payload.Email, hashed); err != nil {
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

	tx.Commit()

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data:      nil,
		Message:   "Successfully created user",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusCreated
	return svcResult, nil
}

// AuthenticateUser returns authentication token according to user credentials
func AuthenticateUser(payload dtos.Login) (*constants.ServiceResult, *constants.ServiceError) {
	var token string
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

	user := types.User{}

	if err = db.Get(&user, "SELECT id, email, password FROM ecs_user WHERE email = $1", payload.Email); err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to query database for user"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError
		return nil, svcError
	}

	if user == (types.User{}) {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Could not find user with specified email"),
			Message:   "Invalid username or password",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusUnauthorized
		return nil, svcError
	}

	match, err := utils.ComparePasswordAndHash(payload.Password, user.Password)

	if err != nil {
		svcError := new(constants.ServiceError)
		svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
			Error:     errors.New("Unable to compare password with hash"),
			Message:   "Internal server error",
			Timestamp: time.Now(),
			Success:   false,
		}
		svcError.HTTPStatus = http.StatusInternalServerError
		return nil, svcError
	}

	if match {
		token, err = utils.CreateToken(user.ID)

		if err != nil {
			svcError := new(constants.ServiceError)
			svcError.HTTPErrorResponse = constants.THTTPErrorResponse{
				Error:     errors.New("Unable to generate token for user"),
				Message:   "Internal server error",
				Timestamp: time.Now(),
				Success:   false,
			}
			svcError.HTTPStatus = http.StatusInternalServerError
			return nil, svcError
		}
	}

	svcResult := new(constants.ServiceResult)
	svcResult.HTTPResponse = constants.THTTPResponse{
		Data: map[string]interface{}{
			"Token": token,
		},
		Message:   "Successfully authenticated user",
		Timestamp: time.Now(),
		Success:   true,
	}
	svcResult.HTTPStatus = http.StatusOK

	return svcResult, nil
}
