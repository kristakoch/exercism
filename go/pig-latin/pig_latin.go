package piglatin

import (
	"strings"
)

// Sentence translates a sentence into pig latin.
func Sentence(input string) string {
	words := strings.Split(input, " ")

	var translated []string
	for _, word := range words {
		translated = append(translated, translate(word))
	}

	return strings.Join(translated, " ")
}

// translate translates a word into pig latin.
func translate(word string) string {
	if len(word) == 1 {
		return word
	}

	var prefixLen int
	for i, r := range strings.ToLower(word) {

		// Treat y as a consonant if it's the first letter, not followed by t.
		if i == 0 && r == 'y' && word[1] != 't' {
			prefixLen++
			continue
		}

		// Break on non-consonants, accounting for the u if a q precedes it.
		if isVowel(r) || word[:2] == "xr" || word[:2] == "yt" {
			if i > 0 && word[i-1] == 'q' && r == 'u' {
				prefixLen++
			}
			break
		}

		prefixLen++
	}

	return word[prefixLen:] + word[:prefixLen] + "ay"
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouy", r)
}
