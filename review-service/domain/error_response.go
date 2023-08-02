package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type BookErr struct {
	Error ErrorResponse
}

var BookNotFound = ErrorResponse{Message: "Book is not found. =("}
var MessageBadRequest = ErrorResponse{Message: "Bad Request =/"}
var MessageInternalServerError = ErrorResponse{Message: "...Oops."}
