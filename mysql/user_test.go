// +build integration

package mysql_test

import (
	"context"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/mysql"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestFindAll(t *testing.T) {
	t.Parallel()
	db, dbName, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, dbName, "user")

	ur := mysql.NewUserRepository(db, db)
	users, err := ur.FindAll(context.Background())
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(users))
	testingutil.Equals(t, int64(1), users[0].ID)
	testingutil.Equals(t, "Myuser", users[0].Name)
	testingutil.Equals(t, "myuser@example.com", users[0].Email)
}

func TestFindByID(t *testing.T) {
	t.Parallel()
	db, dbName, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, dbName, "user")

	expectedID := int64(1)
	ur := mysql.NewUserRepository(db, db)
	user, err := ur.FindByID(context.Background(), expectedID)
	testingutil.Ok(t, err)
	testingutil.Equals(t, expectedID, user.ID)
	testingutil.Equals(t, "Myuser", user.Name)
	testingutil.Equals(t, "myuser@example.com", user.Email)
}

func TestFindByIDWithNotFound(t *testing.T) {
	t.Parallel()
	db, dbName, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, dbName, "user")

	expectedID := int64(1000)
	ur := mysql.NewUserRepository(db, db)
	user, err := ur.FindByID(context.Background(), expectedID)
	testingutil.Ok(t, err)
	testingutil.Equals(t, int64(0), user.ID)
	testingutil.Equals(t, "", user.Name)
	testingutil.Equals(t, "", user.Email)
}

func TestCreate(t *testing.T) {
	t.Parallel()
	db, dbName, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, dbName, "user")

	ur := mysql.NewUserRepository(db, db)
	err := ur.Create(context.Background(), 10, "test@example.com", "test")
	testingutil.Ok(t, err)
}

func TestCreateWithDuplicateID(t *testing.T) {
	t.Parallel()
	db, dbName, cleanup := mysql.CreateTestDatabase(t)
	defer cleanup()

	mysql.LoadFixtures(t, db, dbName, "user")

	ur := mysql.NewUserRepository(db, db)
	err := ur.Create(context.Background(), 1, "test@example.com", "test")
	testingutil.Assert(t, err != nil, "Error should not be nil")
}
