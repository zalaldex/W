package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// DB is the shared SQLite handle used for statistics storage.
var DB *sql.DB

// dbPath is where the SQLite file lives. Railway's filesystem is ephemeral
// across deploys unless a volume is mounted at this path; mount a volume at
// /data in the Railway service settings to persist stats across deploys.
const dbPath = "/data/stats.db"

// initDB opens the SQLite database and ensures the schema exists.
func initDB() error {
	db, err := sql.Open("sqlite", dbPath+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		return fmt.Errorf("open sqlite: %w", err)
	}

	db.SetMaxOpenConns(1) // SQLite + WAL: single writer keeps things simple and avoids lock errors

	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping sqlite: %w", err)
	}

	DB = db
	return createSchema(db)
}
