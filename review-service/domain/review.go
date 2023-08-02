package domain

import (
	"context"

	"gorm.io/gorm"
)

type Review struct {
	BookId uint
	UserId uint
	Rating uint
	Title  string
	Text   string
	gorm.Model
}

type ReviewResponse struct {
	BookAuthor   string `json:"book_author"`
	BookTitle    string `json:"book_title"`
	ReviewAuthor string `json:"review_author"`
	Rating       uint   `json:"rating"`
	ReviewTitle  string `json:"review_title"`
	ReviewText   string `json:"review_text"`
}

type ReviewRepository interface {
	FetchAllReview(c context.Context, BookID uint) (*[]Review, error)
	FetchReview(c context.Context, ReviewID uint) (*Review, error)
	CreateReview(c context.Context, review *Review, BookID uint) error
	UpdateReview(c context.Context, review *Review, Model *Review) error
	DeleteReview(c context.Context, ReviewID uint) error
}

type ReviewUsecase interface {
	FetchAllReview(c context.Context, BookID uint) (*[]Review, error)
	FetchReview(c context.Context, ReviewID uint) (*Review, error)
	CreateReview(c context.Context, review *Review, BookID uint) error
	UpdateReview(c context.Context, review *Review, Model *Review) error
	DeleteReview(c context.Context, ReviewID uint) error
}
