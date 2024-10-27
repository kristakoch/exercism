package armstrong

import (
	"math"
)

// IsNumber determines whether a number n is an armstrong number,
// a number that is the sum of its own digits each raised to the
// power of the number of digits.
func IsNumber(n int) bool {
	if n < 10 {
		return true
	}

	var digits []int
	remainder := n
	for remainder > 0 {
		for remainder > 0 {
			digits = append(digits, remainder-(remainder/10)*10)
			remainder /= 10
		}
	}

	var product float64
	for _, d := range digits {
		product += math.Pow(float64(d), float64(len(digits)))
	}

	return int(product) == n
}
