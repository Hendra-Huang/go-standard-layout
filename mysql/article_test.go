// +build integration

package mysql_test

import (
	"context"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/mysql"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestFindByUserID(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "article")

	tracer := mocktracer.New()
	ar := mysql.NewArticleRepository(tracer, db, db)
	ctx := context.Background()

	testCases := []struct {
		userID         int64
		expectedError  error
		expectedLength int
	}{
		{
			userID:         1,
			expectedError:  nil,
			expectedLength: 2,
		},
		{
			userID:         2,
			expectedError:  nil,
			expectedLength: 0,
		},
	}
	for _, tc := range testCases {
		articles, err := ar.FindByUserID(ctx, tc.userID)
		testingutil.Equals(t, tc.expectedError, err)
		testingutil.Equals(t, tc.expectedLength, len(articles))
	}
}

func TestCreateArticle(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "article")

	tracer := mocktracer.New()
	ar := mysql.NewArticleRepository(tracer, db, db)
	ctx := context.Background()

	testCases := []struct {
		userID        int64
		title         string
		expectedError error
	}{
		{
			userID:        1,
			title:         "Title Test1",
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		err := ar.Create(ctx, tc.userID, tc.title)
		testingutil.Equals(t, tc.expectedError, err)
	}
}
