package formatter

import "strings"

func splitWords(text string) []string {
	var units []string
	var cur strings.Builder
	inSpace := false
	started := false
	for _, r := range text {
		isSpace := r == ' ' || r == '\t' || r == '\n' || r == '\r'
		if !started {
			cur.WriteRune(r)
			inSpace = isSpace
			started = true
			continue
		}
		if isSpace == inSpace {
			cur.WriteRune(r)
			continue
		}
		if inSpace {
			cur.WriteRune(r)
			inSpace = false
			continue
		}
		units = append(units, cur.String())
		cur.Reset()
		cur.WriteRune(r)
		inSpace = true
	}
	if cur.Len() > 0 {
		units = append(units, cur.String())
	}
	return units
}