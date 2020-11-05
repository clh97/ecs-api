package services

import (
	"github.com/clh97/ecs/internal/dtos"
	"github.com/clh97/ecs/internal/utils"
	"github.com/clh97/ecs/store"
)

// Services will return false when something wrong happens

// CreateUser implements user creation service
func CreateUser(payload dtos.AccountCreation) (bool, error) {
	db, err := store.CreateDBInstance()

	defer db.Close()

	if err != nil {
		return false, err
	}

	hashed, err := utils.GenerateFromPassword(payload.Password)

	if err != nil {
		return false, err
	}

	tx := db.MustBegin()

	_, err = tx.Exec("INSERT INTO ecs_user (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", payload.FirstName, payload.LastName, payload.Email, hashed)

	tx.Commit()

	if err != nil {
		return false, err
	}

	return true, nil
}

// AuthenticateUser logs returns authentication token according to payload
func AuthenticateUser(payload *dtos.Login) (bool, error) {
	return true, nil
}
