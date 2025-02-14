package user

import (
	"github.com/google/uuid"
	"todoApp-backend/src/models"
)

type DataInterface interface {
	CreateUserOnDB(user models.User) error

	DBGetUserIdByEmail(email string) (uuid.UUID, error)
	DBGetUserDataById(id uuid.UUID) (models.User, error)
	DBGetUserDataByEmail(email string) (models.User, error)
}
