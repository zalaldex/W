package bot

import tele "gopkg.in/telebot.v3"

// resendMedia sends the same media file back to the chat with the given
// caption, preserving the original media type.
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
