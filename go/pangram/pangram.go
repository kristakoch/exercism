package pangram

import (
	"strings"
	"unicode"
)

// IsPangram determines if a sentence is a pangram.
func IsPangram(str string) bool {
	if len(str) < 26 {
		return false
	}

	alphabetMap := make(map[string]bool)

	for _, r := range str {
		if unicode.IsLetter(r) {
			alphabetMap[strings.ToLower(string(r))] = true

			if len(alphabetMap) == 26 {
				return true
			}
		}
	}

	return false
}
