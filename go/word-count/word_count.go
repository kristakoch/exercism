package wordcount

import (
	"regexp"
	"strings"
)

// isWordRegex is a regex for picking up
// simple words, contractions, and numbers.
const isWordRegex = "[A-Za-z]+('[A-Za-z])?|[0-9]"

// Frequency is a type to hold word frequency data.
type Frequency map[string]int

// WordCount returns the frequency of each
// word in the input phrase.
func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)

	rx := regexp.MustCompile(isWordRegex)
	words := rx.FindAllString(strings.ToLower(phrase), -1)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}
