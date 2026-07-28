package main

import tele "gopkg.in/telebot.v3"

// registerContentHandlers wires up handling of incoming text messages and
// media (photos, videos, voice notes, etc.).
func registerContentHandlers(bot *tele.Bot, st *store, mainMenu *tele.ReplyMarkup) {
	bot.Handle(tele.OnText, func(c tele.Context) error {
		mode := st.get(c.Sender().ID)
		return sendRendered(c, mode, c.Text(), mainMenu)
	})

	bot.Handle(tele.OnMedia, func(c tele.Context) error {
		mode := st.get(c.Sender().ID)
		return handleMedia(c, mode, mainMenu)
	})
}
