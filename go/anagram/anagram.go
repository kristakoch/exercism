package anagram

import (
	"reflect"
	"strings"
)

// Detect returns a sublist of the input candidates that are
// anagrams of the input subject.
func Detect(subject string, candidates []string) []string {
	var anagrams []string

	subjectRunes := make(map[rune]int)
	for _, r := range strings.ToLower(subject) {
		subjectRunes[r]++
	}

	for _, word := range candidates {
		if strings.ToLower(word) == strings.ToLower(subject) {
			continue
		}

		candidateRunes := make(map[rune]int)
		for _, r := range strings.ToLower(word) {
			candidateRunes[r]++
		}

		if reflect.DeepEqual(subjectRunes, candidateRunes) {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams
}
