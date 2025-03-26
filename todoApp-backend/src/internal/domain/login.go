package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"os"
	"time"
)

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewLogin(userName, email, password string) *Login {
	return &Login{
		Username: userName,
		Email:    email,
		Password: password,
	}
}

func NewClaims(email string) Claims {

	return Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}
}
