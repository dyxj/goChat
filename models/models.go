package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func NewDatabase() (*sql.DB, error) {
	dbuser := os.Getenv("TEST_PGSQL_USR")
	dbpass := os.Getenv("TEST_PGSQL_PASS")
	dbname := "gochat"
	dbssl := "disable"
	fmt.Println(dbuser)
	fmt.Println(dbpass)
	// Connect to database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbuser, dbpass, dbname, dbssl)

	pgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Ping database
	err = pgDB.Ping()
	if err != nil {
		return nil, err
	}
	return pgDB, nil
}
