package domain

import "gorm.io/gorm"

type User struct {
	Email    string
	Name     string
	Password string
	gorm.Model
}
