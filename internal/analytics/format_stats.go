package analytics

import (
	"fmt"
	"time"
)

func FormatStats(s Stats, updatedAt time.Time) string {
	return fmt.Sprintf(
		"<b>📊 Statistics</b>\n\n"+
			"🟢 Active users: <b>%d</b>\n"+
			"👥 Unique users (all-time): <b>%d</b>\n\n"+
			"💬 Messages today: <b>%d</b>\n"+
			"💬 Last 24 hours: <b>%d</b>\n"+
			"💬 Last 7 days: <b>%d</b>\n"+
			"💬 Last 30 days: <b>%d</b>\n"+
			"💬 Last 1 year: <b>%d</b>\n\n"+
			"📈 Lifetime total: <b>%d</b>\n\n"+
			"<i>Updated %s UTC</i>",
		s.ActiveUsers,
		s.UniqueUsers,
		s.MessagesToday,
		s.Messages24h,
		s.Messages7d,
		s.Messages30d,
		s.Messages1y,
		s.MessagesLifetime,
		updatedAt.UTC().Format("2006-01-02 15:04:05"),
	)
}