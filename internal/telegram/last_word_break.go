package telegram

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