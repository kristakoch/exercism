package pythagorean

// Triplet holds data for triangle side lengths.
type Triplet [3]int

// Sum returns a list of all Pythagorean triplets
// where the sum a+b+c (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var ts []Triplet

	tsInRange := Range(1, p)

	for _, t := range tsInRange {
		if t[0]+t[1]+t[2] == p {
			ts = append(ts, t)
		}
	}

	return ts
}

// Range returns a list of all Pythagorean triplets
// with sides in the range min to max inclusive.
func Range(min, max int) []Triplet {
	var ts []Triplet

	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}

	return ts
}
