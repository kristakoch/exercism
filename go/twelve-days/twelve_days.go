package twelve

import (
	"fmt"
	"strings"
)

var days = []struct {
	dayNum string
	gift   string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song returns the lyrics of the twelve days of Christmas.
func Song() string {
	var lyrics []string

	for i := 1; i <= 12; i++ {
		lyrics = append(lyrics, Verse(i))
	}

	return strings.Join(lyrics, "\n")
}

// Verse returns verses by day number.
func Verse(verseDay int) string {
	var verseGifts string

	for i := verseDay - 1; i >= 0; i-- {
		if i == 0 && ((verseDay - 1) != 0) {
			verseGifts += ", and "
		} else if i != (verseDay - 1) {
			verseGifts += ", "
		}

		verseGifts += days[i].gift
	}

	verse := fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me: %s.",
		days[verseDay-1].dayNum,
		verseGifts,
	)

	return verse
}
