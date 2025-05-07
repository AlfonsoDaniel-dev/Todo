package app

import "todoApp-backend/src/internal/domain"

type AppService struct {
	Repository domain.AppRepository
}

func NewAppService(repo domain.AppRepository) *AppService {
	return &AppService{
		Repository: repo,
	}
}
