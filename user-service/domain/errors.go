package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	MessageBadRequest          = ErrorResponse{"Bad Request"}
	MessageInternalServerError = ErrorResponse{"Internal Server Error"}
	ErrUserAlreadyExists       = errors.New("user already exists")
)
