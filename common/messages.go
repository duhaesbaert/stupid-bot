package common

import (
	"fmt"
	"time"
)

const (
	Error		string = "ERROR"
	Warning		string = "WARNING"
	Info 		string = "INFO"
)

// NormalizedLog returns normalized message using the indicated log level.
func NormalizedLog(message string, logLevel string) {
	fmt.Println(fmt.Sprintf("|%t|%s|%s", time.Now().String(), logLevel, message))
}