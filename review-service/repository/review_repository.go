package repository

import (
	"context"

	"review-service/domain"

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

func (rr *reviewRepository) FetchReview(ctx context.Context, conditions *domain.ReviewQuery) (*domain.Review, error) {

	var review domain.Review

	if err := rr.database.WithContext(ctx).Where(conditions).First(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (rr *reviewRepository) CreateReview(ctx context.Context, review *domain.Review) error {
	return rr.database.WithContext(ctx).Save(&review).Error
}

func (rr *reviewRepository) DeleteReview(ctx context.Context, reviewId uint) error {

	var review domain.Review

	if err := rr.database.WithContext(ctx).Delete(&review, reviewId).Error; err != nil {
		return err
	}
	return nil
}

func (rr *reviewRepository) DeleteReviewByBookId(ctx context.Context, BookId uint) error {

	var review domain.Review

	if err := rr.database.WithContext(ctx).Delete(&review, "book_id LIKE ?", BookId).Error; err != nil {
		return err
	}
	return nil
}
