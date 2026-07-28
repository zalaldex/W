package render

import "strings"

// renderUnits splits text into units with the given splitter, wraps each
// non-blank unit in its own code span, and reassembles them using the
// original separators so surrounding whitespace/newlines are preserved.
func renderUnits(text string, splitter func(string) []string) string {
	if strings.TrimSpace(text) == "" {
		return text
	}
	units := splitter(text)
	var b strings.Builder
	for _, u := range units {
		if strings.TrimSpace(u) == "" {
			b.WriteString(u)
			continue
		}
		lead, core, trail := splitSurroundingSpace(u)
		b.WriteString(lead)
		b.WriteString(wrapCode(core))
		b.WriteString(trail)
	}
	return b.String()
}
