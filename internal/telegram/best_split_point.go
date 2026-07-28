package telegram

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