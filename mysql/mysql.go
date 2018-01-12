package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import driver for mysql
	"github.com/jmoiron/sqlx"

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
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", opts.DBUser, opts.DBPassword, opts.DBHost, opts.DBPort, opts.DBName)
	mysqlDB, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return &DB{*mysqlDB}, nil
}

func (db *DB) SafePreparex(query string) *sqlx.Stmt {
	statement, err := db.Preparex(query)
	if err != nil {
		log.Errorf("Preparing statement failed. %s. Query: %s", err.Error(), query)
		return nil
	}

	return statement
}
