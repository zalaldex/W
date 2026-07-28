package main

import tele "gopkg.in/telebot.v3"

// newBot creates a telebot instance configured to receive updates via
// webhook on port, using publicURL as its public endpoint.
func newBot(token, port, publicURL string) (*tele.Bot, error) {
	poller := &tele.Webhook{
		Listen:   ":" + port,
		Endpoint: &tele.WebhookEndpoint{PublicURL: publicURL},
	}
	return tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})
}
