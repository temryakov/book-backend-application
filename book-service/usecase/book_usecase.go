package usecase

import (
	"context"
	"time"

	"book-service/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
	producer       *kafka.Producer
}

func NewBookUsecase(bookRepository domain.BookRepository, producer *kafka.Producer, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
		producer:       producer,
	}
}

func (bu *bookUsecase) FetchBookByID(c context.Context, bookId int) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.FetchBookByID(ctx, bookId)
}

func (bu *bookUsecase) FetchBooks(c context.Context) (*[]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.FetchBooks(ctx)
}

func (bu *bookUsecase) CreateBook(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.CreateBook(ctx, book)
}
func (bu *bookUsecase) UpdateBook(c context.Context, book *domain.Book, bookId int) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	model, err := bu.bookRepository.FetchBookByID(ctx, bookId)
	if err != nil {
		return err
	}
	return bu.bookRepository.UpdateBook(ctx, book, model)
}

func (su *bookUsecase) DeleteBook(c context.Context, bookId int) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookRepository.DeleteBook(ctx, bookId)
}
