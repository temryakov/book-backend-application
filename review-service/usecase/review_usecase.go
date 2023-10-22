package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"review-service/bootstrap"
	"review-service/domain"
	"review-service/transport"

	rp "review-service/proto"

	"google.golang.org/protobuf/proto"
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

	review, err := ru.reviewRepository.FetchReview(ctx, conditions)

	if err != nil {
		return nil, err
	}

	resp, err := transport.FetchBookInfo(ctx, ru.config, review.BookId)

	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	bookInfo := &rp.GetBookResponse{}

	err = proto.Unmarshal(body, bookInfo)
	if err != nil {
		fmt.Println("Error unmarshalling protobuf:", err)
		return nil, err
	}

	log.Print(bookInfo)

	if err != nil {
		return nil, err
	}

	return &domain.ReviewResponse{
		BookAuthor:   bookInfo.GetAuthor(),
		BookTitle:    bookInfo.GetTitle(),
		ReviewAuthor: "Test Test",
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
