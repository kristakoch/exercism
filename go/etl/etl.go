package etl

import "strings"

// Transform takes in Scrabble scoring data by point
// and returns a map of the data by letter.
func Transform(legacyMap map[int][]string) map[string]int {
	output := make(map[string]int)

	for ptval, letters := range legacyMap {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = ptval
		}
	}

	return output
}
