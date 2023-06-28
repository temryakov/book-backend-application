package repository

import (
	"context"
	"errors"

	"github.com/temryakov/go-backend-book-app/domain"

	"gorm.io/gorm"
)

type bookProgressRepository struct {
	database   *gorm.DB
	collection string
}

func NewBookProgressRepository(database *gorm.DB, collection string) domain.BookProgressRepository {
	return &bookProgressRepository{
		database:   database,
		collection: collection,
	}
}

func (r *bookProgressRepository) FetchByID(ctx context.Context, id uint) (domain.BookProgress, error) {

	var bookProgress domain.BookProgress

	if err := r.database.WithContext(ctx).First(&bookProgress, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.BookProgress{}, err
		}
		return domain.BookProgress{}, err
	}
	return bookProgress, nil
}
