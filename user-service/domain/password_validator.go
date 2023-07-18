package domain

import (
	"errors"
)

var (
	ErrShortPassword = errors.New("password must be greater than 7")
	ErrLongPassword  = errors.New("password must be lower than 20")
)

func PasswordValidation(password string) error {
	if len(password) < 7 {
		return ErrShortPassword
	}
	if len(password) > 20 {
		return ErrLongPassword
	}
	return nil
}
