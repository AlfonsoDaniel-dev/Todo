package models

import (
	"github.com/google/uuid"
	"time"
)

type Organization struct {
	Id            uuid.UUID
	OwnerId       uuid.UUID
	Name          string
	Description   string
	Collaborators []User
	Projects      []Project[Organization]
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
