package domain

import (
	"context"

	"gorm.io/gorm"
)

type BookProgress struct {
	CompletedChapters uint `json:"completed_chapters" binding:"required"`
	IsBookCompleted   bool `json:"is_book_completed"`
	BookInfo          Book
	gorm.Model
}

const (
	CollectionBookProgress = "BooksProgress"
)

type BookProgressRepository interface {
	Fetch(c context.Context) ([]Book, error)
	FetchByID(c context.Context, BookID uint) (Book, error)
	Save(c context.Context, book *Book) error
	Delete(c context.Context, BookID uint) error
}

type BookProgressUsecase interface {
	Fetch(c context.Context) ([]Book, error)
	FetchByID(c context.Context, BookID uint) (Book, error)
	Save(c context.Context, book *Book) error
	Delete(c context.Context, BookID uint) error
}
