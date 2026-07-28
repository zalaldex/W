package main

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
