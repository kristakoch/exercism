package say

import (
	"strings"
)

var words = map[string][]string{
	"ones":   {"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
	"tens":   {"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"},
	"places": {"", "hundred", "thousand", "million", "billion"},
}

// Say translates numbers between 0 and 999,999,999,999
// into a number in english.
func Say(num int64) (string, bool) {
	if num < 0 || num > 999999999999 {
		return "", false
	}
	if num == 0 {
		return "zero", true
	}

	numGroups := breakIntoThousands(num)

	var numWords []string
	for idx, ng := range numGroups {
		word := translate3Digits(ng)
		if word == "" {
			continue
		}

		if place := len(numGroups) - idx; place > 1 {
			word += " " + words["places"][place]
		}

		numWords = append(numWords, word)
	}

	return strings.Join(numWords, " "), true
}

// breakIntoThousands breaks up a number into thousands.
func breakIntoThousands(num int64) []int64 {
	var thousands []int64
	for num > 0 {
		thousands = append([]int64{num % 1000}, thousands...)
		num /= 1000
	}
	return thousands
}

// translate3Digits translates a 3 digit number into english.
func translate3Digits(num int64) string {
	var translated string

	if num < 20 {
		return words["ones"][num]
	}

	if num < 100 {
		r := num % 10
		translated += words["tens"][(num-r)/10]
		if r != 0 {
			translated += "-" + words["ones"][r]
		}
		return translated
	}

	r := num % 100
	tens := translate3Digits(r)
	translated += words["ones"][num/100] + " " + words["places"][1]
	if tens != "" {
		translated += " " + tens
	}

	return translated
}
