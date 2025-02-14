package models

import (
	"github.com/google/uuid"
	"time"
)

type Task[T ownership] struct {
	Id        uuid.UUID
	Owner     T
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedBy User
	UpdatedAt time.Time
}
