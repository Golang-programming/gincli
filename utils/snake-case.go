// ./utils/snake-case.go
package utils

import (
	"regexp"
	"strings"
)

// ConvertToSnakeCase converts a string to snake_case.
func ConvertToSnakeCase(text string) string {
	// Replace non-word characters with underscores
	reg := regexp.MustCompile(`[^\w]+`)
	snake := reg.ReplaceAllString(text, "_")

	// Insert underscore before any uppercase letters followed by lowercase letters
	snake = regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(snake, `${1}_${2}`)

	// Convert everything to lowercase
	return strings.ToLower(snake)
}
