package telegram

func splitForTelegram(text string, limit int) []string {
	if len([]rune(text)) <= limit {
		return []string{text}
	}

	var chunks []string
	remaining := text
	for len([]rune(remaining)) > limit {
		cut := bestSplitPoint(remaining, limit)
		if cut <= 0 {
			cut = limit
		}
		runes := []rune(remaining)
		chunks = append(chunks, string(runes[:cut]))
		remaining = string(runes[cut:])
	}
	if len(remaining) > 0 {
		chunks = append(chunks, remaining)
	}
	return chunks
}