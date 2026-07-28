package bot

import tele "gopkg.in/telebot.v3"

// settingsMenuButtons groups the buttons on the settings keyboard so
// handlers can be registered against the same button values used to build
// the menu.
type settingsMenuButtons struct {
	Word      tele.Btn
	Sentence  tele.Btn
	Paragraph tele.Btn
	Full      tele.Btn
	Back      tele.Btn
}

// newSettingsMenu builds the settings keyboard (one button per Mode, plus
// Back) and returns it along with its buttons for handler registration.
func newSettingsMenu() (*tele.ReplyMarkup, settingsMenuButtons) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btns := settingsMenuButtons{
		Word:      menu.Text(ModeLabel(ModeWord)),
		Sentence:  menu.Text(ModeLabel(ModeSentence)),
		Paragraph: menu.Text(ModeLabel(ModeParagraph)),
		Full:      menu.Text(ModeLabel(ModeFull)),
		Back:      menu.Text("⬅️ Back"),
	}
	menu.Reply(
		menu.Row(btns.Word, btns.Sentence),
		menu.Row(btns.Paragraph, btns.Full),
		menu.Row(btns.Back),
	)
	return menu, btns
}
