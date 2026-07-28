package main

// isClosingMark reports whether r is a closing quote or bracket that
// should stay attached to the sentence-terminating punctuation before it.
func isClosingMark(r rune) bool {
	switch r {
	case '"', '\'', ')', ']', '}', '”', '’', '»':
		return true
	default:
		return false
	}
}
