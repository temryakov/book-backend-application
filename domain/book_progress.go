package domain

import (
	"context"

	"gorm.io/gorm"
)

type BookProgress struct {
	CompletedChapters uint `json:"completed_chapters" binding:"required"`
	IsBookCompleted   bool `json:"is_book_completed"`
	gorm.Model
}

const (
	CollectionBookProgress = "BooksProgress"
)

type BookProgressRepository interface {
	FetchByID(c context.Context, BookProgressID uint) (BookProgress, error)
}

type BookProgressUsecase interface {
	FetchByID(c context.Context, BookProgressID uint) (BookProgress, error)
}
