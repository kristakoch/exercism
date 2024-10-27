// Package isogram determines isograms from inputs.
package isogram

import "strings"

// IsIsogram returns true if the word or phrase 
// is an isogram and false otherwise.
func IsIsogram(str string) bool {

	// Init map for found letters.
	foundLetters := make(map[rune]bool)

	// Convert string to lower for case-insensitivity.
	str = strings.ToLower(str)

	for _, r := range str {
		if r == ' ' || r == '-' {
			continue
		}
		if _, ok := foundLetters[r]; ok {
			return false
		}
		
		foundLetters[r] = true
	}

	return true
}