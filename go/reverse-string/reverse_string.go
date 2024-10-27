// Package reverse reverses strings.
package reverse

// Reverse takes a string and returns its reversal.
func Reverse(str string) string {

	// Use a rune slice for the reversal to handle 
	// non-ascii multi-byte characters, ex: 世, 界
	var reversed []rune
	stringRunes := []rune(str)

	for i := len(stringRunes) - 1; i >= 0; i-- {
		reversed = append(reversed, stringRunes[i])
	}

	return string(reversed)
}
