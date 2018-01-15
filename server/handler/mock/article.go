package mock

import (
	"context"
	"errors"

	"github.com/Hendra-Huang/go-standard-layout"
)

type ArticleService struct{}

func (as *ArticleService) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	return []myapp.Article{
		myapp.Article{1, userID, "test"},
		myapp.Article{2, userID, "test2"},
	}, nil
}

type ArticleServiceWithError struct{}

func (as *ArticleServiceWithError) FindByUserID(ctx context.Context, userID int64) ([]myapp.Article, error) {
	return nil, errors.New("internal error")
}
