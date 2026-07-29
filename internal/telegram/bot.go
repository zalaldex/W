package telegram

import (
	"os"
	tele "gopkg.in/telebot.v3"
)

func NewBot(token, port, domain string) (*tele.Bot, error) {
	// Ensure domain starts with https://
	publicURL := domain
	if len(domain) < 8 || domain[:8] != "https://" {
		publicURL = "https://" + domain
	}

	// Configure webhook with explicit endpoint path matching your token
	poller := &tele.Webhook{
		Listen: ":" + port,
		Endpoint: &tele.WebhookEndpoint{
			PublicURL: publicURL,
			// Telebot uses the token to securely route incoming Telegram updates
			CustomURL: token, 
		},
	}

	return tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})
}
