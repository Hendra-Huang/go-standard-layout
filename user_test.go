package myapp_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout"
	"github.com/Hendra-Huang/go-standard-layout/mock"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestNewUserService(t *testing.T) {
	ur := &mock.UserRepository{}
	us := myapp.NewUserService(ur)
	testingutil.Assert(t, us != nil, "NewUserService returns nil")
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		ur            myapp.UserRepository
		expectedUsers []myapp.User
		expectedError error
	}{
		{
			ur: &mock.UserRepository{},
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
			expectedUsers: nil,
			expectedError: errors.New("internal error"),
		},
	}

	for _, tc := range testCases {
		us := myapp.NewUserService(tc.ur)
		users, err := us.FindAll(context.Background())
		testingutil.Equals(t, tc.expectedUsers, users)
		testingutil.Equals(t, tc.expectedError, err)
	}
}

func TestFindByID(t *testing.T) {
	testCases := []struct {
		ur            myapp.UserRepository
		id            int64
		expectedUser  myapp.User
		expectedError error
	}{
		{
			ur: &mock.UserRepository{},
			id: 1,
			expectedUser: myapp.User{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			},
			expectedError: nil,
		},
		{
			ur:            &mock.UserRepositoryWithError{},
			id:            1,
			expectedUser:  myapp.User{},
			expectedError: errors.New("internal error"),
		},
	}

	for _, tc := range testCases {
		us := myapp.NewUserService(tc.ur)
		user, err := us.FindByID(context.Background(), tc.id)
		testingutil.Equals(t, tc.expectedUser, user)
		testingutil.Equals(t, tc.expectedError, err)
	}
}
