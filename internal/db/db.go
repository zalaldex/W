package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// DefaultDBPath is the default path where the SQLite file will be stored.
const DefaultDBPath = "/data/stats.db"

// Open opens (or creates) the SQLite database at path and ensures the schema exists.
// If path is empty, DefaultDBPath is used.
func Open(path string) (*sql.DB, error) {
	if path == "" {
		path = DefaultDBPath
	}

	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	db.SetMaxOpenConns(1) // SQLite + WAL: single writer keeps things simple and avoids lock errors

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	if err := createSchema(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
