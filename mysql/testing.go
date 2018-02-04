package mysql

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Hendra-Huang/go-standard-layout/testingutil"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

const (
	dbHost     = "127.0.0.1"
	dbPort     = "3307"
	dbUser     = "root"
	dbPassword = "root"
	dbName     = "myapp"
)

// CreateTestDatabase will create a test-database and test-schema
func CreateTestDatabase(t *testing.T) (*DB, string, func()) {
	db, dbErr := New(Options{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	})
	if dbErr != nil {
		t.Fatalf("Fail to connect database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	testDBName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE DATABASE " + testDBName)
	if err != nil {
		t.Fatalf("Fail to create database %s. %s", testDBName, err.Error())
	}

	testDB, dbErr := New(Options{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     testDBName,
	})
	if dbErr != nil {
		t.Fatalf("Fail to connect database. %s", dbErr.Error())
	}

	return testDB, testDBName, func() {
		_, err := db.Exec("DROP DATABASE " + testDBName)
		if err != nil {
			t.Fatalf("Fail to drop database %s. %s", testDBName, err.Error())
		}
	}
}

func LoadFixtures(t *testing.T, db *DB, fixtureName string) {
	loadSchema(t, db)
	content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s.sql", fixtureName))
	testingutil.Ok(t, err)

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			testingutil.Ok(t, err)
		}
	}
}

func loadSchema(t *testing.T, db *DB) {
	content, err := ioutil.ReadFile("./schema.sql")
	testingutil.Ok(t, err)

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			testingutil.Ok(t, err)
		}
	}
}
