package domain

import "errors"

var (
	UserAlreadyExists            = errors.New("User already exists")
	ErrTokenIsNotValid           = errors.New("Token is not valid")
	ErrNotEnoughOrValidArguments = errors.New("Not enough arguments")
	ErrIdIsNotValid              = errors.New("Id is not valid")
	ErrNotFound                  = errors.New("User not found")
	ErrInvalidToken              = errors.New("invalid token")
	ErrWrongPassword             = errors.New("wrong password")
	ErrInvalidLoginForm          = errors.New("username or Password are empty")
)
