package render

import "strings"

// Render converts text into Telegram monospace formatting according to
// mode. Each unit (word, sentence, paragraph, or the whole text) is wrapped
// individually in its own code span, separated by the same whitespace that
// originally separated the units, so content and spacing are preserved.
func Render(mode string, text string) string {
	switch mode {
	case "word":
		return renderUnits(text, splitWords)
	case "sentence":
		return renderUnits(text, splitSentences)
	case "paragraph":
		return renderUnits(text, splitParagraphs)
	case "full":
		fallthrough
	default:
		if strings.TrimSpace(text) == "" {
			return text
		}
		return wrapCode(text)
	}
}
