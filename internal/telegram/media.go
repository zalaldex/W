package telegram

import (
	tele "gopkg.in/telebot.v3"
	"telegram-monospace-bot/internal/formatter"
)

func handleMedia(c tele.Context, mode formatter.Mode, menu *tele.ReplyMarkup) error {
	msg := c.Message()

	var chunks []string
	if msg.Caption != "" {
		rendered := formatter.Render(mode, msg.Caption)
		chunks = splitForTelegram(rendered, telegramCaptionLimit)
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

func resendMedia(c tele.Context, msg *tele.Message, caption string, opts *tele.SendOptions) error {
	switch {
	case msg.Photo != nil:
		m := *msg.Photo
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Video != nil:
		m := *msg.Video
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Animation != nil:
		m := *msg.Animation
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Audio != nil:
		m := *msg.Audio
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Voice != nil:
		m := *msg.Voice
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Document != nil:
		m := *msg.Document
		m.Caption = caption
		return c.Send(&m, opts)
	case msg.Sticker != nil:
		return c.Send(msg.Sticker, opts)
	case msg.VideoNote != nil:
		return c.Send(msg.VideoNote, opts)
	default:
		return c.Send("Unsupported media type.", opts.ReplyMarkup)
	}
}