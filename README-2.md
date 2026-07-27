# Monospace Telegram Bot

A Telegram bot that converts any text you send — plus captions on photos,
videos, voice notes, and other media — into Telegram monospace formatting.

You choose how the text is chunked before each chunk is wrapped in its own
monospace span:

- **Word** — each word is individually monospaced.
- **Sentence** — each sentence is individually monospaced.
- **Paragraph** — each paragraph is individually monospaced.
- **Full** — the entire message is monospaced as one block.

Original spacing, line breaks, and content are preserved. If the result
exceeds Telegram's message length limit, it's automatically split into
multiple messages, preferring to break at paragraph, then sentence, then
word, then character boundaries — so nothing is cut off mid-word or
mid-sentence unless there is no other option.

## Usage

Start a chat with the bot and use the persistent keyboard:

- **Start** — shows a welcome message.
- **Settings** — choose your active mode (Word, Sentence, Paragraph, Full).

Once a mode is set, just send text or media with a caption, and the bot
replies with the converted version. Each user's mode is remembered
independently for the lifetime of the running process (in memory — see
[Notes](#notes)).

## Requirements

- Go 1.22+
- A Telegram bot token from [@BotFather](https://t.me/BotFather)

## Configuration

The bot requires exactly one environment variable:

| Variable    | Description                          |
| ----------- | ------------------------------------- |
| `BOT_TOKEN` | Your Telegram bot token from BotFather |

`PORT` is read automatically if present (Render sets this for Web
Services) and defaults to `8080` otherwise. It only controls the
health-check server described below — it has no effect on Telegram
functionality.

## Running locally

```bash
export BOT_TOKEN=your-telegram-bot-token
go run .
```

The bot uses long polling, so no public URL or webhook configuration is
needed.

## Deploying on Render (free tier)

Render's free tier only supports **Web Services**, which must bind to a
port and respond to HTTP health checks — a pure long-polling process
doesn't do that on its own. To satisfy this, the bot runs a minimal
`net/http` server alongside its Telegram polling loop. This server has no
Telegram functionality; it only exists so Render considers the service
healthy. Telegram updates are still received exclusively via long
polling.

1. Push this repository to GitHub (or GitLab/Bitbucket).
2. In the Render dashboard, create a new **Web Service**.
3. Connect your repository. Render will detect the `Dockerfile` and build
   from it automatically.
4. Add an environment variable `BOT_TOKEN` with your bot token. Render
   automatically provides `PORT`; the app reads it and falls back to
   `8080` if it's ever unset.
5. Deploy. The bot will start polling Telegram immediately.

**Free tier caveat:** Render's free Web Services spin down after 15
minutes without incoming HTTP traffic, and only spin back up on the next
HTTP request. Telegram long polling does not count as incoming HTTP
traffic to Render, so a free instance will eventually go idle and stop
receiving Telegram updates until something (e.g. a browser visit to the
service URL, or an external uptime pinger) wakes it back up. For a bot
that must always be responsive, an external scheduled pinger (hitting `/`
every ~10 minutes) or a paid Render plan avoids this.

## Building with Docker

```bash
docker build -t monospace-bot .
docker run -e BOT_TOKEN=your-telegram-bot-token -p 8080:8080 monospace-bot
```

## Notes

- Modes are stored in memory per user and reset if the process restarts.
- Stickers and video notes are re-sent as-is (they carry no caption in the
  Telegram Bot API, so there is nothing to convert).
- Unsupported message types are politely reported back to the user.
