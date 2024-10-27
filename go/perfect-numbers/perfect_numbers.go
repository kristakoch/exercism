package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("input is not a positive integer")

// Classify determines the classification of a number
// based on Nichomachus's scheme for natural numbers.
func Classify(number int64) (Classification, error) {
	var c Classification

	if number < 1 {
		return c, ErrOnlyPositive
	}

	var aliquotSum int64
	for i := number / 2; i > 0; i-- {
		if number%i == 0 {
			aliquotSum += i
		}
	}

	switch {
	case aliquotSum == number:
		c = ClassificationPerfect
	case aliquotSum > number:
		c = ClassificationAbundant
	case aliquotSum < number:
		c = ClassificationDeficient
	}

	return c, nil
}
