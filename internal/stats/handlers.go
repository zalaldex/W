package stats

import (
	"log"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

var statsRefreshBtn = tele.InlineButton{Unique: "stats_refresh", Text: "Refresh"}

// RegisterHandlers wires the "📊 Statistics" settings button and the
// Refresh inline button on the resulting message.
func RegisterHandlers(bot *tele.Bot, m *Manager) {
	bot.Handle("📊 Statistics", func(c tele.Context) error {
		if m == nil {
			return c.Send("Statistics are unavailable.")
		}

		s, err := m.Load(c.Context())
		if err != nil {
			log.Printf("stats: load: %v", err)
			return c.Send("Couldn't load statistics right now — please try again.")
		}
		return c.Send(Format(s, time.Now()), statsMenu(), tele.ModeHTML)
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

		err = c.Edit(Format(s, time.Now()), statsMenu(), tele.ModeHTML)
		if err != nil {
			if isNotModifiedErr(err) {
				return c.Respond(&tele.CallbackResponse{Text: "Already up to date."})
			}
			log.Printf("stats: edit: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}

		return c.Respond(&tele.CallbackResponse{Text: "Refreshed ✅"})
	})
}

// isNotModifiedErr reports whether err is Telegram's "message is not
// modified" error, returned when editing a message with identical content.
func isNotModifiedErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "message is not modified")
}
