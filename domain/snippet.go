package domain

import "context"

type Snippet struct {
	ID          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Author      string `json: "author"`
}

type SnippetUsecase interface {
	GetSnippetById(c context.Context, userID string) (*Snippet, error)
}
