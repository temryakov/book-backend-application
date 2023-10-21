package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	review, err := ru.reviewRepository.FetchReview(ctx, conditions)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	url := "http://localhost:8080/api/book/" + strconv.FormatUint(uint64(review.BookId), 10)
	req, _ := http.NewRequest("GET", url, nil)

	log.Print(url)

	req.Header.Set("Accept", "application/x-protobuf")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	log.Print(body)

	if err != nil {
		return nil, err
	}

	return review, nil
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
