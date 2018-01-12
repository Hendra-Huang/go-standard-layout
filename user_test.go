package myapp_test

import (
	"context"
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
	ur := &mock.UserRepository{}
	us := myapp.NewUserService(ur)

	users, err := us.FindAll(context.Background())
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(users))
}

func TestFindByID(t *testing.T) {
	ur := &mock.UserRepository{}
	us := myapp.NewUserService(ur)

	expectedID := int64(1)
	user, err := us.FindByID(context.Background(), expectedID)
	testingutil.Ok(t, err)
	testingutil.Equals(t, expectedID, user.ID)
}
