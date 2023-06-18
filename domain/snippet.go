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
	FetchBySnippetID(c context.Context, userID string) ([]Snippet, error)
}

type SnippetUsecase interface {
	Create(c context.Context, snippet *Snippet) error
	FetchBySnippetID(c context.Context, userID string) ([]Snippet, error)
}
