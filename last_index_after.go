package main

import "strings"

// lastIndexAfter returns the rune index immediately after the last
// occurrence of sep in s, or -1 if not found.
func lastIndexAfter(s, sep string) int {
	idx := strings.LastIndex(s, sep)
	if idx < 0 {
		return -1
	}
	return len([]rune(s[:idx+len(sep)]))
}
