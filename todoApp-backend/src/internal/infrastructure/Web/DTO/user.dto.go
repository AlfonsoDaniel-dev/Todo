package DTO

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type GetUser struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateUserName struct {
	NewUserName string `json:"newUserName"`
}

type UpdateUserEmail struct {
	NewEmail string `json:"newEmail"`
}

type UpdateUserPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type DeleteUser struct {
	Password string `json:"oldPassword"`
}
