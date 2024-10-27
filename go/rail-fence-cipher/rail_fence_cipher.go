package railfence

import (
	"strings"
)

// Encode encodes a string using the rail fence cypher.
func Encode(input string, numRails int) string {
	input = strings.ReplaceAll(input, " ", "")

	// Add letters one at a time to the rail they belong in.
	rails := make([]string, numRails)

	adder, railNum := 1, 0
	for _, r := range input {
		rails[railNum] += string(r)
		if railNum += adder; railNum == 0 || railNum == numRails-1 {
			adder = -adder
		}
	}

	// Join all of the rails into one encoded string.
	return strings.Join(rails, "")
}

// Decode decodes a rail fence cypher encoded string.
func Decode(input string, numRails int) string {

	// Determine the length of each rail.
	railLengths := make([]int, numRails)

	adder, railNum := 1, 0
	for range input {
		railLengths[railNum]++
		if railNum += adder; railNum == 0 || railNum == numRails-1 {
			adder = -adder
		}
	}

	// Fill in the runes for each rail based on rail length.
	rails := make([][]rune, numRails)

	segmentStart := 0
	for idx, railLength := range railLengths {
		rails[idx] = []rune(input)[segmentStart : segmentStart+railLength]
		segmentStart += railLength
	}

	// Add the runes to the result in order, one by one.
	decoded := make([]rune, len(input))

	adder, railNum = 1, 0
	for idx := range input {
		decoded[idx] = rails[railNum][0]
		rails[railNum] = rails[railNum][1:]
		if railNum += adder; railNum == 0 || railNum == numRails-1 {
			adder = -adder
		}
	}

	return string(decoded)
}
