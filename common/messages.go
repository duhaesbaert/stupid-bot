package common

import (
	"fmt"
	"time"
)

const (
	error		string = "ERROR"
	warning		string = "WARN"
	info 		string = "INFO"
)

// NormalizedLog returns normalized message using the indicated log level.
func logMessage(message, logLevel string) {
	fmt.Println(normalizeMessage(message, logLevel))
}
func normalizeMessage(message, logLevel string) string {
	return fmt.Sprintf("|%s|%s|%s", time.Now().String(), logLevel, message)
}

func ErrorLog(message string) {
	logMessage(message, error)
}

func WarningLog(message string) {
	logMessage(message, warning)
}

func InfoLog(message string) {
	logMessage(message, info)
}
