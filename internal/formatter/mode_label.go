package formatter

func ModeLabel(m Mode) string {
	switch m {
	case ModeWord:
		return "Word"
	case ModeSentence:
		return "Sentence"
	case ModeParagraph:
		return "Paragraph"
	case ModeFull:
		return "Full"
	default:
		return string(m)
	}
}