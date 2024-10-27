package prime

// Factors finds the prime factorization of a number n.
func Factors(n int64) []int64 {
	factors := make([]int64, 0)

	var divisor int64 = 2
	for n > 1 {
		if n%divisor == 0 {
			n /= divisor
			factors = append(factors, divisor)
			continue
		}
		divisor++
	}

	return factors
}
