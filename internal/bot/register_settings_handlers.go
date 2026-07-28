package bot

import (
	"github.com/zalaldex/W/internal/mode"
	tele "gopkg.in/telebot.v3"
)

// registerSettingsHandlers wires up the Word/Sentence/Paragraph/Full
// buttons on the settings keyboard.
func registerSettingsHandlers(bot *tele.Bot, st *store.Store, settingsBtns settingsMenuButtons, mainMenu *tele.ReplyMarkup) {
	bot.Handle(&settingsBtns.Word, modeHandler(st, mainMenu, mode.ModeWord))
	bot.Handle(&settingsBtns.Sentence, modeHandler(st, mainMenu, mode.ModeSentence))
	bot.Handle(&settingsBtns.Paragraph, modeHandler(st, mainMenu, mode.ModeParagraph))
	bot.Handle(&settingsBtns.Full, modeHandler(st, mainMenu, mode.ModeFull))
}
