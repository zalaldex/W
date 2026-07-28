package render

import "strings"

// splitSurroundingSpace separates leading/trailing whitespace from the core
// content of a unit, so wrapping only touches the visible text.
func splitSurroundingSpace(s string) (lead, core, trail string) {
	trimmedLeft := strings.TrimLeft(s, " \t\r\n")
	lead = s[:len(s)-len(trimmedLeft)]
	trimmedBoth := strings.TrimRight(trimmedLeft, " \t\r\n")
	trail = trimmedLeft[len(trimmedBoth):]
	core = trimmedBoth
	return
}
