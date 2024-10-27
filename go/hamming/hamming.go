// Package hamming performs hamming distance calculation.
package hamming

import "errors"

// Distance returns the hamming distance between DNA strands 
// represented by a and b.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands %s and %s have inequal lengths")
	}

	// Each difference at an index adds one to the distance.
	var distance int 
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		distance++
	}

	return distance, nil
}
