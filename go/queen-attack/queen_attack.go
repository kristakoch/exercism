package queenattack

import (
	"fmt"
	"math"
	"strconv"
)

// position is a type for 2d position data.
type position struct {
	x int
	y int
}

// CanQueenAttack returns whether queens in 2 positions
// are able to attack each other.
func CanQueenAttack(w, b string) (bool, error) {
	var err error

	// Input strings should be two characters and not the same.
	if w == b || len(b) != 2 || len(w) != 2 {
		return false, fmt.Errorf("input %s:%s invalid", w, b)
	}

	var wpos position
	wpos.x = int(([]rune(w[:1])[0] - 96))
	if wpos.y, err = strconv.Atoi(w[1:2]); nil != err {
		return false, err
	}
	var bpos position
	bpos.x = int(([]rune(b[:1])[0] - 96))
	if bpos.y, err = strconv.Atoi(b[1:2]); nil != err {
		return false, err
	}

	// Chess boards are 8 squares across, so that's our range.
	if wpos.x < 1 || wpos.y < 1 || wpos.x > 8 || wpos.y > 8 ||
		bpos.x < 1 || bpos.y < 1 || bpos.x > 8 || bpos.y > 8 {
		return false, fmt.Errorf("input %s:%s invalid", w, b)
	}

	// Matching x's means same row, matching y's means same col,
	// and an equal abs. val diff between x's and y's means diag.
	if wpos.x == bpos.x || wpos.y == bpos.y ||
		math.Abs(float64(wpos.x-bpos.x)) == math.Abs(float64(wpos.y-bpos.y)) {
		return true, nil
	}

	return false, nil
}
