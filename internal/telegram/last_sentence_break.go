package telegram

import "telegram-monospace-bot/internal/formatter"

func lastSentenceBreak(s string) int {
	runes := []rune(s)
	best := -1
	i := 0
	for i < len(runes) {
		r := runes[i]
		if r == '.' || r == '!' || r == '?' {
			j := i + 1
			for j < len(runes) && formatter.IsClosingMark(runes[j]) {
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