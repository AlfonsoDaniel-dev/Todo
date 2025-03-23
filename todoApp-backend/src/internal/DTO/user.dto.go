package DTO

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUserDTO(userName, password, email string) (*UserDTO, error) {

	if userName == "" || password == "" || email == "" {
		return nil, errors.New("username or password or email is empty")
	}

	return &UserDTO{
		Username: userName,
		Password: password,
		Email:    email,
	}, nil
	
}

type GetUser struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
