# Integrating the Statistics feature

These are the files to add/edit that I couldn't see directly. Everything
else (db.go, db_schema.go, track_message.go, stats.go, load_stats.go,
format_stats.go, stats_menu.go, register_stats_handlers.go,
is_not_modified_err.go, track_middleware.go) is ready to drop into your repo
root as-is.

## 1. go.mod — add the SQLite driver

```
go get modernc.org/sqlite
```

This is a pure-Go SQLite driver (no cgo), which keeps Railway builds simple.

## 2. main.go — initialize the DB and register everything

In your main() function, before you start polling/serving webhooks, add:

```go
func main() {
    // ... your existing config loading ...

    if err := initDB(); err != nil {
        log.Fatalf("stats db: %v", err)
    }
    defer DB.Close()

    bot := newBot(cfg) // however you currently construct it

    bot.Use(trackMiddleware) // record every incoming update as one message

    registerMainMenuHandlers(bot)
    registerSettingsHandlers(bot)
    registerContentHandlers(bot)
    registerStatsHandlers(bot) // <-- add this line

    // ... your existing webhook/start logic ...
}
```

Order matters only in that `bot.Use(trackMiddleware)` should be called
before `bot.Start()` (standard telebot middleware registration — it applies
to all handlers registered on this bot instance regardless of order
relative to the register* calls, but it's cleanest to put it right after
bot construction).

## 3. settings_menu.go — add the Statistics button

See settings_menu.go.patch.txt for the exact change. In short: add a
`📊 Statistics` button to your Settings reply keyboard. If you use a
different label, update the string match in register_stats_handlers.go
to match it exactly.

## 4. Railway — persist the database across deploys

Railway's container filesystem is ephemeral: a new deploy wipes it. To keep
stats across deploys/restarts:

1. In the Railway dashboard, add a **Volume** to your service.
2. Mount it at `/data`.
3. That's it — db.go already points at `/data/stats.db`.

Without a volume, stats reset on every deploy (but still survive normal
restarts/crashes within the same deploy, since Railway containers keep
their filesystem between restarts, just not between deploys).

If you'd rather not use a volume, change `dbPath` in db.go to a relative
path — it'll just reset more often.

## 5. Sanity check

After wiring this in:
- Send the bot a few messages of different types (text, sticker, photo).
- Open Settings → 📊 Statistics.
- Numbers should reflect what you just sent (Active users = 1, Messages
  today/24h/etc. all incrementing).
- Tap Refresh — the same message should update in place, no new message
  sent, and Telegram will *not* show a "not modified" error even if you
  tap twice quickly (handled via isNotModifiedErr).
