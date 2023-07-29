package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrBadRequest          = ErrorResponse{"Bad Request"}
	ErrInternalServerError = ErrorResponse{"Internal Server Error"}
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrBadCredenrials      = errors.New("incorrect login or password")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)
