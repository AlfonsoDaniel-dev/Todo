package app

import (
	"errors"
	"todoApp-backend/src/internal/DTO"
	"todoApp-backend/src/internal/domain"
)

type UserServices struct {
	Repository domain.UserRepository
}

func NewUserServices(repo domain.UserRepository) *UserServices {
	return &UserServices{
		Repository: repo,
	}
}

func (S *UserServices) CreateUser(UserDTO *DTO.UserDTO) (*DTO.GetUser, error) {
	if UserDTO.Username == "" || UserDTO.Email == "" || UserDTO.Password == "" {
		return nil, errors.New("username or email or password is empty")
	}

	User, err := domain.NewUser(UserDTO.Username, UserDTO.Email, UserDTO.Password)
	if err != nil {
		return nil, err
	}

	err = S.Repository.Save(&User)
	if err != nil {
		return nil, err
	}

	savedUserData := DTO.GetUser{
		Id:        User.Id,
		Username:  User.Username,
		Email:     User.Email.Value,
		CreatedAt: User.CreatedAt,
	}

	return &savedUserData, nil
}
