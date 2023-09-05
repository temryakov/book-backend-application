package repository

import (
	"context"

	"github.com/review-service/domain"

	"gorm.io/gorm"
)

type reviewRepository struct {
	database *gorm.DB
}

func NewReviewRepository(database *gorm.DB) domain.ReviewRepository {
	return &reviewRepository{
		database: database,
	}
}

func (r *reviewRepository) FetchReview(ctx context.Context, reviewId uint) (*domain.Review, error) {

	var review *domain.Review

	if err := r.database.WithContext(ctx).First(&review, reviewId).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, reviewId uint) error {

	var review *domain.Review

	if err := r.database.WithContext(ctx).Delete(&review, reviewId).Error; err != nil {
		return err
	}
	return nil
}
