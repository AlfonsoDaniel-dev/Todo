package user

import (
	"errors"
	"github.com/google/uuid"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/DTO"
)

func (S *UserServices) Login(loginDTO DTO.LoginDTO) (string, error) {
	if loginDTO.Email == "" || loginDTO.Password == "" {
		return "", errors.New("Username or Password is empty")
	}

	// TODO Login proccess
	// 1. get password from db with the email
	// 2. compare form password with db password
	// 3. if it's ok generate jwt and send it through client request, if it's not ok we will return a 401 code
	//  verify request
	// 1. request will pass for auth middleware who will know it's a valid token
	// 2. if pass handler will extract email from the JWT and send it to the useService who need it

	return "", nil

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
