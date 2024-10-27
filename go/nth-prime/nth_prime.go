package prime

import "math"

// Nth determines the nth prime based on an input n.
func Nth(n int) (int, bool) {
	if n == 0 {
		return n, false
	}

	number := 1
	prime := 2
	counter := 1

	for counter < n {
		number += 2
		if isPrime(number) {
			prime = number
			counter++
		}
	}

	return prime, true
}

// isPrime determines whether n is prime through the
// primality test of trial division.
// https://en.wikipedia.org/wiki/Trial_division
func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)) + 1)

	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
