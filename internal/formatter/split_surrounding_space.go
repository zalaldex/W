package formatter

import "strings"

func splitSurroundingSpace(s string) (lead, core, trail string) {
	trimmedLeft := strings.TrimLeft(s, " \t\r\n")
	lead = s[:len(s)-len(trimmedLeft)]
	trimmedBoth := strings.TrimRight(trimmedLeft, " \t\r\n")
	trail = trimmedLeft[len(trimmedBoth):]
	core = trimmedBoth
	return
}