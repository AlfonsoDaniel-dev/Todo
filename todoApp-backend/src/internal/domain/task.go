package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedBy User
	UpdatedAt time.Time
}

func NewTask(ownerId uuid.UUID, title, Body string) (*Task, error) {
	if title == "" || Body == "" {
		return nil, errors.New("title or body is empty")
	}

	return &Task{
		Id:      uuid.New(),
		OwnerId: ownerId,
		Title:   title,
		Body:    Body,
	}, nil
}
