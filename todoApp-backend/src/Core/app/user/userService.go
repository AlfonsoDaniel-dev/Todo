package user

import "todoApp-backend/src/Core/domain"

type UserServices struct {
	Repository domain.UserRepository
}

func NewUserServices(repo domain.UserRepository) *UserServices {
	return &UserServices{
		Repository: repo,
	}
}
