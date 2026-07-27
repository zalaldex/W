package main

import (
	"log"
	"os"
	"sync"

	tele "gopkg.in/telebot.v3"
)

// store holds each user's selected mode in memory, keyed by Telegram user
// ID. It is safe for concurrent use.
type store struct {
	mu    sync.RWMutex
	modes map[int64]Mode
}

func newStore() *store {
	return &store{modes: make(map[int64]Mode)}
}

func (s *store) get(id int64) Mode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if m, ok := s.modes[id]; ok {
		return m
	}
	return DefaultMode
}

func (s *store) set(id int64, m Mode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modes[id] = m
}

const welcomeText = "Send me any text, photo, video, or voice message, and I'll convert " +
	"it (and any caption) to Telegram monospace.\n\n" +
	"Use Settings to choose how it's chunked: Word, Sentence, Paragraph, or Full."

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	webhookURL := "https://" + os.Getenv("RAILWAY_PUBLIC_DOMAIN") + "/" + token
	

	poller := &tele.Webhook{
		Listen:   ":" + port,
		Endpoint: &tele.WebhookEndpoint{PublicURL: webhookURL},
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})
	if err != nil {
		log.Fatalf("failed to start bot: %v", err)
	}

	st := newStore()

	// Main persistent keyboard.
	mainMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnStart := mainMenu.Text("▶️ Start")
	btnSettings := mainMenu.Text("⚙️ Settings")
	mainMenu.Reply(mainMenu.Row(btnStart, btnSettings))

	// Settings keyboard: one mode per button, plus Back.
	settingsMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnWord := settingsMenu.Text(ModeLabel(ModeWord))
	btnSentence := settingsMenu.Text(ModeLabel(ModeSentence))
	btnParagraph := settingsMenu.Text(ModeLabel(ModeParagraph))
	btnFull := settingsMenu.Text(ModeLabel(ModeFull))
	btnBack := settingsMenu.Text("⬅️ Back")
	settingsMenu.Reply(
		settingsMenu.Row(btnWord, btnSentence),
		settingsMenu.Row(btnParagraph, btnFull),
		settingsMenu.Row(btnBack),
	)

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&btnStart, func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&btnSettings, func(c tele.Context) error {
		current := st.get(c.Sender().ID)
		return c.Send("Current mode: "+ModeLabel(current)+"\n\nChoose a mode:", settingsMenu)
	})

	bot.Handle(&btnBack, func(c tele.Context) error {
		return c.Send("Back to main menu.", mainMenu)
	})

	modeHandler := func(mode Mode) func(tele.Context) error {
		return func(c tele.Context) error {
			st.set(c.Sender().ID, mode)
			return c.Send("Mode set to "+ModeLabel(mode)+".", mainMenu)
		}
	}
	bot.Handle(&btnWord, modeHandler(ModeWord))
	bot.Handle(&btnSentence, modeHandler(ModeSentence))
	bot.Handle(&btnParagraph, modeHandler(ModeParagraph))
	bot.Handle(&btnFull, modeHandler(ModeFull))

	bot.Handle(tele.OnText, func(c tele.Context) error {
		mode := st.get(c.Sender().ID)
		return sendRendered(c, mode, c.Text(), mainMenu)
	})

	bot.Handle(tele.OnMedia, func(c tele.Context) error {
		mode := st.get(c.Sender().ID)
		return handleMedia(c, mode, mainMenu)
	})

	log.Println("bot started")
	bot.Start()
}

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
