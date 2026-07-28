package main

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
