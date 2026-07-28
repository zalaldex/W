package main

import tele "gopkg.in/telebot.v3"

// handleMedia re-sends the received media with its caption rendered in the
// selected mode. If the rendered caption exceeds Telegram's caption limit,
// the media is sent with the first chunk as caption and any remaining
// chunks are sent as follow-up text messages so no content is lost.
func handleMedia(c tele.Context, mode Mode, menu *tele.ReplyMarkup) error {
	msg := c.Message()

	var chunks []string
	if msg.Caption != "" {
		rendered := Render(mode, msg.Caption)
		chunks = SplitForTelegram(rendered, telegramCaptionLimit)
	}

	firstCaption := ""
	if len(chunks) > 0 {
		firstCaption = chunks[0]
	}

	mediaOpts := &tele.SendOptions{ParseMode: tele.ModeMarkdown}
	if len(chunks) <= 1 {
		mediaOpts.ReplyMarkup = menu
	}

	if err := resendMedia(c, msg, firstCaption, mediaOpts); err != nil {
		return err
	}

	if len(chunks) <= 1 {
		return nil
	}
	rest := chunks[1:]
	for i, chunk := range rest {
		opts := &tele.SendOptions{ParseMode: tele.ModeMarkdown}
		if i == len(rest)-1 {
			opts.ReplyMarkup = menu
		}
		if err := c.Send(chunk, opts); err != nil {
			return err
		}
	}
	return nil
}
