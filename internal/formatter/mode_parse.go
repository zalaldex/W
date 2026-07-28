package formatter

func ParseMode(label string) (Mode, bool) {
	switch label {
	case "Word":
		return ModeWord, true
	case "Sentence":
		return ModeSentence, true
	case "Paragraph":
		return ModeParagraph, true
	case "Full":
		return ModeFull, true
	default:
		return "", false
	}
}