package telegram

import "strings"

func lastIndexAfter(s, sep string) int {
	idx := strings.LastIndex(s, sep)
	if idx < 0 {
		return -1
	}
	return len([]rune(s[:idx+len(sep)]))
}