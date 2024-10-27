package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on the
// nth square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("number not in range of 1-64")
	}

	squareTotal := math.Pow(2, float64(n-1))

	return uint64(squareTotal), nil
}

// Total returns the total number of grains
// on the chessboard.
func Total() uint64 {

	var total uint64
	for n := 1; n < 65; n++ {

		// Don't need to handle the error since
		// we know the input will be valid.
		squareTotal, _ := Square(n)

		total += squareTotal
	}

	return total
}
