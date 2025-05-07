package user

import (
	"errors"
	"github.com/google/uuid"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/DTO"
)

func (S *UserServices) Login(loginDTO DTO.LoginDTO) error {
	if loginDTO.Email == "" || loginDTO.Password == "" {
		return domain.ErrInvalidLoginForm
	}

	// 1. verify User exists
	exists, err := S.Repository.CheckUserExists(loginDTO.Email)
	if !errors.Is(err, domain.ErrNotFound) { // check if it is an error while getting the data from DB
		return err
	} else if !exists {
		return domain.ErrNotFound
	}
	// get password from db with the email

	userId, err := S.Repository.GetIdByEmail(loginDTO.Email)
	if err != nil {
		return err
	}

	userPassword, err := S.Repository.GetUserPassword(userId)
	if err != nil {
		return err
	}

	// compare form password with db password

	ok := domain.ComparePassword(loginDTO.Password, userPassword)
	if !ok {
		return domain.ErrWrongPassword
	}

	return nil
}

func (S *UserServices) OAuthLogin(userName, email string) error {
	if userName == "" || email == "" {
		return errors.New("username or email are required")
	}

	id, err := S.Repository.GetIdByEmail(email)
	if id != uuid.Nil && err == nil {

		return domain.UserAlreadyExists

	} else if errors.Is(err, domain.ErrNotFound) {
		userPassword := domain.GeneratePassword()

		user, err := domain.NewUser(userName, email, userPassword)
		if err != nil {
			return err
		}

		err = S.Repository.Save(&user)
		if err != nil {
			return err
		}

		return domain.ErrNotFound
	}

	return err
}
