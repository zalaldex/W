package formatter

import "strings"

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