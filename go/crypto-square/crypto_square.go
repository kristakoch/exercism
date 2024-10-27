package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode outputs the square encoding of the input.
func Encode(input string) string {
	var clean string
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			clean += strings.ToLower(string(r))
		}
	}

	// Divide the string evenly into cols using the length's sq root
	// and then use that number to determine how many rows we need.
	numChars := float64(len(clean))
	numCols := int(math.Ceil(math.Sqrt(numChars)))
	numRows := int(math.Ceil(numChars / float64(numCols)))

	// Place letters one at a time into their row in the square
	// code slice based on their column offset as the index.
	squareCode := make([]string, numCols)
	for i := 0; i < numCols*numRows; i++ {
		rowNum := i % numCols
		if i > len(clean)-1 {
			squareCode[rowNum] += " "
			continue
		}
		squareCode[rowNum] += string(clean[i])
	}

	return strings.Join(squareCode, " ")
}
