package telegram

import (
	tele "gopkg.in/telebot.v3"
)

func NewBot(token, port string, _ string /* ignored domain arg */) (*tele.Bot, error) {
	// Hardcoded domain straight into the repository
	publicURL := "https://w-production-ee10.up.railway.app"

	// Configure webhook poller
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
