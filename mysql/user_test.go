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

	expectedID := int64(1)
	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	user, err := ur.FindByID(context.Background(), expectedID)
	testingutil.Ok(t, err)
	testingutil.Equals(t, expectedID, user.ID)
	testingutil.Equals(t, "Myuser", user.Name)
	testingutil.Equals(t, "myuser@example.com", user.Email)
}

func TestFindByIDWithNotFound(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	expectedID := int64(1000)
	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	user, err := ur.FindByID(context.Background(), expectedID)
	testingutil.Ok(t, err)
	testingutil.Equals(t, int64(0), user.ID)
	testingutil.Equals(t, "", user.Name)
	testingutil.Equals(t, "", user.Email)
}

func TestCreate(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	err := ur.Create(context.Background(), 10, "test@example.com", "test")
	testingutil.Ok(t, err)
}

func TestCreateWithDuplicateID(t *testing.T) {
	t.Parallel()
	db, _, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, "user")

	tracer := mocktracer.New()
	ur := mysql.NewUserRepository(tracer, db, db)
	err := ur.Create(context.Background(), 1, "test@example.com", "test")
	testingutil.Assert(t, err != nil, "Error should not be nil")
}
