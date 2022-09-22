package common

import (
	"fmt"
	"time"
)

type Logger struct {}

const (
	error		string = "ERROR"
	warning		string = "WARN"
	info 		string = "INFO"
)

// NewLogger retuns a new instance of Logger, which encapsulates
// all logging functions.
func NewLogger() Logger {
	return Logger{}
}

// ErrorLog logs a normalized message of log level error
func (l Logger) ErrorLog(message string) {
	l.logMessage(message, error)
}

// WarningLog logs a normalized message of log level warning
func (l Logger) WarningLog(message string) {
	l.logMessage(message, warning)
}

// InfoLog logs a normalized message of log level info
func (l Logger) InfoLog(message string) {
	l.logMessage(message, info)
}

// NormalizedLog returns normalized message using the indicated log level.
func (l Logger) logMessage(message, logLevel string) {
	fmt.Println(l.normalizeMessage(message, logLevel))
}

// normalizeMessage returns a normalized string according to the system logs,
// using the log level and message received as arugments.
func (l Logger) normalizeMessage(message, logLevel string) string {
	return fmt.Sprintf("|%s|%s|%s", time.Now().String(), logLevel, message)
}
