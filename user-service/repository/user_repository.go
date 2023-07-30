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

func (r *userRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	err := r.database.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FetchByID(ctx context.Context, ID uint) (*domain.User, error) {

	var user domain.User

	err := r.database.WithContext(ctx).First(&user, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
