package domain

import "context"

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
}
