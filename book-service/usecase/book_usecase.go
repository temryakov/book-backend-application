package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/book-service/domain"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
	}
}

func (su *bookUsecase) FetchBookByID(c context.Context, bookId uint) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.FetchBookByID(ctx, bookId)
}

func (su *bookUsecase) FetchBooks(c context.Context) (*[]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.FetchBooks(ctx)
}

func (su *bookUsecase) CreateBook(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.CreateBook(ctx, book)
}
func (su *bookUsecase) UpdateBook(c context.Context, book *domain.Book, bookId uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	model, err := su.bookRepository.FetchBookByID(ctx, bookId)
	if err != nil {
		return err
	}
	return su.bookRepository.UpdateBook(ctx, book, model)
}

func (su *bookUsecase) DeleteBook(c context.Context, bookId uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.DeleteBook(ctx, bookId)
}
