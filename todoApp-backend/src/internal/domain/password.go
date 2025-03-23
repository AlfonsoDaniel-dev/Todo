package domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type password struct {
	Hash []byte
}

func newPassword(passwordString string) (*password, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordString), 6)
	if err != nil {
		return nil, errors.New("couldn't encrypt password. error: " + err.Error())
	}

	PasswordObj := password{
		Hash: encryptedPassword,
	}

	return &PasswordObj, nil
}

func comparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("error trying to compare password. error: %s\n", err)
	}

	return true
}
