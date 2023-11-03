package usecase

import (
	"context"
	"time"

	"review-service/bootstrap"
	"review-service/domain"
	"review-service/transport"

	"gorm.io/gorm"
)

type reviewUsecase struct {
	reviewRepository domain.ReviewRepository
	config           bootstrap.Config
	contextTimeout   time.Duration
}

func NewReviewUsecase(reviewRepository domain.ReviewRepository, cfg *bootstrap.Config, timeout time.Duration) domain.ReviewUsecase {
	return &reviewUsecase{
		reviewRepository: reviewRepository,
		config:           *cfg,
		contextTimeout:   timeout,
	}
}

func (ru *reviewUsecase) FetchReview(c context.Context, conditions *domain.ReviewQuery) (*domain.ReviewResponse, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()

	bookCh, userCh := make(chan transport.BookInfo), make(chan transport.UserInfo)

	review, err := ru.reviewRepository.FetchReview(ctx, conditions)
	if err != nil {
		return nil, err
	}

	go transport.FetchBookInfo(ctx, ru.config, review.BookId, bookCh)
	go transport.FetchUserInfo(ctx, ru.config, review.UserId, userCh)

	bookInfo := <-bookCh

	if bookInfo.Error != nil {
		return nil, *bookInfo.Error
	}

	userInfo := <-userCh

	if userInfo.Error != nil {
		return nil, *userInfo.Error
	}

	return &domain.ReviewResponse{
		BookAuthor:   *bookInfo.Author,
		BookTitle:    *bookInfo.Title,
		ReviewAuthor: *userInfo.Name,
		Rating:       review.Rating,
		ReviewTitle:  review.Title,
		ReviewText:   review.Text,
	}, nil
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
		return domain.ErrReviewIsExist
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return ru.reviewRepository.CreateReview(ctx, review)
}

func (ru *reviewUsecase) DeleteReview(c context.Context, reviewId uint) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.reviewRepository.DeleteReview(ctx, reviewId)
}
