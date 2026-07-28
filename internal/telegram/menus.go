package telegram

import (
	"telegram-monospace-bot/internal/formatter"

	tele "gopkg.in/telebot.v3"
)

// ==========================================
// Main Menu (Persistent Reply Keyboard)
// ==========================================

type mainMenuButtons struct {
	Start    tele.Btn
	Settings tele.Btn
}

func newMainMenu() (*tele.ReplyMarkup, mainMenuButtons) {
	menu := &tele.ReplyMarkup{
		ResizeKeyboard: true,
		IsPersistent:   true, // Forces Telegram to always show this keyboard
	}
	btns := mainMenuButtons{
		Start:    menu.Text("▶️ Start"),
		Settings: menu.Text("⚙️ Settings"),
	}
	menu.Reply(menu.Row(btns.Start, btns.Settings))
	return menu, btns
}

// ==========================================
// Settings Menu (Dynamic Inline Keyboard)
// ==========================================

var (
	btnModeWord      = tele.Btn{Unique: "mode_word"}
	btnModeSentence  = tele.Btn{Unique: "mode_sentence"}
	btnModeParagraph = tele.Btn{Unique: "mode_paragraph"}
	btnModeFull      = tele.Btn{Unique: "mode_full"}
	btnStats         = tele.Btn{Unique: "stats_view"}
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

// ==========================================
// Stats Menu (Inline Keyboard)
// ==========================================

var statsRefreshBtn = tele.Btn{
	Unique: "stats_refresh",
	Text:   "🔄 Refresh",
}

func statsMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	menu.Inline(
		menu.Row(statsRefreshBtn),
	)
	return menu
}
