package usecase

import (
	"context"
	"time"

	"github.com/review-service/domain"
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

func (ru *reviewUsecase) FetchReview(c context.Context, reviewId uint) (*domain.Review, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.reviewRepository.FetchReview(ctx, reviewId)
}
