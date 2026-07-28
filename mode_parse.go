package main

// ParseMode maps a settings button label back to a Mode.
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
