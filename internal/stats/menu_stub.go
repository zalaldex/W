package stats

// statsMenu returns the inline menu displayed with statistics messages.
// The original repository kept a statsMenu in stats_menu.go; here we provide
// a minimal placeholder that callers in bot package can override if needed.
func statsMenu() interface{} { // using interface{} to avoid bringing tele dependency here
	// callers in bot package will pass the actual tele.ReplyMarkup
	return nil
}
