package main

import (
	"log"
	"time"
)

// trackMessage records one incoming Telegram update as a message event for
// userID, and upserts that user's first/last seen timestamps. It is safe to
// call from every update handler (text, caption, media, sticker, document,
// etc.) — each call counts as exactly one message.
func trackMessage(userID int64) {
	now := time.Now().UTC()

	if _, err := DB.Exec(
		`INSERT INTO users (user_id, first_seen, last_seen) VALUES (?, ?, ?)
		 ON CONFLICT(user_id) DO UPDATE SET last_seen = excluded.last_seen`,
		userID, now, now,
	); err != nil {
		log.Printf("trackMessage: upsert user %d: %v", userID, err)
		return
	}

	if _, err := DB.Exec(
		`INSERT INTO messages (user_id, created_at) VALUES (?, ?)`,
		userID, now,
	); err != nil {
		log.Printf("trackMessage: insert message for user %d: %v", userID, err)
	}
}
