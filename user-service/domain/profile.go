package domain

import "context"

type Profile struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ProfileUsecase interface {
	FetchByID(c context.Context, id uint) (*User, error)
}
