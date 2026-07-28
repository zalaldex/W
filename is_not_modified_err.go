package main

import "strings"

// isNotModifiedErr reports whether err is Telegram's "message is not
// modified" error, returned when editing a message with identical content.
func isNotModifiedErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()), "message is not modified")
}
