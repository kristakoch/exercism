package strain

// Ints is a type for a slice of ints.
type Ints []int

// Strings is a type for a slice of strings.
type Strings []string

// Lists is a type for a slice of int slices.
type Lists [][]int

// Keep filters an Ints into ints that meet the predicate.
func (is Ints) Keep(predicate func(int) bool) Ints {
	var keeping Ints
	for _, i := range is {
		if predicate(i) {
			keeping = append(keeping, i)
		}
	}
	return keeping
}

// Discard filters an Ints into ints that don't meet the predicate.
func (is Ints) Discard(predicate func(int) bool) Ints {
	var discarding Ints
	for _, i := range is {
		if !predicate(i) {
			discarding = append(discarding, i)
		}
	}
	return discarding
}

// Keep filters a Strings into strings that meet the predicate.
func (ss Strings) Keep(predicate func(string) bool) Strings {
	var keeping Strings
	for _, s := range ss {
		if predicate(s) {
			keeping = append(keeping, s)
		}
	}
	return keeping
}

// Keep filters a Lists into int slices that meet the predicate.
func (ls Lists) Keep(predicate func([]int) bool) Lists {
	var keeping Lists
	for _, l := range ls {
		if predicate(l) {
			keeping = append(keeping, l)
		}
	}
	return keeping
}
