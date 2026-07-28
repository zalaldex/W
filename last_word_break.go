package main

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
