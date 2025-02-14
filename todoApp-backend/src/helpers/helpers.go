package helpers

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func TimeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{
		Time: t,
	}
	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

func IntToNull(number int64) sql.NullInt64 {
	null := sql.NullInt64{
		Int64: number,
	}

	if null.Int64 != 0 {
		null.Valid = true
	}

	return null
}

func StringToNull(char string) sql.NullString {
	null := sql.NullString{
		String: char,
	}

	if null.String != "" {
		null.Valid = true
	}

	return null
}

func EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return "", errors.New("couldn't encrypt password. error: " + err.Error())
	}

	return string(encryptedPassword), nil
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("error trying to compare password. error: %s\n", err)
	}

	return true
}
