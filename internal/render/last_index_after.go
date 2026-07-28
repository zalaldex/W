package render

// lastIndexAfter finds the last index of substr in s and returns the index
// immediately after it, or -1 if substr is not found.
func lastIndexAfter(s, substr string) int {
	last := -1
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			last = i + len(substr)
		}
	}
	return last
}
