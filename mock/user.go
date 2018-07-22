package mock

import (
	"context"
	"errors"

	"github.com/Hendra-Huang/go-standard-layout"
)

type UserRepository struct{}

func (ur *UserRepository) FindAll(ctx context.Context) ([]myapp.User, error) {
	users := []myapp.User{
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
	}

	return users, nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	user := myapp.User{
		ID:    id,
		Name:  "test",
		Email: "test@example.com",
	}

	return user, nil
}

func (ur *UserRepository) Create(ctx context.Context, email, name string) error {
	return nil
}

type UserRepositoryWithError struct{}

func (ur *UserRepositoryWithError) FindAll(ctx context.Context) ([]myapp.User, error) {
	return nil, errors.New("internal error")
}

func (ur *UserRepositoryWithError) FindByID(ctx context.Context, id int64) (myapp.User, error) {
	user := myapp.User{}

	return user, errors.New("internal error")
}

func (ur *UserRepositoryWithError) Create(ctx context.Context, email, name string) error {
	return errors.New("internal error")
}
