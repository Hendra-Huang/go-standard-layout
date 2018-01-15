package responseutil_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/server/responseutil"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestJSON(t *testing.T) {
	testCases := []struct {
		options            []responseutil.WriterOption
		data               interface{}
		expectedHeaders    map[string]string
		expectedStatusCode int
	}{
		{
			options: nil,
			data:    nil,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			options: []responseutil.WriterOption{
				responseutil.WithHeader("Content-Type", "text/plain"),
				responseutil.WithStatus(http.StatusCreated),
			},
			data: map[string]interface{}{
				"id":   1,
				"name": "test1",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "text/plain",
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			options: []responseutil.WriterOption{
				responseutil.WithHeader("Set-Cookie", "test"),
				responseutil.WithStatus(http.StatusCreated),
			},
			data: map[string]interface{}{
				"id":   1,
				"name": "test1",
			},
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Set-Cookie":   "test",
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		rr := httptest.NewRecorder()
		responseutil.JSON(rr, tc.data, tc.options...)

		headers := rr.Header()
		for headerName, headerValue := range tc.expectedHeaders {
			testingutil.Equals(t, headerValue, headers.Get(headerName))
		}

		expectedResByte, err := json.Marshal(tc.data)
		testingutil.Ok(t, err)
		testingutil.Equals(t, string(expectedResByte), rr.Body.String())
		testingutil.Equals(t, tc.expectedStatusCode, rr.Code)
	}
}
