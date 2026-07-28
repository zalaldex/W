package analytics

import tele "gopkg.in/telebot.v3"

func TrackMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if sender := c.Sender(); sender != nil {
			trackMessage(sender.ID)
		}
		return next(c)
	}
}