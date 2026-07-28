package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
	"github.com/zalaldex/W/internal/stats"
)

var statsRefreshBtn = tele.InlineButton{Unique: "stats_refresh", Text: "Refresh"}

// RegisterStatsHandlers wires the "📊 Statistics" settings button and the
// Refresh inline button on the resulting message.
func RegisterStatsHandlers(bot *tele.Bot, m *stats.Manager) {
	bot.Handle("📊 Statistics", func(c tele.Context) error {
		if m == nil {
			return c.Send("Statistics are unavailable.")
		}

		s, err := m.Load(c.Context())
		if err != nil {
			log.Printf("stats: load: %v", err)
			return c.Send("Couldn't load statistics right now — please try again.")
		}
		return c.Send(stats.Format(s, time.Now()), statsMarkup(), tele.ModeHTML)
	})

	bot.Handle(&statsRefreshBtn, func(c tele.Context) error {
		if m == nil {
			return c.Respond(&tele.CallbackResponse{Text: "Statistics are unavailable."})
		}

		s, err := m.Load(c.Context())
		if err != nil {
			log.Printf("stats: refresh: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}

		err = c.Edit(stats.Format(s, time.Now()), statsMarkup(), tele.ModeHTML)
		if err != nil {
			// Telegram errors if the content is byte-for-byte identical to
			// the current message (e.g. two refreshes within the same
			// second). That's not a real failure, so just ack silently.
			if isNotModifiedErr(err) {
				return c.Respond(&tele.CallbackResponse{Text: "Already up to date."})
			}
			log.Printf("stats: edit: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}

		return c.Respond(&tele.CallbackResponse{Text: "Refreshed ✅"})
	})
}

// statsMarkup builds the inline markup used with Statistics messages.
func statsMarkup() *tele.ReplyMarkup {
	m := &tele.ReplyMarkup{}
	m.InlineKeyboard = [][]tele.InlineButton{{statsRefreshBtn}}
	return m
}
