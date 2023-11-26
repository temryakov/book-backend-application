package domain

import (
	"context"

	"gorm.io/gorm"
)

type Review struct {
	BookId uint   `gorm:"integer";not null;default:null`
	UserId uint   `gorm:"integer";not null;default:null`
	Rating uint   `gorm:"integer";not null;default:null`
	Title  string `gorm:"varchar";not null;default:null`
	Text   string `gorm:"varchar";not null;default:null`
	gorm.Model
}

type ReviewRequest struct {
	BookId uint   `json:"book_id" binding:"required"`
	Rating uint   `json:"rating" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type ReviewResponse struct {
	BookAuthor   string `json:"book_author"`
	BookTitle    string `json:"book_title"`
	ReviewAuthor string `json:"review_author"`
	Rating       uint   `json:"rating"`
	ReviewTitle  string `json:"review_title"`
	ReviewText   string `json:"review_text"`
}

type ReviewQuery struct {
	BookId uint
	UserId uint
	ID     uint
}

type BookInfo struct {
	Title  string
	Author string
}

type UserInfo struct {
	Name string
}

type ReviewRepository interface {
	// FetchAllReview(c context.Context, BookID uint) (*[]Review, error)
	FetchReview(c context.Context, conditions *ReviewQuery) (*Review, error)
	CreateReview(c context.Context, review *Review) error
	// UpdateReview(c context.Context, review *Review, Model *Review) error
	DeleteReview(c context.Context, ReviewID uint) error
}

type ReviewUsecase interface {
	// FetchAllReview(c context.Context, BookID uint) (*[]Review, error)
	FetchReview(c context.Context, conditions *ReviewQuery) (*ReviewResponse, error)
	CreateReview(c context.Context, review *Review) error
	// UpdateReview(c context.Context, review *Review, Model *Review) error
	DeleteReview(c context.Context, ReviewID uint) error
}
