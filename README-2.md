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

## Running locally

```bash
export BOT_TOKEN=your-telegram-bot-token
go run .
```

The bot uses long polling, so no public URL or webhook configuration is
needed.

## Deploying on Render

1. Push this repository to GitHub (or GitLab/Bitbucket).
2. In the Render dashboard, create a new **Background Worker** (not a Web
   Service — this bot uses long polling and doesn't listen on a port).
3. Connect your repository. Render will detect the `Dockerfile` and build
   from it automatically.
4. Add an environment variable `BOT_TOKEN` with your bot token.
5. Deploy. The bot will start polling Telegram immediately.

## Building with Docker

```bash
docker build -t monospace-bot .
docker run -e BOT_TOKEN=your-telegram-bot-token monospace-bot
```

## Notes

- Modes are stored in memory per user and reset if the process restarts.
- Stickers and video notes are re-sent as-is (they carry no caption in the
  Telegram Bot API, so there is nothing to convert).
- Unsupported message types are politely reported back to the user.
