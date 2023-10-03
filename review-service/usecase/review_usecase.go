package usecase

import (
	"context"
	"errors"
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

func (ru *reviewUsecase) CreateReview(c context.Context, review *domain.Review) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.reviewRepository.CreateReview(ctx, review)
}

func (ru *reviewUsecase) DeleteReview(c context.Context, reviewId uint, userId uint) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	review, err := ru.FetchReview(c, reviewId)
	if err != nil {
		return err
	}
	if review.UserId != userId {
		return errors.New("user is not have permission to delete")
	}
	return ru.reviewRepository.DeleteReview(ctx, reviewId)
}
