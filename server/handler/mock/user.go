package mock

import (
	"context"
	"errors"

	"github.com/Hendra-Huang/go-standard-layout"
)

type UserService struct{}

func (us *UserService) FindAll(ctx context.Context) ([]myapp.User, error) {
	return []myapp.User{
		myapp.User{1, "test@example.com", "test"},
		myapp.User{2, "test2@example.com", "test2"},
	}, nil
}

func (us *UserService) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	return myapp.User{1, "test@example.com", "test"}, nil
}

type UserServiceWithError struct{}

func (us *UserServiceWithError) FindAll(ctx context.Context) ([]myapp.User, error) {
	return nil, errors.New("internal error")
}

func (us *UserServiceWithError) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	return myapp.User{}, errors.New("internal error")
}
