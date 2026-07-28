package main

import "time"

// activeWindow is how recently a user must have been seen to count as
// "currently using the bot".
const activeWindow = 5 * time.Minute

// loadStats queries SQLite for a fresh Stats snapshot as of now.
func loadStats() (Stats, error) {
	var s Stats
	now := time.Now().UTC()

	row := DB.QueryRow(`SELECT COUNT(*) FROM users WHERE last_seen >= ?`, now.Add(-activeWindow))
	if err := row.Scan(&s.ActiveUsers); err != nil {
		return s, err
	}

	row = DB.QueryRow(`SELECT COUNT(*) FROM users`)
	if err := row.Scan(&s.UniqueUsers); err != nil {
		return s, err
	}

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if err := scanMessageCount(&s.MessagesToday, midnight); err != nil {
		return s, err
	}
	if err := scanMessageCount(&s.Messages24h, now.Add(-24*time.Hour)); err != nil {
		return s, err
	}
	if err := scanMessageCount(&s.Messages7d, now.AddDate(0, 0, -7)); err != nil {
		return s, err
	}
	if err := scanMessageCount(&s.Messages30d, now.AddDate(0, 0, -30)); err != nil {
		return s, err
	}
	if err := scanMessageCount(&s.Messages1y, now.AddDate(-1, 0, 0)); err != nil {
		return s, err
	}

	row = DB.QueryRow(`SELECT COUNT(*) FROM messages`)
	if err := row.Scan(&s.MessagesLifetime); err != nil {
		return s, err
	}

	return s, nil
}

// scanMessageCount counts messages created at or after since, into dest.
func scanMessageCount(dest *int64, since time.Time) error {
	row := DB.QueryRow(`SELECT COUNT(*) FROM messages WHERE created_at >= ?`, since)
	return row.Scan(dest)
}
