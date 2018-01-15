package myapp_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/mock"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestNewArticleService(t *testing.T) {
	ar := &mock.ArticleRepository{}
	as := myapp.NewArticleService(ar)
	testingutil.Assert(t, as != nil, "NewArticleService returns nil")
}

func TestFindByUserID(t *testing.T) {
	testCases := []struct {
		ar               myapp.ArticleRepository
		userID           int64
		expectedArticles []myapp.Article
		expectedError    error
	}{
		{
			ar:     &mock.ArticleRepository{},
			userID: 1,
			expectedArticles: []myapp.Article{
				myapp.Article{
					ID:     1,
					UserID: 1,
					Title:  "test1",
				},
				myapp.Article{
					ID:     2,
					UserID: 1,
					Title:  "test2",
				},
			},
			expectedError: nil,
		},
		{
			ar:               &mock.ArticleRepositoryWithError{},
			userID:           1,
			expectedArticles: nil,
			expectedError:    errors.New("internal error"),
		},
	}

	for _, tc := range testCases {
		as := myapp.NewArticleService(tc.ar)
		articles, err := as.FindByUserID(context.Background(), tc.userID)
		testingutil.Equals(t, tc.expectedArticles, articles)
		testingutil.Equals(t, tc.expectedError, err)
	}
}
