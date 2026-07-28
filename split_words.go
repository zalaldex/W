package main

import "strings"

// splitWords splits text into words, keeping the whitespace between them
// attached to the following word so reassembly is lossless.
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
			// transition from space to word: keep space attached ahead
			cur.WriteRune(r)
			inSpace = false
			continue
		}
		// transition from word to space: flush word, start new unit with space
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
