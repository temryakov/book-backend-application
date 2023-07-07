package repository

import (
	"context"

	"github.com/temryakov/go-backend-book-app/user-service/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: database,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.database.WithContext(ctx).Save(&user).Error
}
