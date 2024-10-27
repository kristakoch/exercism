package series

// All returns a list of all substrings of s
// with length n.
func All(n int, s string) []string {
	var substrings []string

	if len(s) < n {
		return substrings
	}

	for i := 0; i <= len(s)-n; i++ {
		substr := s[i : i+n]
		substrings = append(substrings, substr)
	}

	return substrings
}

// UnsafeFirst returns the first substring of s
// with length n.
func UnsafeFirst(n int, s string) string {
	res := All(n, s)

	if len(res) == 0 {
		return ""
	}

	return res[0]
}

// First returns the first substring of s with
// length n and whether the result is valid.
func First(n int, s string) (first string, ok bool) {
	res := UnsafeFirst(n, s)

	if res == "" {
		return "", false
	}

	return res, true
}
