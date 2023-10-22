package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Email    string
	Name     string
	Password string
	gorm.Model
}

type UserUsecase interface {
	FetchByID(c context.Context, id uint) (*User, error)
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	FetchByEmail(c context.Context, email string) (*User, error)
	FetchByID(c context.Context, id uint) (*User, error)
}
