package domain

import "errors"

var (
	ErrReviewNotFound = errors.New("review not found")
	ErrReviewIsExist  = errors.New("review is exist")
)
