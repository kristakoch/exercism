package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Song returns the full song.
func Song() string {
	s, _ := Verses(99, 0)
	return s
}

// Verses returns verses from the upper to the lower bound.
func Verses(ubound, lbound int) (string, error) {
	var err error

	if ubound <= lbound || ubound > 99 || lbound < 0 {
		return "", errors.New("bounds must be in format: upper, lower, and both in range of 0-99")
	}

	var vs []string
	for ubound >= lbound {
		var v string
		if v, err = Verse(ubound); nil != err {
			return "", err
		}
		vs = append(vs, v)
		ubound--
	}

	return strings.Join(vs, "\n") + "\n", nil
}

// Verse returns one verse of the song.
func Verse(n int) (string, error) {
	switch {
	case n > 99, n < 0:
		return "", errors.New("number must be in range of 0-99")
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}

	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
}
