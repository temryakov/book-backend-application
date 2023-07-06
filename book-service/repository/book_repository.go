package repository

import (
	"context"

	"github.com/temryakov/go-backend-book-app/book-service/domain"

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

func (r *bookRepository) FetchBookByID(ctx context.Context, id uint) (*domain.Book, error) {

	var book *domain.Book

	if err := r.database.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) FetchBooks(ctx context.Context) (*[]domain.Book, error) {

	var books *[]domain.Book

	if err := r.database.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) CreateBook(ctx context.Context, book *domain.Book) error {

	return r.database.WithContext(ctx).Save(&book).Error
}
func (r *bookRepository) UpdateBook(ctx context.Context, book *domain.Book, Model *domain.Book) error {

	r.database.WithContext(ctx).Model(&Model).Updates(&book)
	return nil
}

func (r *bookRepository) DeleteBook(ctx context.Context, bookId uint) error {

	var book domain.Book

	if err := r.database.WithContext(ctx).Delete(&book, bookId).Error; err != nil {
		return err
	}
	return nil
}
