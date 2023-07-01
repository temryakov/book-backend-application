package domain

import (
	"context"

	"gorm.io/gorm"
)

type Book struct {
	Title          string `json:"title" binding:"required"`
	Author         string `json:"author" binding:"required"`
	ChaptersAmount uint   `json:"chapters_amount" binding:"required"`
	gorm.Model
}

const (
	CollectionBook = "Books"
)

type BookRepository interface {
	Fetch(c context.Context) (*[]Book, error)
	FetchByID(c context.Context, BookID uint) (*Book, error)
	Create(c context.Context, book *Book) error
	Update(c context.Context, book *Book, Model *Book) error
	Delete(c context.Context, BookID uint) error
}

type BookUsecase interface {
	Fetch(c context.Context) (*[]Book, error)
	FetchByID(c context.Context, BookID uint) (*Book, error)
	Create(c context.Context, book *Book) error
	Update(c context.Context, book *Book, BookID uint) error
	Delete(c context.Context, BookID uint) error
}
