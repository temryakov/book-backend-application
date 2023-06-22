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

const (
	CollectionSnippet = "Snippets"
)

type SnippetRepository interface {
	Fetch(c context.Context) ([]Snippet, error)
	FetchByID(c context.Context, snippetID uint16) (Snippet, error)
}

type SnippetUsecase interface {
	Fetch(c context.Context) ([]Snippet, error)
	FetchByID(c context.Context, snippetID uint16) (Snippet, error)
}
