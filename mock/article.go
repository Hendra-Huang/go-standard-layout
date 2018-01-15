package mock

import (
	"context"
	"errors"

	"github.com/Hendra-Huang/go-standard-layout"
)

type ArticleRepository struct{}

func (ar *ArticleRepository) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	articles := []myapp.Article{
		myapp.Article{
			ID:     1,
			UserID: userID,
			Title:  "test1",
		},
		myapp.Article{
			ID:     2,
			UserID: userID,
			Title:  "test2",
		},
	}

	return articles, nil
}

type ArticleRepositoryWithError struct{}

func (ar *ArticleRepositoryWithError) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	return nil, errors.New("internal error")
}
