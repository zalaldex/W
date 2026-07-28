package bot

import (
	"fmt"
	"os"

	tele "gopkg.in/telebot.v3"
)

// NewBotFromEnv constructs a telebot instance using environment variables
// BOT_TOKEN, PORT, and RAILWAY_PUBLIC_DOMAIN.
func NewBotFromEnv() (*tele.Bot, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("BOT_TOKEN is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	domain := os.Getenv("RAILWAY_PUBLIC_DOMAIN")
	if domain == "" {
		return nil, fmt.Errorf("RAILWAY_PUBLIC_DOMAIN is not set")
	}

	p := &tele.Webhook{
		Listen:   ":" + port,
		Endpoint: &tele.WebhookEndpoint{PublicURL: webhookURL(domain, token)},
	}

	return tele.NewBot(tele.Settings{Token: token, Poller: p})
}
