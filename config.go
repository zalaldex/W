package main

import (
	"fmt"
	"os"
)

// config holds the settings needed to start the bot.
type config struct {
	Token  string
	Port   string
	Domain string
}

// loadConfig reads configuration from environment variables. BOT_TOKEN is
// required; PORT defaults to 8080 if unset.
func loadConfig() (config, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return config{}, fmt.Errorf("BOT_TOKEN environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return config{
		Token:  token,
		Port:   port,
		Domain: os.Getenv("RAILWAY_PUBLIC_DOMAIN"),
	}, nil
}
