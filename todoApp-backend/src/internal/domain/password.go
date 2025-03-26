package domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
)

type password struct {
	Hash []byte
}

func GeneratePassword() string {
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z", "A",
		"B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z", "0", "1",
		"2", "3", "4", "5", "6", "7", "8", "9", "!",
		"@", "#", "$", "%", "^", "&", "*", "(", ")",
		"_", "-", "+", "=", "<", ">", "/", "?", "[",
	}

	passwordLenght := 25

	var Password string
	for i := 0; i < passwordLenght; i++ {
		RandomIndex := rand.Intn(len(chars))

		Password += chars[RandomIndex]
	}

	return Password
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
