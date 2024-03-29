package repository

import (
	"context"

	"book-service/domain"

	"gorm.io/gorm"
)

type bookRepository struct {
	database *gorm.DB
}

func NewBookRepository(database *gorm.DB) domain.BookRepository {
	return &bookRepository{
		database: database,
	}
}

func (r *bookRepository) FetchBookByID(ctx context.Context, id int) (*domain.Book, error) {

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

func (r *bookRepository) DeleteBook(ctx context.Context, bookId int) error {

	var book domain.Book

	if err := r.database.WithContext(ctx).Delete(&book, bookId).Error; err != nil {
		return err
	}
	return nil
}
