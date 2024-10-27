package summultiples

// SumMultiples returns the sum of all unique
// multiples of n up to but not including n.
func SumMultiples(limit int, divisors ...int) int {
	var sum int

	if len(divisors) == 0 {
		return sum
	}

	seen := make(map[int]bool)

	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if i < div || div == 0 || seen[i] {
				continue
			}
			if i%div == 0 {
				sum += i
				seen[i] = true
			}
		}
	}

	return sum
}
