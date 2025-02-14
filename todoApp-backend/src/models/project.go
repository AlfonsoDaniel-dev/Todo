package models

import (
	"github.com/google/uuid"
	"time"
)

type Project[T ownership] struct {
	Id            uuid.UUID
	Owner         T
	Name          string
	Description   string
	Tasks         []Task[T]
	Collaborators []User
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
