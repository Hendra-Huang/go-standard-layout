// +build integration

package mysql_test

import (
	"context"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/mysql"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestFindAll(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	users, err := ur.FindAll(context.Background())
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(users))
	testingutil.Equals(t, int64(1), users[0].ID)
	testingutil.Equals(t, "Myuser", users[0].Name)
	testingutil.Equals(t, "myuser@example.com", users[0].Email)
}

func TestFindByID(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	ctx := context.Background()

	testCases := []struct {
		userID        int64
		expectedName  string
		expectedEmail string
		expectedError error
	}{
		{
			userID:        1,
			expectedName:  "Myuser",
			expectedEmail: "myuser@example.com",
			expectedError: nil,
		},
		{
			userID:        1000,
			expectedName:  "",
			expectedEmail: "",
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		user, err := ur.FindByID(ctx, tc.userID)
		testingutil.Equals(t, tc.expectedError, err)
		testingutil.Equals(t, tc.expectedName, user.Name)
		testingutil.Equals(t, tc.expectedEmail, user.Email)
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	ctx := context.Background()

	testCases := []struct {
		email         string
		name          string
		expectedError error
	}{
		{
			email:         "test@example.com",
			name:          "test",
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		err := ur.Create(ctx, tc.email, tc.name)
		testingutil.Equals(t, tc.expectedError, err)
	}
}
