package app

import (
	"errors"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/DTO"
)

type UserServices struct {
	Repository domain.UserRepository
}

func NewUserServices(repo domain.UserRepository) *UserServices {
	return &UserServices{
		Repository: repo,
	}
}

func (S *UserServices) Login(loginDTO DTO.LoginDTO) (string, error) {
	if loginDTO.Email == "" || loginDTO.Password == "" {
		return "", errors.New("Username or Password is empty")
	}

	// get user Email coincides with the password
	// 1. get password from db with the email
	// 2. compare form password with db password
	// 3. if it's ok generate jwt and send it through client request, if it's not ok we will return a 401 code
	//  verify request
	// 1. request will pass for auth middleware who will know it's a valid token
	// 2. if pass handler will extract email from the JWT and send it to the useService who need it

	return "", nil

}

func (S *UserServices) CreateUser(UserDTO *DTO.UserDTO) error {
	if UserDTO.Username == "" || UserDTO.Email == "" || UserDTO.Password == "" {
		return errors.New("username or email or password is empty")
	}

	User, err := domain.NewUser(UserDTO.Username, UserDTO.Email, UserDTO.Password)
	if err != nil {
		return err
	}

	err = S.Repository.Save(&User)
	if err != nil {
		return err
	}

	return nil
}
