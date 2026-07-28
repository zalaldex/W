package main

import tele "gopkg.in/telebot.v3"

// registerSettingsHandlers wires up the Word/Sentence/Paragraph/Full
// buttons on the settings keyboard.
func registerSettingsHandlers(bot *tele.Bot, st *store, settingsBtns settingsMenuButtons, mainMenu *tele.ReplyMarkup) {
	bot.Handle(&settingsBtns.Word, modeHandler(st, mainMenu, ModeWord))
	bot.Handle(&settingsBtns.Sentence, modeHandler(st, mainMenu, ModeSentence))
	bot.Handle(&settingsBtns.Paragraph, modeHandler(st, mainMenu, ModeParagraph))
	bot.Handle(&settingsBtns.Full, modeHandler(st, mainMenu, ModeFull))
}
