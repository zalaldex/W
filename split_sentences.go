package main

// splitSentences splits text into sentences, ending each unit after a
// sentence-terminating punctuation mark (. ! ?) plus any trailing quote or
// bracket, keeping following whitespace attached to the next sentence.
func splitSentences(text string) []string {
	var units []string
	runes := []rune(text)
	start := 0
	i := 0
	for i < len(runes) {
		r := runes[i]
		if r == '.' || r == '!' || r == '?' {
			j := i + 1
			for j < len(runes) && isClosingMark(runes[j]) {
				j++
			}
			units = append(units, string(runes[start:j]))
			start = j
			i = j
			continue
		}
		i++
	}
	if start < len(runes) {
		units = append(units, string(runes[start:]))
	}
	return units
}
