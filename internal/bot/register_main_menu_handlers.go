package bot

import (
	"github.com/zalaldex/W/internal/mode"
	tele "gopkg.in/telebot.v3"
)

// registerMainMenuHandlers wires up /start plus the Start, Settings, and
// Back buttons.
func registerMainMenuHandlers(bot *tele.Bot, st *store.Store, mainBtns mainMenuButtons, mainMenu *tele.ReplyMarkup, settingsBtns settingsMenuButtons, settingsMenu *tele.ReplyMarkup) {
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&mainBtns.Start, func(c tele.Context) error {
		return c.Send(welcomeText, mainMenu)
	})

	bot.Handle(&mainBtns.Settings, func(c tele.Context) error {
		current := st.Get(c.Sender().ID)
		return c.Send("Current mode: "+mode.ModeLabel(current)+"\n\nChoose a mode:", settingsMenu)
	})

	bot.Handle(&settingsBtns.Back, func(c tele.Context) error {
		return c.Send("Back to main menu.", mainMenu)
	})
}
