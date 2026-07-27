package main

import "strings"

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

// telegramMessageLimit is Telegram's hard cap on text message length, in
// UTF-16 code units. We use it conservatively as a rune count, which is
// always <= the UTF-16 length and therefore always safe.
const telegramMessageLimit = 4096

// telegramCaptionLimit is Telegram's hard cap on media caption length.
const telegramCaptionLimit = 1024

// ModeLabel returns the human-readable label shown on settings buttons.
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

// Render converts text into Telegram monospace formatting according to
// mode. Each unit (word, sentence, paragraph, or the whole text) is wrapped
// individually in its own code span, separated by the same whitespace that
// originally separated the units, so content and spacing are preserved.
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

// renderUnits splits text into units with the given splitter, wraps each
// non-blank unit in its own code span, and reassembles them using the
// original separators so surrounding whitespace/newlines are preserved.
func renderUnits(text string, splitter func(string) []string) string {
	if strings.TrimSpace(text) == "" {
		return text
	}
	units := splitter(text)
	var b strings.Builder
	for _, u := range units {
		if strings.TrimSpace(u) == "" {
			b.WriteString(u)
			continue
		}
		lead, core, trail := splitSurroundingSpace(u)
		b.WriteString(lead)
		b.WriteString(wrapCode(core))
		b.WriteString(trail)
	}
	return b.String()
}

// wrapCode wraps s in a Telegram monospace code span. Backticks inside s
// are escaped so the span cannot be broken out of.
func wrapCode(s string) string {
	escaped := strings.ReplaceAll(s, "`", "'")
	return "`" + escaped + "`"
}

// splitSurroundingSpace separates leading/trailing whitespace from the core
// content of a unit, so wrapping only touches the visible text.
func splitSurroundingSpace(s string) (lead, core, trail string) {
	trimmedLeft := strings.TrimLeft(s, " \t\r\n")
	lead = s[:len(s)-len(trimmedLeft)]
	trimmedBoth := strings.TrimRight(trimmedLeft, " \t\r\n")
	trail = trimmedLeft[len(trimmedBoth):]
	core = trimmedBoth
	return
}

// splitWords splits text into words, keeping the whitespace between them
// attached to the following word so reassembly is lossless.
func splitWords(text string) []string {
	var units []string
	var cur strings.Builder
	inSpace := false
	started := false
	for _, r := range text {
		isSpace := r == ' ' || r == '\t' || r == '\n' || r == '\r'
		if !started {
			cur.WriteRune(r)
			inSpace = isSpace
			started = true
			continue
		}
		if isSpace == inSpace {
			cur.WriteRune(r)
			continue
		}
		if inSpace {
			// transition from space to word: keep space attached ahead
			cur.WriteRune(r)
			inSpace = false
			continue
		}
		// transition from word to space: flush word, start new unit with space
		units = append(units, cur.String())
		cur.Reset()
		cur.WriteRune(r)
		inSpace = true
	}
	if cur.Len() > 0 {
		units = append(units, cur.String())
	}
	return units
}

// splitSentences splits text into sentences, ending each unit after a
// sentence-terminating punctuation mark (. ! ?) plus any trailing quote or
// bracket, keeping following whitespace attached to the next sentence.
func splitSentences(text string) []string {
	var units []string
	runes := []rune(text)
	start := 0
	i := 0
	for i < len(runes) {
		r := runes[i]
		if r == '.' || r == '!' || r == '?' {
			j := i + 1
			for j < len(runes) && isClosingMark(runes[j]) {
				j++
			}
			units = append(units, string(runes[start:j]))
			start = j
			i = j
			continue
		}
		i++
	}
	if start < len(runes) {
		units = append(units, string(runes[start:]))
	}
	return units
}

func isClosingMark(r rune) bool {
	switch r {
	case '"', '\'', ')', ']', '}', '”', '’', '»':
		return true
	default:
		return false
	}
}

// splitParagraphs splits text on blank-line boundaries (one or more blank
// lines), keeping the separating newlines attached so spacing is preserved.
func splitParagraphs(text string) []string {
	var units []string
	var cur strings.Builder
	lines := strings.SplitAfter(text, "\n")
	for _, line := range lines {
		cur.WriteString(line)
		if strings.TrimSpace(line) == "" {
			units = append(units, cur.String())
			cur.Reset()
		}
	}
	if cur.Len() > 0 {
		units = append(units, cur.String())
	}
	return units
}

// SplitForTelegram splits rendered text into chunks that each fit within
// limit runes, preferring to break at paragraph, then sentence, then word,
// then character boundaries so content is never truncated or lost.
func SplitForTelegram(text string, limit int) []string {
	if len([]rune(text)) <= limit {
		return []string{text}
	}

	var chunks []string
	remaining := text
	for len([]rune(remaining)) > limit {
		cut := bestSplitPoint(remaining, limit)
		if cut <= 0 {
			cut = limit
		}
		runes := []rune(remaining)
		chunks = append(chunks, string(runes[:cut]))
		remaining = string(runes[cut:])
	}
	if len(remaining) > 0 {
		chunks = append(chunks, remaining)
	}
	return chunks
}

// bestSplitPoint returns a rune index <= limit at which to cut remaining,
// preferring (in order) a paragraph break, a sentence break, a word break,
// falling back to a hard character cut at limit.
func bestSplitPoint(remaining string, limit int) int {
	runes := []rune(remaining)
	window := string(runes[:limit])

	if i := lastIndexAfter(window, "\n\n"); i > 0 {
		return i
	}
	if i := lastSentenceBreak(window); i > 0 {
		return i
	}
	if i := lastWordBreak(window); i > 0 {
		return i
	}
	return limit
}

// lastWordBreak returns the rune index immediately after the last
// whitespace run within s, or -1 if none is found.
func lastWordBreak(s string) int {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if r == ' ' || r == '\t' || r == '\n' {
			return i + 1
		}
	}
	return -1
}

// lastIndexAfter returns the rune index immediately after the last
// occurrence of sep in s, or -1 if not found.
func lastIndexAfter(s, sep string) int {
	idx := strings.LastIndex(s, sep)
	if idx < 0 {
		return -1
	}
	return len([]rune(s[:idx+len(sep)]))
}

// lastSentenceBreak returns the rune index immediately after the last
// sentence-terminating punctuation (plus trailing closing marks and
// whitespace) within s, or -1 if none is found.
func lastSentenceBreak(s string) int {
	runes := []rune(s)
	best := -1
	i := 0
	for i < len(runes) {
		r := runes[i]
		if r == '.' || r == '!' || r == '?' {
			j := i + 1
			for j < len(runes) && isClosingMark(runes[j]) {
				j++
			}
			for j < len(runes) && (runes[j] == ' ' || runes[j] == '\n' || runes[j] == '\t') {
				j++
			}
			best = j
			i = j
			continue
		}
		i++
	}
	return best
}
