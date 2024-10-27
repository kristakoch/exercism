package atbash

import (
	"strings"
	"unicode"
)

const groupSize = 5

// Atbash returns the atbash cipher of an input string.
func Atbash(s string) string {
	var cipher string

	s = strings.ToLower(s)

	var counter int
	for _, r := range []rune(s) {
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			continue
		}

		if unicode.IsLetter(r) {
			cipher += string('z' - (r - 'a'))
		} else {
			cipher += string(r)
		}

		if counter++; counter%groupSize == 0 {
			cipher += " "
		}
	}

	return strings.TrimSpace(cipher)
}
