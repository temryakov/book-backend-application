package repository

import (
	"context"
	"errors"
	"snippetapp/domain"

	"gorm.io/gorm"
)

type snippetRepository struct {
	db *gorm.DB
}

func (r *snippetRepository) GetSnippetByID(ctx context.Context, id int) (*domain.Snippet, error) {

	var snippet domain.Snippet

	if err := r.db.WithContext(ctx).First(&snippet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("snippet not found")
		}
		return nil, err
	}
	return &snippet, nil
}
