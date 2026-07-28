package main

import "log"

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := newBot(cfg.Token, cfg.Port, webhookURL(cfg.Domain, cfg.Token))
	if err != nil {
		log.Fatalf("failed to start bot: %v", err)
	}

	st := newStore()

	mainMenu, mainBtns := newMainMenu()
	settingsMenu, settingsBtns := newSettingsMenu()

	registerMainMenuHandlers(bot, st, mainBtns, mainMenu, settingsBtns, settingsMenu)
	registerSettingsHandlers(bot, st, settingsBtns, mainMenu)
	registerContentHandlers(bot, st, mainMenu)

	log.Println("bot started")
	bot.Start()
}
