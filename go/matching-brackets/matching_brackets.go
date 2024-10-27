package brackets

// Bracket determines whether a string containing brackets
// contains valid pairs, matched and nested correctly.
func Bracket(input string) bool {
	var leftBrackets []rune

	for _, r := range []rune(input) {
		switch r {
		case '{', '[', '(':
			leftBrackets = append(leftBrackets, r)

		case '}', ']', ')':
			if len(leftBrackets) == 0 {
				return false
			}
			leftBracket := leftBrackets[len(leftBrackets)-1]

			bracketPair := string(leftBracket) + string(r)
			if !(bracketPair == "{}" || bracketPair == "[]" || bracketPair == "()") {
				return false
			}

			leftBrackets = leftBrackets[:len(leftBrackets)-1]
		}
	}

	return len(leftBrackets) == 0
}
