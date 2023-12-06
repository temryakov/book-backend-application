package usecase

import (
	"context"
	"sync"
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

	res := &domain.ReviewResponse{
		Rating:      review.Rating,
		ReviewTitle: review.Title,
		ReviewText:  review.Text,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := transport.FetchBookInfo(ctx, ru.config, review.BookId, bookCh); err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		if err := transport.FetchUserInfo(ctx, ru.config, review.UserId, userCh); err != nil {
			return
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case bookResult := <-bookCh:
			res.BookAuthor = *bookResult.Author
			res.BookTitle = *bookResult.Title

		case userResult := <-userCh:
			res.ReviewAuthor = *userResult.Name

		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	wg.Wait()

	return res, nil
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

func (ru *reviewUsecase) DeleteReview(c context.Context, reviewId uint, userId uint) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)

	defer cancel()

	query := domain.ReviewQuery{
		ReviewID: reviewId,
		UserId:   userId,
	}

	_, err := ru.reviewRepository.FetchReview(ctx, &query)

	err := ru.reviewRepository.DeleteReview(ctx, reviewId)

	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}
