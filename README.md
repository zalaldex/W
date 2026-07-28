# Monospace Telegram Bot (Restructured)
    
This repository has been refactored into a Domain-Driven ("Package by Feature") layout.

- **cmd/bot/**: Contains the main application entrypoint.
- **internal/config/**: Environment variable management.
- **internal/state/**: In-memory mode storage per user.
- **internal/analytics/**: SQLite persistence and statistics handlers.
- **internal/formatter/**: Core domain logic for parsing text and converting to monospace.
- **internal/telegram/**: Bot-specific routing, menus, API limits, and inline UI checkmarks.

To run locally:
```bash
export BOT_TOKEN=your-token
export RAILWAY_PUBLIC_DOMAIN=your-tunnel-domain
go run ./cmd/bot
```