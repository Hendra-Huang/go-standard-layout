package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import driver for postgres

	"github.com/Hendra-Huang/go-standard-layout/log"
)

type (
	// DB of database
	DB struct {
		sqlx.DB
	}

	// Options of database
	Options struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
	}
)

// New database connection
func New(opts Options) (*DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", opts.DBHost, opts.DBPort, opts.DBUser, opts.DBPassword, opts.DBName)
	postgresDB, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &DB{*postgresDB}, nil
}

func (db *DB) SafePreparex(query string) *sqlx.Stmt {
	statement, err := db.Preparex(query)
	if err != nil {
		log.Errorf("Preparing statement failed. %s. Query: %s", err.Error(), query)
		return nil
	}

	return statement
}
