// ./utils/capitalize.go
package utils

import (
	"strings"
	"unicode"
)

// Capitalize capitalizes the first letter of each word in the string.
func Capitalize(text string) string {
	if len(text) == 0 {
		return text
	}
	words := strings.Fields(text)
	for i, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, "")
}
