package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Email     email
	Password  password
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserRepository interface {
	Save(user *User) error
	GetUserData(id uuid.UUID) (User, error)
	GetIdByEmail(email string) (uuid.UUID, error)
	GetEmailById(id uuid.UUID) (string, error)
	GetUserPassword(id uuid.UUID) (string, error)
	UpdateName(NewName string, id uuid.UUID) error
	UpdateEmail(Email string, id uuid.UUID) error
	UpdatePassword(Password string, id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
	CheckUserExists(email string) (bool, error)
}

func NewUser(UserName, Email string, Password string) (User, error) {
	if UserName == "" || Email == "" || len(Password) == 0 {
		return User{}, errors.New("UserName or Email or Password is empty, All Data is Required")
	}

	createdAt := time.Now().UTC()

	userNewPassword, err := newPassword(Password)
	if err != nil {
		return User{}, errors.New("Error creating new password: " + err.Error())
	}
	userNewEmail, err := newEmail(Email)
	if err != nil {
		return User{}, errors.New("Error creating new email: " + err.Error())
	}

	return User{
		Id:        uuid.New(),
		Username:  UserName,
		Email:     *userNewEmail,
		Password:  *userNewPassword,
		CreatedAt: createdAt,
	}, nil
}
