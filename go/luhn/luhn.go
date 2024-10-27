// Package luhn uses the Luhn formula to validate numbers.
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns the luhn validation of the input.
func Valid(input string) bool {
	var err error

	// Input must be more than 1 in length with spaces removed.
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 2 {
		return false
	}

	// Range over the numbers, starting from the right.
	var sum int
	double := false
	for i := len(input) - 1; i >= 0; i-- {
		ch := string(input[i])

		var num int
		if num, err = strconv.Atoi(ch); nil != err {
			return false
		}

		// Double every other and subtract 9 if the result is over 9.
		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		double = !double

		// We'll need the sum of all the numbers to validate.
		sum += num
	}

	// The input is valid if it is evenly divisible by 10.
	return sum%10 == 0
}