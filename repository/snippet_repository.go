package repository

import (
	"context"
	"errors"
	"snippetapp/domain"

	"gorm.io/gorm"
)

type snippetRepository struct {
	database   *gorm.DB
	collection string
}

func NewSnippetRepository(database *gorm.DB, collection string) domain.SnippetRepository {
	return &snippetRepository{
		database:   database,
		collection: collection,
	}
}

func (r *snippetRepository) FetchByID(ctx context.Context, id uint) (domain.Snippet, error) {

	var snippet domain.Snippet

	if err := r.database.WithContext(ctx).First(&snippet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Snippet{}, err
		}
		return domain.Snippet{}, err
	}
	return snippet, nil
}

func (r *snippetRepository) Fetch(ctx context.Context) ([]domain.Snippet, error) {

	var snippets []domain.Snippet

	if err := r.database.WithContext(ctx).Find(&snippets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domain.Snippet{}, err
		}
		return []domain.Snippet{}, err
	}
	return snippets, nil
}

func (r *snippetRepository) Create(ctx context.Context, snippet *domain.Snippet) error {

	return r.database.WithContext(ctx).Create(&snippet).Error
}

func (r *snippetRepository) Delete(ctx context.Context, id uint) (domain.Snippet, error) {

	var snippet domain.Snippet

	if err := r.database.WithContext(ctx).First(&snippet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Snippet{}, err
		}
		return domain.Snippet{}, err
	}
	return snippet, nil
}
