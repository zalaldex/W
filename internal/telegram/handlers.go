package telegram

import (
	"log"
	"time"

	"telegram-monospace-bot/internal/analytics"
	"telegram-monospace-bot/internal/formatter"
	"telegram-monospace-bot/internal/state"

	tele "gopkg.in/telebot.v3"
)

func RegisterHandlers(bot *tele.Bot, st *state.Store) {
	mainMenu, mainBtns := newMainMenu()

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&mainBtns.Start, func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&mainBtns.Settings, func(c tele.Context) error {
		current := st.Get(c.Sender().ID)
		menu := newSettingsMenu(current)
		return c.Send("Current mode: "+formatter.ModeLabel(current)+"\n\nChoose a mode:", menu)
	})

	bot.Handle(&btnModeWord, modeHandler(st, formatter.ModeWord))
	bot.Handle(&btnModeSentence, modeHandler(st, formatter.ModeSentence))
	bot.Handle(&btnModeParagraph, modeHandler(st, formatter.ModeParagraph))
	bot.Handle(&btnModeFull, modeHandler(st, formatter.ModeFull))

	bot.Handle(&btnStats, func(c tele.Context) error {
		s, err := analytics.LoadStats()
		if err != nil {
			log.Printf("stats: load: %v", err)
			c.Respond(&tele.CallbackResponse{Text: "Error loading stats."})
			return c.Send("Couldn't load statistics right now — please try again.")
		}
		c.Respond(&tele.CallbackResponse{})
		return c.Send(analytics.FormatStats(s, time.Now()), statsMenu(), tele.ModeHTML)
	})

	bot.Handle(&statsRefreshBtn, func(c tele.Context) error {
		s, err := analytics.LoadStats()
		if err != nil {
			log.Printf("stats: refresh: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}
		err = c.Edit(analytics.FormatStats(s, time.Now()), statsMenu(), tele.ModeHTML)
		if err != nil {
			if isNotModifiedErr(err) {
				return c.Respond(&tele.CallbackResponse{Text: "Already up to date."})
			}
			log.Printf("stats: edit: %v", err)
			return c.Respond(&tele.CallbackResponse{Text: "Couldn't refresh — try again."})
		}
		return c.Respond(&tele.CallbackResponse{Text: "Refreshed ✅"})
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		mode := st.Get(c.Sender().ID)
		return sendRendered(c, mode, c.Text(), mainMenu)
	})

	bot.Handle(tele.OnMedia, func(c tele.Context) error {
		mode := st.Get(c.Sender().ID)
		return handleMedia(c, mode, mainMenu)
	})
}

func modeHandler(st *state.Store, mode formatter.Mode) func(tele.Context) error {
	return func(c tele.Context) error {
		st.Set(c.Sender().ID, mode)
		menu := newSettingsMenu(mode)
		err := c.Edit("Current mode: "+formatter.ModeLabel(mode)+"\n\nChoose a mode:", menu)
		if err != nil && isNotModifiedErr(err) {
			return c.Respond(&tele.CallbackResponse{Text: "Already selected!"})
		}
		return c.Respond(&tele.CallbackResponse{Text: "Mode set to " + formatter.ModeLabel(mode)})
	}
}