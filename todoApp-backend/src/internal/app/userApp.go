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
