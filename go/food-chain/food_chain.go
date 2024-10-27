package foodchain

import (
	"fmt"
	"strings"
)

var songText = []struct {
	noun   string
	phrase string
}{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

// Song returns the full food chain song.
func Song() string {
	return Verses(1, 8)
}

// Verses returns verses n to m.
func Verses(n, m int) string {

	// The range numbers should be in ascending order.
	if n > m {
		n, m = m, n
	}

	var verses []string
	for i := n; i < m+1; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}

// Verse returns one verse of the food chain song.
func Verse(n int) string {
	var verse []string

	verse = append(verse, fmt.Sprintf("I know an old lady who swallowed a %s.", songText[n-1].noun))
	verse = append(verse, songText[n-1].phrase)

	// The first and last verses end after the above step.
	if n == 1 || n == 8 {
		return strings.Join(verse, "\n")
	}

	for n > 1 {
		if n == 3 {
			verse = append(verse, fmt.Sprintf("She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.", songText[n-1].noun, songText[n-2].noun))

			n--
			continue
		}

		verse = append(verse, fmt.Sprintf("She swallowed the %s to catch the %s.", songText[n-1].noun, songText[n-2].noun))
		n--
	}
	verse = append(verse, songText[0].phrase)

	return strings.Join(verse, "\n")
}
