# Monospace Telegram Bot

A Telegram bot that converts text and media captions into Telegram monospace formatting.

## Build & Run (local)

```bash
export BOT_TOKEN=your-telegram-bot-token
export RAILWAY_PUBLIC_DOMAIN=your-public-domain
export PORT=8080
# Enable stats (optional)
export ENABLE_STATS=true
# Optionally set DB path (defaults to /data/stats.db)
export DB_PATH=/data/stats.db

# Build and run
go build -o bot ./cmd/monospace-bot
./bot
```

## Docker

The Docker image exposes port 8080. The bot stores its SQLite database at `/data/stats.db` by default; to persist statistics across restarts you should mount a volume at `/data`.

Example run with a host volume:

```bash
docker build -t monospace-bot .

docker run -e BOT_TOKEN=your-telegram-bot-token \
  -e RAILWAY_PUBLIC_DOMAIN=your-public-domain \
  -e ENABLE_STATS=true \
  -v /path/on/host/data:/data \
  -p 8080:8080 monospace-bot
```

If you don't set ENABLE_STATS=true the bot will run but statistics-related functionality will be disabled.
