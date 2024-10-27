package sieve

// Sieve uses the Sieve of Erastosthenes algorithm
// to find the primes from 2 to the input limit.
func Sieve(limit int) []int {
	marked := make(map[int]bool)
	for num := 2; num <= limit; num++ {
		marked[num] = false
	}

	var primes []int
	for num := 2; num <= limit; num++ {
		if marked[num] {
			continue
		}

		primes = append(primes, num)

		for multiple := num; multiple <= limit; multiple += num {
			marked[multiple] = true
		}
	}

	return primes
}
