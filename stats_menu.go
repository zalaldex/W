package main

import tele "gopkg.in/telebot.v3"

// statsRefreshBtn is the inline button that reloads the Statistics message
// in place. Its Unique string is what register_stats_handlers.go binds to.
var statsRefreshBtn = tele.InlineButton{
	Unique: "stats_refresh",
	Text:   "🔄 Refresh",
}

// statsMenu builds the inline keyboard shown under the Statistics message.
func statsMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	menu.InlineKeyboard = [][]tele.InlineButton{
		{statsRefreshBtn},
	}
	return menu
}
