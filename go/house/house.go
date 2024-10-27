package house

import (
	"fmt"
	"strings"
)

type verse struct {
	subj string
	verb string
}

var verses = []verse{
	{subj: "house", verb: "Jack built"},
	{subj: "malt", verb: "lay in"},
	{subj: "rat", verb: "ate"},
	{subj: "cat", verb: "killed"},
	{subj: "dog", verb: "worried"},
	{subj: "cow with the crumpled horn", verb: "tossed"},
	{subj: "maiden all forlorn", verb: "milked"},
	{subj: "man all tattered and torn", verb: "kissed"},
	{subj: "priest all shaven and shorn", verb: "married"},
	{subj: "rooster that crowed in the morn", verb: "woke"},
	{subj: "farmer sowing his corn", verb: "kept"},
	{subj: "horse and the hound and the horn", verb: "belonged to"},
}

// Song returns the full song.
func Song() string {
	var song []string
	for i := 0; i < 12; i++ {
		song = append(song, Verse(i+1))
	}

	return strings.Join(song, "\n\n")
}

// Verse returns the nth verse.
func Verse(n int) string {
	return "This is " + recurse(n)
}

// recurse builds the nth verse recursively.
func recurse(n int) string {
	parts := []string{
		fmt.Sprintf("the %s", verses[n-1].subj),
		fmt.Sprintf("that %s", verses[n-1].verb),
	}

	if n == 1 {
		return strings.Join(parts, " ") + "."
	}

	return strings.Join(parts, "\n") + " " + recurse(n-1)
}
