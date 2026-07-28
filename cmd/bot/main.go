package main

import (
	"log"

	"telegram-monospace-bot/internal/analytics"
	"telegram-monospace-bot/internal/config"
	"telegram-monospace-bot/internal/state"
	"telegram-monospace-bot/internal/telegram"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := analytics.InitDB(); err != nil {
		log.Fatalf("stats db: %v", err)
	}
	defer analytics.CloseDB()

	bot, err := telegram.NewBot(cfg.Token, cfg.Port, cfg.Domain)
	if err != nil {
		log.Fatalf("failed to start bot: %v", err)
	}
	bot.Use(analytics.TrackMiddleware)

	st := state.NewStore()

	telegram.RegisterHandlers(bot, st)

	log.Println("bot started")
	bot.Start()
}