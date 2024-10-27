package transpose

import (
	"strings"
)

// Transpose transposes a matrix in the form of a string slice.
func Transpose(rows []string) []string {

	// Get the x and y dimensions by looking at slice and string lengths.
	dx := len(rows)
	var dy int
	for _, row := range rows {
		if len(row) > dy {
			dy = len(row)
		}
	}

	// Initialize the matrix with empty strings for each element.
	ret := make([][]string, dy)
	for i := range ret {
		ret[i] = make([]string, dx)
	}

	// Build out the matrix, adding left padding as necessary.
	for ridx, row := range rows {
		for cidx, ch := range row {
			for j := 0; j < ridx; j++ {
				if ret[cidx][j] == "" {
					ret[cidx][j] = " "
				}
			}
			ret[cidx][ridx] = string(ch)
		}
	}

	// Convert the matrix back into a slice by stringifying rows.
	transposed := make([]string, dy)
	for i, rs := range ret {
		transposed[i] = strings.Join(rs, "")
	}

	return transposed
}
