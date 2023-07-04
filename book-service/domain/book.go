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
	FetchBooks(c context.Context) (*[]Book, error)
	FetchBookByID(c context.Context, BookID uint) (*Book, error)
	CreateBook(c context.Context, book *Book) error
	UpdateBook(c context.Context, book *Book, Model *Book) error
	DeleteBook(c context.Context, BookID uint) error
}

type BookUsecase interface {
	FetchBooks(c context.Context) (*[]Book, error)
	FetchBookByID(c context.Context, BookID uint) (*Book, error)
	CreateBook(c context.Context, book *Book) error
	UpdateBook(c context.Context, book *Book, BookID uint) error
	DeleteBook(c context.Context, BookID uint) error
}
