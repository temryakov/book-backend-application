package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-book-app/domain"
)

type bookProgressUsecase struct {
	bookProgressRepository domain.BookProgressRepository
	contextTimeout         time.Duration
}

func NewBookProgressUsecase(bookProgressRepository domain.BookProgressRepository, timeout time.Duration) domain.BookProgressRepository {
	return &bookProgressUsecase{
		bookProgressRepository: bookProgressRepository,
		contextTimeout:         timeout,
	}
}

func (su *bookProgressUsecase) FetchByID(c context.Context, bookProgressId uint) (domain.BookProgress, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.bookProgressRepository.FetchByID(ctx, bookProgressId)
}
