package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@accounts-db:5432/transactions?sslmode=disable"
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("sql.Open failed: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("db.Ping failed: %w", err)
	}

	return nil
}
