package domain

import (
	"context"

	"gorm.io/gorm"
)

type Snippet struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	gorm.Model
}

const (
	CollectionSnippet = "Snippets"
)

type SnippetRepository interface {
	Fetch(c context.Context) ([]Snippet, error)
	FetchByID(c context.Context, snippetID uint) (Snippet, error)
	Create(c context.Context, snippet *Snippet) error
	Update(c context.Context, snippet *Snippet, snippetID uint) error
	Delete(c context.Context, snippetID uint) error
}

type SnippetUsecase interface {
	Fetch(c context.Context) ([]Snippet, error)
	FetchByID(c context.Context, snippetID uint) (Snippet, error)
	Create(c context.Context, snippet *Snippet) error
	Update(c context.Context, snippet *Snippet, snippetID uint) error
	Delete(c context.Context, snippetID uint) error
}
