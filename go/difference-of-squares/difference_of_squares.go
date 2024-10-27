// Package diffsquares computes the difference between square operations.
package diffsquares

// SquareOfSum returns the squared sum of 1 up to n.
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i < n+1; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns the sum of 1 squared up to n squared.
func SumOfSquares(n int) int {
	var sum int 
	for i := 1; i < n+1; i++ {
		sum += (i * i)
	}
	
	return sum
}

// Difference returns the difference between the square of
// the sum of n and the sum of squares of n.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
