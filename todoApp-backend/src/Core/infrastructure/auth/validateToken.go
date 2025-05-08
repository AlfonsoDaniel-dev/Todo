package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"todoApp-backend/src/Core/domain"
)

func ValidateToken(tokenString string) (domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, verifyFunc)
	if err != nil {
		return domain.Claims{}, err
	}

	if !token.Valid {
		fmt.Println("Fallo en auth")
		return domain.Claims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*domain.Claims)
	if !ok {
		return domain.Claims{}, errors.New("No se pudieron obtener los tokens")
	}

	return *claims, nil
}

func verifyFunc(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
