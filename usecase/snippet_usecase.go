package usecase

import (
	"context"
	"time"

	"github.com/temryakov/go-backend-snippet-app/domain"
)

type snippetUsecase struct {
	snippetRepository domain.SnippetRepository
	contextTimeout    time.Duration
}

func NewSnippetUsecase(snippetRepository domain.SnippetRepository, timeout time.Duration) domain.SnippetUsecase {
	return &snippetUsecase{
		snippetRepository: snippetRepository,
		contextTimeout:    timeout,
	}
}

func (su *snippetUsecase) FetchByID(c context.Context, snippetID uint) (domain.Snippet, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.snippetRepository.FetchByID(ctx, snippetID)
}

func (su *snippetUsecase) Fetch(c context.Context) ([]domain.Snippet, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.snippetRepository.Fetch(ctx)
}

func (su *snippetUsecase) Create(c context.Context, snippet *domain.Snippet) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.snippetRepository.Create(ctx, snippet)
}

func (su *snippetUsecase) Delete(c context.Context, snippetID uint) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.snippetRepository.Delete(ctx, snippetID)
}
