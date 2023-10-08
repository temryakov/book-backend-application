package usecase

import (
	"context"
	"time"

	"github.com/review-service/domain"
	"gorm.io/gorm"
)

type reviewUsecase struct {
	reviewRepository domain.ReviewRepository
	contextTimeout   time.Duration
}

func NewReviewUsecase(reviewRepository domain.ReviewRepository, timeout time.Duration) domain.ReviewUsecase {
	return &reviewUsecase{
		reviewRepository: reviewRepository,
		contextTimeout:   timeout,
	}
}

func (ru *reviewUsecase) FetchReview(c context.Context, conditions *domain.ReviewQuery) (*domain.Review, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.reviewRepository.FetchReview(ctx, conditions)
}

func (ru *reviewUsecase) CreateReview(c context.Context, review *domain.Review) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)

	query := domain.ReviewQuery{
		BookId: review.BookId,
		UserId: review.UserId,
	}

	defer cancel()

	_, err := ru.reviewRepository.FetchReview(ctx, &query)

	if err == nil {
		cancel()
		return domain.ErrReviewIsExist
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		cancel()
		return err
	}
	return ru.reviewRepository.CreateReview(ctx, review)
}

func (ru *reviewUsecase) DeleteReview(c context.Context, reviewId uint) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.reviewRepository.DeleteReview(ctx, reviewId)
}
