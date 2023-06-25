package repository

import (
	"context"
	"errors"

	"github.com/temryakov/go-backend-snippet-app/domain"

	"gorm.io/gorm"
)

type bookRepository struct {
	database   *gorm.DB
	collection string
}

func NewBookRepository(database *gorm.DB, collection string) domain.BookRepository {
	return &bookRepository{
		database:   database,
		collection: collection,
	}
}

func (r *bookRepository) FetchByID(ctx context.Context, id uint) (domain.Book, error) {

	var book domain.Book

	if err := r.database.WithContext(ctx).First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Book{}, err
		}
		return domain.Book{}, err
	}
	return book, nil
}

func (r *bookRepository) Fetch(ctx context.Context) ([]domain.Book, error) {

	var books []domain.Book

	if err := r.database.WithContext(ctx).Find(&books).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domain.Book{}, err
		}
		return []domain.Book{}, err
	}
	return books, nil
}

func (r *bookRepository) Save(ctx context.Context, book *domain.Book) error {

	return r.database.WithContext(ctx).Save(&book).Error
}

func (r *bookRepository) Delete(ctx context.Context, id uint) error {

	var book domain.Book

	if err := r.database.WithContext(ctx).Delete(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
