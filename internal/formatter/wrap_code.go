package formatter

import "strings"

func wrapCode(s string) string {
	escaped := strings.ReplaceAll(s, "`", "'")
	return "`" + escaped + "`"
}