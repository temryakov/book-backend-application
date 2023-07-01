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

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
	}
}

func (su *bookUsecase) FetchByID(c context.Context, bookId uint) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.FetchByID(ctx, bookId)
}

func (su *bookUsecase) Fetch(c context.Context) (*[]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Fetch(ctx)
}

func (su *bookUsecase) Create(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Create(ctx, book)
}
func (su *bookUsecase) Update(c context.Context, book *domain.Book, bookId uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	model, err := su.bookRepository.FetchByID(ctx, bookId)
	if err != nil {
		return err
	}
	return su.bookRepository.Update(ctx, book, model)
}

func (su *bookUsecase) Delete(c context.Context, bookId uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.Delete(ctx, bookId)
}
