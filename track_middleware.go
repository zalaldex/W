package main

import tele "gopkg.in/telebot.v3"

// trackMiddleware records every incoming update (text, caption, media,
// sticker, document, callback, etc.) as exactly one message for the
// sending user, then passes control to the next handler. Registering this
// once via bot.Use(trackMiddleware) covers every handler in the bot
// without needing to edit each one individually.
func trackMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if sender := c.Sender(); sender != nil {
			trackMessage(sender.ID)
		}
		return next(c)
	}
}
