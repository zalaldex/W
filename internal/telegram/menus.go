package telegram

import (
	"telegram-monospace-bot/internal/formatter"
	tele "gopkg.in/telebot.v3"
)

type mainMenuButtons struct {
	Start    tele.Btn
	Settings tele.Btn
}

func newMainMenu() (*tele.ReplyMarkup, mainMenuButtons) {
	menu := &tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	btns := mainMenuButtons{
		Start:    menu.Text("▶️ Start"),
		Settings: menu.Text("⚙️ Settings"),
	}
	menu.Reply(menu.Row(btns.Start, btns.Settings))
	return menu, btns
}

var (
	btnModeWord      = tele.InlineButton{Unique: "mode_word"}
	btnModeSentence  = tele.InlineButton{Unique: "mode_sentence"}
	btnModeParagraph = tele.InlineButton{Unique: "mode_paragraph"}
	btnModeFull      = tele.InlineButton{Unique: "mode_full"}
	btnStats         = tele.InlineButton{Unique: "stats_view"}
)

func newSettingsMenu(currentMode formatter.Mode) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	label := func(m formatter.Mode) string {
		if m == currentMode {
			return "✔️ " + formatter.ModeLabel(m)
		}
		return formatter.ModeLabel(m)
	}

	w, s, p, f, st := btnModeWord, btnModeSentence, btnModeParagraph, btnModeFull, btnStats
	w.Text = label(formatter.ModeWord)
	s.Text = label(formatter.ModeSentence)
	p.Text = label(formatter.ModeParagraph)
	f.Text = label(formatter.ModeFull)
	st.Text = "📊 Statistics"

	menu.Inline(
		menu.Row(w, s),
		menu.Row(p, f),
		menu.Row(st),
	)
	return menu
}

var statsRefreshBtn = tele.InlineButton{
	Unique: "stats_refresh",
	Text:   "🔄 Refresh",
}

func statsMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	menu.InlineKeyboard = [][]tele.InlineButton{
		{statsRefreshBtn},
	}
	return menu
}