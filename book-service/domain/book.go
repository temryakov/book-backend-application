package domain

import (
	"context"

	"gorm.io/gorm"
)

type Book struct {
	Title          string `json:"title" binding:"required" gorm:"varchar;not null;default:null"`
	Author         string `json:"author" binding:"required" gorm:"varchar;not null;default:null"`
	ChaptersAmount uint   `json:"chapters_amount" binding:"required" gorm:"integer;not null;default:null"`
	gorm.Model
}

type BookData struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	ChaptersAmount uint   `json:"chapters_amount"`
}

type BookResponse struct {
	Message string   `json:"message"`
	Data    BookData `json:"data"`
}

type AllBookResponse struct {
	Message string     `json:"message"`
	Data    []BookData `json:"data"`
}

type BookRepository interface {
	FetchBooks(c context.Context) (*[]Book, error)
	FetchBookByID(c context.Context, BookID int) (*Book, error)
	CreateBook(c context.Context, book *Book) error
	UpdateBook(c context.Context, book *Book, Model *Book) error
	DeleteBook(c context.Context, BookID int) error
}

type BookUsecase interface {
	FetchBooks(c context.Context) (*[]Book, error)
	FetchBookByID(c context.Context, BookID int) (*Book, error)
	CreateBook(c context.Context, book *Book) error
	UpdateBook(c context.Context, book *Book, BookID int) error
	DeleteBook(c context.Context, BookID int) error
}

type BookProducer interface {
	DeleteBook(bookId int)
}
