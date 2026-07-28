package render

import "strings"

// wrapCode wraps s in a Telegram monospace code span. Backticks inside s
// are escaped so the span cannot be broken out of.
func wrapCode(s string) string {
	escaped := strings.ReplaceAll(s, "`", "'")
	return "`" + escaped + "`"
}
