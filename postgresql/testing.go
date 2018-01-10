package postgresql

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"database/sql"

	_ "github.com/lib/pq" // postgresql driver
)

const (
	dbPort     = 5439
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "test"
)

// CreateTestDatabase will create a test-database and test-schema
func CreateTestDatabase(t *testing.T) (*sql.DB, string, func()) {
	connectionString := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable", dbPort, dbUser, dbPassword, dbName)
	db, dbErr := sql.Open("postgres", connectionString)
	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	schemaName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE SCHEMA " + schemaName)
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	return db, schemaName, func() {
		_, err := db.Exec("DROP SCHEMA " + schemaName + " CASCADE")
		if err != nil {
			t.Fatalf("Fail to drop database. %s", err.Error())
		}
	}
}
