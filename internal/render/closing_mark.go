package render

// isClosingMark reports whether r is a closing quote or bracket that should
// be considered part of a sentence-ending punctuation cluster.
func isClosingMark(r rune) bool {
	switch r {
	case '\'', '"', ')', ']', '}' :
		return true
	default:
		return false
	}
}
