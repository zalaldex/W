package telegram

import tele "gopkg.in/telebot.v3"

func NewBot(token, port, domain string) (*tele.Bot, error) {
	publicURL := "https://" + domain + "/" + token
	poller := &tele.Webhook{
		Listen:   ":" + port,
		Endpoint: &tele.WebhookEndpoint{PublicURL: publicURL},
	}
	return tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})
}