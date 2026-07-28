package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/zalaldex/W/internal/bot"
	"github.com/zalaldex/W/internal/db"
	"github.com/zalaldex/W/internal/stats"
	"github.com/zalaldex/W/internal/store"

	tele "gopkg.in/telebot.v3"
)

func main() {
	// load config from environment
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN not set")
	}
	domain := os.Getenv("RAILWAY_PUBLIC_DOMAIN")
	if domain == "" {
		log.Fatal("RAILWAY_PUBLIC_DOMAIN not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// optional stats DB
	var dbh *sql.DB
	if os.Getenv("ENABLE_STATS") == "true" {
		p := os.Getenv("DB_PATH")
		if p == "" {
			p = db.DefaultDBPath
		}
		var err error
		dbh, err = db.Open(p)
		if err != nil {
			log.Fatalf("open db: %v", err)
		}
		defer dbh.Close()
	}

	botInstance, err := bot.NewBotFromEnv()
	if err != nil {
		log.Fatalf("new bot: %v", err)
	}

	// register stats handlers (Manager may be nil)
	var statsMgr *stats.Manager
	if dbh != nil {
		statsMgr = stats.NewManager(dbh)
	}
	bot.RegisterStatsHandlers(botInstance, statsMgr)

	// register track middleware if db is available
	if dbh != nil {
		botInstance.Use(store.Middleware(dbh))
	}

	// TODO: register other handlers (rendering, media handling, menus)

	log.Printf("starting bot")
	botInstance.Start()
}
