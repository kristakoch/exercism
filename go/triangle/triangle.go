package triangle

import "math"

// Kind is an enumerated type for triangles.
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	// Validate the input for non-triangles.
	switch {
	case a <= 0 || b <= 0 || c <= 0,
		a+b < c || a+c < b || b+c < a,
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c),
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1):
		return NaT
	}

	// Determine the type of triangle from the input.
	var k Kind
	switch {
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}

	return k
}
