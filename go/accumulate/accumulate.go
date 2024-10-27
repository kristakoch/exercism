package accumulate

// Accumulate returns the result of performing the
// given operation on each element of the input.
func Accumulate(
	input []string,
	converter func(string) string,
) []string {
	for i := range input {
		input[i] = converter(input[i])
	}

	return input
}
