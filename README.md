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

## Project structure

Every function lives in its own file, named after what it does, so
individual pieces are easy to find, read, and modify:

```
main.go                          entry point — wires everything together
config.go                        loads BOT_TOKEN / PORT / RAILWAY_PUBLIC_DOMAIN
new_bot.go                       constructs the telebot instance (webhook)
webhook_url.go                   builds the public webhook URL
store.go                         per-user mode storage (in memory)
welcome_text.go                  the /start welcome message

mode.go                          Mode type + constants
mode_label.go                    Mode -> button label
mode_parse.go                    button label -> Mode
mode_handler.go                  handler factory for mode-select buttons
limits.go                        Telegram message/caption length limits

main_menu.go                     builds the Start/Settings keyboard
settings_menu.go                 builds the Word/Sentence/Paragraph/Full keyboard
register_main_menu_handlers.go   wires /start, Start, Settings, Back
register_settings_handlers.go    wires the four mode buttons
register_content_handlers.go     wires incoming text + media

send_rendered.go                 sends rendered text, chunked to fit
handle_media.go                  re-sends media with rendered caption
resend_media.go                  re-sends a message's media by type

render.go                        Render: text -> monospaced text, by mode
render_units.go                  wraps each split unit in its own code span
wrap_code.go                     wraps a string in a Telegram code span
split_surrounding_space.go       separates leading/trailing whitespace

split_words.go                   splits text into words
split_sentences.go                splits text into sentences
split_paragraphs.go              splits text into paragraphs
closing_mark.go                  recognizes trailing quote/bracket marks

split_for_telegram.go            splits long output into multiple messages
best_split_point.go              picks the best boundary to cut at
last_word_break.go               finds the last whitespace break
last_sentence_break.go           finds the last sentence break
last_index_after.go              finds the index just after a substring
```

Adding a feature (a new mode, a new command, a new media type) usually
means adding one new file rather than editing a large existing one.

## Requirements

- Go 1.22+
- A Telegram bot token from [@BotFather](https://t.me/BotFather)

## Configuration

| Variable                 | Description                                                              |
| ------------------------- | -------------------------------------------------------------------------- |
| `BOT_TOKEN`               | Required. Your Telegram bot token from BotFather.                        |
| `PORT`                    | Optional. Port the webhook server listens on. Defaults to `8080`.        |
| `RAILWAY_PUBLIC_DOMAIN`   | Required. Public domain Telegram will send webhook updates to (no scheme). |

The bot runs in **webhook** mode: it starts an HTTP server on `PORT` and
registers `https://$RAILWAY_PUBLIC_DOMAIN/$BOT_TOKEN` with Telegram as the
webhook endpoint. It needs a publicly reachable domain (as provided
automatically on Railway) — it does not use long polling.

## Running locally

```bash
export BOT_TOKEN=your-telegram-bot-token
export RAILWAY_PUBLIC_DOMAIN=your-public-domain   # e.g. via a tunnel like ngrok
go run .
```

## Deploying on Railway

1. Push this repository to GitHub (or GitLab/Bitbucket).
2. Create a new project on Railway and connect your repository. Railway
   will detect the `Dockerfile` and build from it automatically.
3. Add an environment variable `BOT_TOKEN` with your bot token.
   `RAILWAY_PUBLIC_DOMAIN` and `PORT` are provided by Railway automatically.
4. Deploy. The bot will register its webhook and start receiving updates.

## Building with Docker

```bash
docker build -t monospace-bot .
docker run -e BOT_TOKEN=your-telegram-bot-token \
  -e RAILWAY_PUBLIC_DOMAIN=your-public-domain \
  -p 8080:8080 monospace-bot
```

## Notes

- Modes are stored in memory per user and reset if the process restarts.
- Stickers and video notes are re-sent as-is (they carry no caption in the
  Telegram Bot API, so there is nothing to convert).
- Unsupported message types are politely reported back to the user.
