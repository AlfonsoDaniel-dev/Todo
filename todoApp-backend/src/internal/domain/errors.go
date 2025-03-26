package domain

import "errors"

var (
	UserAlreadyExists  = errors.New("User already exists")
	ErrTokenIsNotValid = errors.New("Token is not valid")
	ErrNotFound        = errors.New("User not found")
)
