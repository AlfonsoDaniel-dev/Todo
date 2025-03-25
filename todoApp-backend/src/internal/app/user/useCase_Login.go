package user

import (
	"errors"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/DTO"
)

func (S *UserServices) Login(loginDTO DTO.LoginDTO) (string, error) {
	if loginDTO.Email == "" || loginDTO.Password == "" {
		return "", errors.New("Username or Password is empty")
	}

	// Login proccess
	// 1. get password from db with the email
	// 2. compare form password with db password
	// 3. if it's ok generate jwt and send it through client request, if it's not ok we will return a 401 code
	//  verify request
	// 1. request will pass for auth middleware who will know it's a valid token
	// 2. if pass handler will extract email from the JWT and send it to the useService who need it

	return "", nil

}
