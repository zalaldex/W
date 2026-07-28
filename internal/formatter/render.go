package formatter

import "strings"

func Render(mode Mode, text string) string {
	switch mode {
	case ModeWord:
		return renderUnits(text, splitWords)
	case ModeSentence:
		return renderUnits(text, splitSentences)
	case ModeParagraph:
		return renderUnits(text, splitParagraphs)
	case ModeFull:
		fallthrough
	default:
		if strings.TrimSpace(text) == "" {
			return text
		}
		return wrapCode(text)
	}
}