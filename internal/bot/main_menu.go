package bot

import tele "gopkg.in/telebot.v3"

// mainMenuButtons groups the buttons on the main persistent keyboard so
// handlers can be registered against the same button values used to build
// the menu.
type mainMenuButtons struct {
	Start    tele.Btn
	Settings tele.Btn
}

// newMainMenu builds the main persistent keyboard (Start, Settings) and
// returns it along with its buttons for handler registration.
func newMainMenu() (*tele.ReplyMarkup, mainMenuButtons) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btns := mainMenuButtons{
		Start:    menu.Text("▶️ Start"),
		Settings: menu.Text("⚙️ Settings"),
	}
	menu.Reply(menu.Row(btns.Start, btns.Settings))
	return menu, btns
}
