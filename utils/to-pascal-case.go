package utils

import (
	"unicode"
)

func ToPascalCase(input string) string {
	var result []rune
	shouldCapitalize := true

	for _, r := range input {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			shouldCapitalize = true
		} else {
			if shouldCapitalize {
				result = append(result, unicode.ToUpper(r))
				shouldCapitalize = false
			} else {
				result = append(result, unicode.ToLower(r))
			}
		}
	}

	return string(result)
}
