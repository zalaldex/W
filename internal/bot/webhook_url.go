package bot

// webhookURL builds the public URL Telegram will use to deliver updates,
// combining Railway's public domain with the bot token as the path secret.
func webhookURL(domain, token string) string {
	return "https://" + domain + "/" + token
}
