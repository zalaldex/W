package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token  string
	Port   string
	Domain string
}

func LoadConfig() (Config, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return Config{}, fmt.Errorf("BOT_TOKEN environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Token:  token,
		Port:   port,
		Domain: os.Getenv("PUBLIC_DOMAIN"),     //# os.Getenv("RAILWAY_PUBLIC_DOMAIN"),
	}, nil
}
