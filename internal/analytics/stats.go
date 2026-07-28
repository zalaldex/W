package analytics

type Stats struct {
	ActiveUsers      int64
	UniqueUsers      int64
	MessagesToday    int64
	Messages24h      int64
	Messages7d       int64
	Messages30d      int64
	Messages1y       int64
	MessagesLifetime int64
}