package formatter

type Mode string

const (
	ModeWord      Mode = "word"
	ModeSentence  Mode = "sentence"
	ModeParagraph Mode = "paragraph"
	ModeFull      Mode = "full"
)

const DefaultMode = ModeFull