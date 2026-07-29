package analytics

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const dbPath = "./stats.db"

func InitDB() error {
	d, err := sql.Open("sqlite", dbPath+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)")
	if err != nil {
		return fmt.Errorf("open sqlite: %w", err)
	}

	d.SetMaxOpenConns(1)

	if err := d.Ping(); err != nil {
		return fmt.Errorf("ping sqlite: %w", err)
	}

	db = d
	return createSchema(d)
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
