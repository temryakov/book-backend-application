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

type UserRepository interface {
	Create(c context.Context, user *User) error
}
