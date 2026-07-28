package formatter

import "strings"

func splitParagraphs(text string) []string {
	var units []string
	var cur strings.Builder
	lines := strings.SplitAfter(text, "\n")
	for _, line := range lines {
		cur.WriteString(line)
		if strings.TrimSpace(line) == "" {
			units = append(units, cur.String())
			cur.Reset()
		}
	}
	if cur.Len() > 0 {
		units = append(units, cur.String())
	}
	return units
}