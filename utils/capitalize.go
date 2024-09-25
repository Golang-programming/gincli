package utils

import (
	"strings"
)

func Capitalize(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}
