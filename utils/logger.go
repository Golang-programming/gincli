// ./utils/logger.go
package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// LogInfo logs informational messages in cyan.
func LogInfo(message string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(message))
}

// LogSuccess logs success messages in green.
func LogSuccess(message string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(message))
}

// LogWarning logs warning messages in yellow.
func LogWarning(message string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow(message))
}

// LogError logs error messages in red and exits the application.
func LogError(message string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red(message))
	os.Exit(1)
}
