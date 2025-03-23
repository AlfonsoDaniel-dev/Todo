package domain

import (
	"errors"
	"strings"
)

type email struct {
	Value string
}

func newEmail(value string) (*email, error) {
	if value == "" || !strings.Contains(value, "@") {
		return nil, errors.New("value is invalid")
	}
	return &email{
		Value: value,
	}, nil
}
