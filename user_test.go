package myapp_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/mock"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestNewUserService(t *testing.T) {
	ur := &mock.UserRepository{}
	tracer := mocktracer.New()
	us := myapp.NewUserService(tracer, ur)
	testingutil.Assert(t, us != nil, "NewUserService returns nil")
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		ur            myapp.UserRepository
		tracer        opentracing.Tracer
		expectedUsers []myapp.User
		expectedError error
	}{
		{
			ur:     &mock.UserRepository{},
			tracer: mocktracer.New(),
			expectedUsers: []myapp.User{
				myapp.User{
					ID:    1,
					Name:  "test1",
					Email: "test1@example.com",
				},
				myapp.User{
					ID:    2,
					Name:  "test2",
					Email: "test2@example.com",
				},
			},
			expectedError: nil,
		},
		{
			ur:            &mock.UserRepositoryWithError{},
			tracer:        mocktracer.New(),
			expectedUsers: nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, tc := range testCases {
		us := myapp.NewUserService(tc.tracer, tc.ur)
		users, err := us.FindAll(context.Background())
		testingutil.Equals(t, tc.expectedUsers, users)
		testingutil.Equals(t, tc.expectedError, err)
	}
}

func TestFindByID(t *testing.T) {
	testCases := []struct {
		ur            myapp.UserRepository
		tracer        opentracing.Tracer
		id            int64
		expectedUser  myapp.User
		expectedError error
	}{
		{
			ur:     &mock.UserRepository{},
			tracer: mocktracer.New(),
			id:     1,
			expectedUser: myapp.User{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			},
			expectedError: nil,
		},
		{
			ur:            &mock.UserRepositoryWithError{},
			tracer:        mocktracer.New(),
			id:            1,
			expectedUser:  myapp.User{},
			expectedError: errors.New("internal error"),
		},
	}

	for _, tc := range testCases {
		us := myapp.NewUserService(tc.tracer, tc.ur)
		user, err := us.FindByID(context.Background(), tc.id)
		testingutil.Equals(t, tc.expectedUser, user)
		testingutil.Equals(t, tc.expectedError, err)
	}
}
