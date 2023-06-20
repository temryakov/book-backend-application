package domain

import (
	"context"

	"gorm.io/gorm"
)

type Snippet struct {
	Title string
	Text  string
	gorm.Model
}

type SnippetRepository interface {
	Create(c context.Context, snippet *Snippet) error
	Fetch(c context.Context, userID string) ([]Snippet, error)
	FetchByID(c context.Context, userID string) ([]Snippet, error)
	Update(c context.Context, snippet *Snippet) error
	Delete(c context.Context, snippet *Snippet) error
}

type SnippetUsecase interface {
	Create(c context.Context, snippet *Snippet) error
	Fetch(c context.Context, userID string) ([]Snippet, error)
	FetchByID(c context.Context, userID string) ([]Snippet, error)
	Update(c context.Context, snippet *Snippet) error
	Delete(c context.Context, snippet *Snippet) error
}
