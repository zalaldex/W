package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

// registerStatsHandlers wires the "📊 Statistics" settings button and the
// Refresh inline button on the resulting message.
//
// NOTE: update the text match below if your Statistics button label in
// settings_menu.go differs from "📊 Statistics".
func registerStatsHandlers(bot *tele.Bot) {
	bot.Handle("📊 Statistics", func(c tele.Context) error {
		s, err := loadStats()
		if err != nil {
			log.Printf("stats: load: %v", err)
			return c.Send("Couldn't load statistics right now — please try again.")
		}
		return c.Send(formatStats(s, time.Now()), statsMenu(), tele.ModeHTML)
	})

	bot.Handle(&statsRefreshBtn, func(c tele.Context) error {
		s, err := loadStats()
		if err != nil {
			log.Printf("stats: refresh: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}

		err = c.Edit(formatStats(s, time.Now()), statsMenu(), tele.ModeHTML)
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
