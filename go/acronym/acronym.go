// Package acronym performs acronym conversion.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	var acronym string

	// Break up the string by word-breaking characters.
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
	})

	// Build the acronym.
	for _, word := range words {
		firstLetter := string(word[0])
		acronym += strings.ToUpper(firstLetter)
	}

	return acronym
}
