package stats

import "time"

// Stats holds a snapshot of bot usage counts, as shown on the Statistics page.
type Stats struct {
	ActiveUsers      int64 // distinct users seen in the last 5 minutes
	UniqueUsers      int64 // distinct users, all-time
	MessagesToday    int64 // messages since local midnight UTC
	Messages24h      int64 // messages in the trailing 24 hours
	Messages7d       int64 // messages in the trailing 7 days
	Messages30d      int64 // messages in the trailing 30 days
	Messages1y       int64 // messages in the trailing 1 year
	MessagesLifetime int64 // all messages ever recorded
}

// UpdatedAt returns the given time formatted for display (UTC).
func (s Stats) UpdatedAt(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}
