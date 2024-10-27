package lsproduct

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct returns the largest product
// for a contiguous substring of digits of length n.
func LargestSeriesProduct(digitStr string, span int) (int, error) {
	var err error

	// Negative numbers and numbers larger than the
	// number of digits we have are invalid.
	if span > len(digitStr) || span < 0 {
		return 0, fmt.Errorf("span %d invalid for string %s", span, digitStr)
	}

	var digits []int
	for _, r := range digitStr {
		var num int
		if num, err = strconv.Atoi(string(r)); err != nil {
			return 0, err
		}

		digits = append(digits, num)
	}

	// Calculate combinations of length span,
	// keeping the largest product we find.
	var maxProduct int
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		for j := 0; j < span; j++ {
			product *= digits[i+j]
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
