// Package raindrops generates raindrop sounds.
package raindrops

import "fmt"

// Convert takes a number and returns raindrop sounds
// based on factors or 3, 5, and 7.
func Convert(num int) string {
	var raindropStr string

	hasFactorThree := num%3 == 0
	hasFactorFive := num%5 == 0
	hasFactorSeven := num%7 == 0

	if hasFactorThree {
		raindropStr += "Pling"
	}
	if hasFactorFive {
		raindropStr += "Plang"
	}
	if hasFactorSeven {
		raindropStr += "Plong"
	}

	if raindropStr == "" {
		raindropStr = fmt.Sprintf("%d", num)
	}

	return raindropStr
}
