package main

// Mode identifies how incoming text is chunked before being wrapped in
// Telegram monospace formatting.
type Mode string

const (
	ModeWord      Mode = "word"
	ModeSentence  Mode = "sentence"
	ModeParagraph Mode = "paragraph"
	ModeFull      Mode = "full"
)

// DefaultMode is used for users who have not chosen one yet.
const DefaultMode = ModeFull
