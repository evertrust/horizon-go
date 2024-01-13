package log

import "log"

var LogEnabled = false

func Debug(message string) {
	if LogEnabled {
		log.Println("[DEBUG]", message)
	}
}
