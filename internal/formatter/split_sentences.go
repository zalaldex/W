package formatter

func splitSentences(text string) []string {
	var units []string
	runes := []rune(text)
	start := 0
	i := 0
	for i < len(runes) {
		r := runes[i]
		if r == '.' || r == '!' || r == '?' {
			j := i + 1
			for j < len(runes) && IsClosingMark(runes[j]) {
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