package telegram

import (
	tele "gopkg.in/telebot.v3"
)

func NewBot(token, port, domain string) (*tele.Bot, error) {
	// Ensure domain starts with https://
	publicURL := domain
	if len(domain) < 8 || domain[:8] != "https://" {
		publicURL = "https://" + domain
	}

	// Configure webhook poller correctly for telebot v3
	poller := &tele.Webhook{
		Listen: ":" + port,
		Endpoint: &tele.WebhookEndpoint{
			PublicURL: publicURL + "/" + token,
		},
	}

	return tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})
}
