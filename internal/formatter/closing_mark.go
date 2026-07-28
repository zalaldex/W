package formatter

func IsClosingMark(r rune) bool {
	switch r {
	case '"', '\'', ')', ']', '}', '”', '’', '»':
		return true
	default:
		return false
	}
}