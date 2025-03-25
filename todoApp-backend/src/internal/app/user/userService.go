package user

import "todoApp-backend/src/internal/domain"

type UserServices struct {
	Repository domain.UserRepository
}

func NewUserServices(repo domain.UserRepository) *UserServices {
	return &UserServices{
		Repository: repo,
	}
}
