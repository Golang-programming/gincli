package utils

import (
	"regexp"
	"strings"
)

func ConvertToSnakeCase(text string) string {
	// Regular expression to match spaces or punctuation
	reg := regexp.MustCompile(`[^\w]+`)

	// Replace spaces and punctuation with underscores
	snake := reg.ReplaceAllString(text, "_")

	// Insert underscore before any uppercase letters followed by lowercase letters
	snake = regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(snake, `${1}_${2}`)

	// Convert everything to lowercase
	return strings.ToLower(snake)
}
