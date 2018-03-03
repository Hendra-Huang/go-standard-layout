package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/router"
	"github.com/Hendra-Huang/go-standard-layout/server/handler"
	"github.com/Hendra-Huang/go-standard-layout/server/handler/mock"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestNewArticleHandler(t *testing.T) {
	as := &mock.ArticleService{}
	articleHandler := handler.NewArticleHandler(as)
	testingutil.Assert(t, articleHandler != nil, "ArticleHandler is nil")
}

func TestGetAllArticlesByUserID(t *testing.T) {
	testCases := []struct {
		as                 handler.ArticleServicer
		requestPath        string
		expectedStatusCode int
		expectedArticles   []myapp.Article
	}{
		{
			as:                 &mock.ArticleService{},
			requestPath:        "/api/v1/user/1/articles",
			expectedStatusCode: http.StatusOK,
			expectedArticles: []myapp.Article{
				myapp.Article{1, 1, "test"},
				myapp.Article{2, 1, "test2"},
			},
		},
		{
			as:                 &mock.ArticleServiceWithError{},
			requestPath:        "/api/v1/user/1/articles",
			expectedStatusCode: http.StatusInternalServerError,
			expectedArticles:   nil,
		},
	}

	for _, tc := range testCases {
		articleHandler := handler.NewArticleHandler(tc.as)

		tracer := mocktracer.New()
		r := router.New(router.Options{
			Timeout: 1 * time.Second,
		}, tracer)
		r.Get("/api/v1/user/{user_id}/articles", articleHandler.GetAllArticlesByUserID)
		ts := httptest.NewServer(r)
		defer ts.Close()

		res, err := http.Get(ts.URL + tc.requestPath)
		testingutil.Ok(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		testingutil.Ok(t, err)

		testingutil.Equals(t, tc.expectedStatusCode, res.StatusCode)

		var expectedResp string
		if len(tc.expectedArticles) > 0 {
			expectedRespByte, err := json.Marshal(tc.expectedArticles)
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			expectedResp = "\n"
		}
		testingutil.Equals(t, expectedResp, string(body))
	}
}
