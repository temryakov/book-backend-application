package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type BookErr struct {
	Error ErrorResponse
}

var (
	ReviewNotFound             = ErrorResponse{Message: "Review is not found. =("}
	MessageBadRequest          = ErrorResponse{Message: "Bad Request =/"}
	MessageInternalServerError = ErrorResponse{Message: "...Oops."}
	MessageForbidden           = ErrorResponse{Message: "You don't have permission. %("}
)
