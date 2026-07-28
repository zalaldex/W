package store

import tele "gopkg.in/telebot.v3"

// Middleware returns a telebot middleware that records every incoming update
// as a single message for the sender and then calls the next handler.
func Middleware(db *sql.DB) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if sender := c.Sender(); sender != nil {
				TrackMessage(db, sender.ID)
			}
			return next(c)
		}
	}
}
