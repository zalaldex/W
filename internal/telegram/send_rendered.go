package telegram

import (
	tele "gopkg.in/telebot.v3"
	"telegram-monospace-bot/internal/formatter"
)

func sendRendered(c tele.Context, mode formatter.Mode, text string, menu *tele.ReplyMarkup) error {
	rendered := formatter.Render(mode, text)
	chunks := splitForTelegram(rendered, telegramMessageLimit)
	for i, chunk := range chunks {
		opts := &tele.SendOptions{ParseMode: tele.ModeMarkdown}
		if i == len(chunks)-1 {
			opts.ReplyMarkup = menu
		}
		if err := c.Send(chunk, opts); err != nil {
			return err
		}
	}
	return nil
}