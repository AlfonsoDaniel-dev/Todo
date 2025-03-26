package auth

import (
	"github.com/golang-jwt/jwt"
	"todoApp-backend/src/internal/domain"
)

func GenerateToken(email string) (string, error) {
	claim := domain.NewClaims(email)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
