package stats

import (
	"context"
	"database/sql"
	"time"
)

// Manager provides operations to load statistics from a database.
type Manager struct {
	db  *sql.DB
	now func() time.Time
}

// NewManager constructs a Manager.
func NewManager(db *sql.DB) *Manager {
	return &Manager{db: db, now: time.Now}
}

// Load queries SQLite for a fresh Stats snapshot as of now.
func (m *Manager) Load(ctx context.Context) (Stats, error) {
	var s Stats
	now := m.now().UTC()

	row := m.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE last_seen >= ?`, now.Add(-5*time.Minute))
	if err := row.Scan(&s.ActiveUsers); err != nil {
		return s, err
	}

	row = m.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users`)
	if err := row.Scan(&s.UniqueUsers); err != nil {
		return s, err
	}

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if err := m.scanMessageCount(&s.MessagesToday, midnight); err != nil {
		return s, err
	}
	if err := m.scanMessageCount(&s.Messages24h, now.Add(-24*time.Hour)); err != nil {
		return s, err
	}
	if err := m.scanMessageCount(&s.Messages7d, now.AddDate(0, 0, -7)); err != nil {
		return s, err
	}
	if err := m.scanMessageCount(&s.Messages30d, now.AddDate(0, 0, -30)); err != nil {
		return s, err
	}
	if err := m.scanMessageCount(&s.Messages1y, now.AddDate(-1, 0, 0)); err != nil {
		return s, err
	}

	row = m.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM messages`)
	if err := row.Scan(&s.MessagesLifetime); err != nil {
		return s, err
	}

	return s, nil
}

func (m *Manager) scanMessageCount(dest *int64, since time.Time) error {
	row := m.db.QueryRow(`SELECT COUNT(*) FROM messages WHERE created_at >= ?`, since)
	return row.Scan(dest)
}
