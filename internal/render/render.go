package render

import (
	"strings"

	"github.com/zalaldex/W/internal/mode"
)

// Render converts text into Telegram monospace formatting according to
// mode. Each unit (word, sentence, paragraph, or the whole text) is wrapped
// individually in its own code span, separated by the same whitespace that
// originally separated the units, so content and spacing are preserved.
func Render(m mode.Mode, text string) string {
	switch m {
	case mode.ModeWord:
		return renderUnits(text, splitWords)
	case mode.ModeSentence:
		return renderUnits(text, splitSentences)
	case mode.ModeParagraph:
		return renderUnits(text, splitParagraphs)
	case mode.ModeFull:
		fallthrough
	default:
		if strings.TrimSpace(text) == "" {
			return text
		}
		return wrapCode(text)
	}
}
