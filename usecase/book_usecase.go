package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/domain"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookRepository {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
	}
}

func (su *bookUsecase) FetchByID(c context.Context, bookId uint) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.FetchByID(ctx, bookId)
}

func (su *bookUsecase) Fetch(c context.Context) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Fetch(ctx)
}

func (su *bookUsecase) Save(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Save(ctx, book)
}

func (su *bookUsecase) Delete(c context.Context, bookId uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Delete(ctx, bookId)
}
