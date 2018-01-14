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
)

func TestNewUserHandler(t *testing.T) {
	us := &mock.UserService{}
	userHandler := handler.NewUserHandler(us)
	testingutil.Assert(t, userHandler != nil, "UserHandler is nil")
}

func TestGetAllUsers(t *testing.T) {
	testCases := []struct {
		us                 handler.UserServicer
		expectedStatusCode int
		expectedUsers      []myapp.User
	}{
		{
			us:                 &mock.UserService{},
			expectedStatusCode: http.StatusOK,
			expectedUsers: []myapp.User{
				myapp.User{1, "test@example.com", "test"},
				myapp.User{2, "test2@example.com", "test2"},
			},
		},
		{
			us:                 &mock.UserServiceWithError{},
			expectedStatusCode: http.StatusInternalServerError,
			expectedUsers:      nil,
		},
	}

	for _, tc := range testCases {
		userHandler := handler.NewUserHandler(tc.us)
		req, err := http.NewRequest("GET", "/", nil)
		testingutil.Ok(t, err)
		rr := httptest.NewRecorder()
		httphandler := http.HandlerFunc(userHandler.GetAllUsers)
		httphandler.ServeHTTP(rr, req)

		testingutil.Equals(t, tc.expectedStatusCode, rr.Code)

		var expectedResp string
		if len(tc.expectedUsers) > 0 {
			expectedRespByte, err := json.Marshal(tc.expectedUsers)
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			expectedResp = "\n"
		}

		testingutil.Equals(t, expectedResp, rr.Body.String())
	}
}

func TestGetUserByID(t *testing.T) {
	testCases := []struct {
		us                 handler.UserServicer
		requestPath        string
		expectedStatusCode int
		expectedUser       myapp.User
	}{
		{
			us:                 &mock.UserService{},
			requestPath:        "/api/v1/user/1",
			expectedStatusCode: http.StatusOK,
			expectedUser:       myapp.User{1, "test@example.com", "test"},
		},
		{
			us:                 &mock.UserService{},
			requestPath:        "/api/v1/user/0",
			expectedStatusCode: http.StatusNotFound,
			expectedUser:       myapp.User{},
		},
		{
			us:                 &mock.UserServiceWithError{},
			requestPath:        "/api/v1/user/1",
			expectedStatusCode: http.StatusInternalServerError,
			expectedUser:       myapp.User{},
		},
	}

	for _, tc := range testCases {
		userHandler := handler.NewUserHandler(tc.us)

		r := router.New(router.Options{
			Timeout: 1 * time.Second,
		})
		r.Get("/api/v1/user/{id}", userHandler.GetUserByID)
		ts := httptest.NewServer(r)
		defer ts.Close()

		res, err := http.Get(ts.URL + tc.requestPath)
		testingutil.Ok(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		testingutil.Ok(t, err)

		testingutil.Equals(t, tc.expectedStatusCode, res.StatusCode)

		var expectedResp string
		if tc.expectedUser.ID > 0 {
			expectedRespByte, err := json.Marshal(tc.expectedUser)
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			expectedResp = "\n"
		}
		testingutil.Equals(t, expectedResp, string(body))
	}
}
