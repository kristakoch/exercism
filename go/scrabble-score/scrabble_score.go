// Package scrabble calculates scrabble scores.
package scrabble

import "strings"

// letterScores holds scrabble letter values.
var scoreMap = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score returns the scrabble score for the input string.
func Score(word string) int {
	var score int 

	// Convert to upper to match the case in the score map.
	word = strings.ToUpper(word)

	for _, r := range word {
		for scoreGroup, points := range scoreMap {
			if strings.ContainsRune(scoreGroup,r) {
				score += points
				continue
			}
		}
	}

	return score
}