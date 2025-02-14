package user

import (
	"errors"
	"github.com/google/uuid"
	"time"
	"todoApp-backend/src/helpers"
	"todoApp-backend/src/models"
)

func (U *User) CreateNewUser(Name, Email, Password string) error {

	userEncryptedPassword, err := helpers.EncryptPassword(Password)
	if err != nil {
		return err
	}

	NewUser := models.User{
		Id:       uuid.New(),
		UserName: Name,
		Email:    Email,
		Password: userEncryptedPassword,
		Created:  time.Now(),
	}

	err = U.DataInterface.CreateUserOnDB(NewUser)
	if err != nil {
		return err
	}
	return nil
}

func (U *User) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		err := errors.New("email is empty")
		return models.User{}, err
	}

	user, err := U.DataInterface.DBGetUserDataByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (U *User) GetUserById(id uuid.UUID) (models.User, error) {
	if id == uuid.Nil {
		return models.User{}, errors.New("id is empty")
	}

	user, err := U.DataInterface.DBGetUserDataById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (U *User) GetUserIdByEmail(Email string) (uuid.UUID, error) {
	if Email == "" {
		return uuid.UUID{}, errors.New("email cannot be empty")
	}

	id, err := U.DataInterface.DBGetUserIdByEmail(Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
