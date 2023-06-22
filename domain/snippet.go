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
	FetchByID(c context.Context, snippetID uint16) (Snippet, error)
}

type SnippetUsecase interface {
	FetchByID(c context.Context, snippetID uint16) (Snippet, error)
}
