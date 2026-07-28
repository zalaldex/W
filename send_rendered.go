package main

import tele "gopkg.in/telebot.v3"

// sendRendered renders text under mode and sends it, splitting across
// multiple messages if needed to respect Telegram's length limit.
func sendRendered(c tele.Context, mode Mode, text string, menu *tele.ReplyMarkup) error {
	rendered := Render(mode, text)
	chunks := SplitForTelegram(rendered, telegramMessageLimit)
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
