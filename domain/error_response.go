package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type SnippetErr struct {
	Error ErrorResponse
}

var SnippetNotFound = ErrorResponse{Message: "Snippet is not found. =("}
var MessageBadRequest = ErrorResponse{Message: "Bad Request =/"}
var MessageInternalServerError = ErrorResponse{Message: "...Oops."}
