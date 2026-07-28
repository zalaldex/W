package main

import tele "gopkg.in/telebot.v3"

// modeHandler returns a handler that sets the sender's mode to mode and
// confirms the change, returning to the main menu.
func modeHandler(st *store, mainMenu *tele.ReplyMarkup, mode Mode) func(tele.Context) error {
	return func(c tele.Context) error {
		st.set(c.Sender().ID, mode)
		return c.Send("Mode set to "+ModeLabel(mode)+".", mainMenu)
	}
}
