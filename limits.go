package main

// telegramMessageLimit is Telegram's hard cap on text message length, in
// UTF-16 code units. We use it conservatively as a rune count, which is
// always <= the UTF-16 length and therefore always safe.
const telegramMessageLimit = 4096

// telegramCaptionLimit is Telegram's hard cap on media caption length.
const telegramCaptionLimit = 1024
