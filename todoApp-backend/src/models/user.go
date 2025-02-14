package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id            uuid.UUID
	UserName      string
	Email         string
	Password      string
	Organizations []Organization
	MemberOf      []Organization
	Projects      []Project[User]
	Created       time.Time
	Updated       time.Time
	DeletedAt     time.Time
}
