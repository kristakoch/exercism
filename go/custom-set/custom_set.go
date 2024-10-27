package stringset

import (
	"fmt"
	"strings"
)

// Set is a type for sets.
type Set struct {
	items map[string]string
}

// New is a factory for sets.
func New() Set {
	var s Set
	s.items = make(map[string]string)
	return s
}

// String returns a Set's string representation.
func (s Set) String() string {
	var stringRep []string
	for _, si := range s.items {
		stringRep = append(stringRep, fmt.Sprintf("\"%s\"", si))
	}
	return fmt.Sprintf("{%s}", strings.Join(stringRep, ", "))
}

// NewFromSlice is a factory for creating a Set from a slice.
func NewFromSlice(strings []string) Set {
	s := New()
	for _, str := range strings {
		s.items[str] = str
	}
	return s
}

// Add adds an item to the Set.
func (s Set) Add(str string) {
	s.items[str] = str
}

// IsEmpty returns whether a Set is empty.
func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

// Has returns whether an item exists in a Set.
func (s Set) Has(str string) bool {
	_, ok := s.items[str]
	return ok
}

// Subset returns whether s1 is a subset of s2.
// This is all items in s1 being in s2.
func Subset(s1, s2 Set) bool {
	for _, s1i := range s1.items {
		if !s2.Has(s1i) {
			return false
		}
	}
	return true
}

// Disjoint returns whether s1 and s2 are disjoint.
// This is having no items in common.
func Disjoint(s1, s2 Set) bool {
	for _, s1i := range s1.items {
		if s2.Has(s1i) {
			return false
		}
	}
	return true
}

// Equal returns whether s1 and s2 are equal.
// This is having all items in common.
func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}
	return Subset(s1, s2)
}

// Intersection returns the intersection of s1 and s2.
// This is items that exist in both.
func Intersection(s1, s2 Set) Set {
	s := New()
	for _, s1i := range s1.items {
		if s2.Has(s1i) {
			s.Add(s1i)
		}
	}
	return s
}

// Difference returns the difference of s1 and s2.
// This is items that exist in either but not both.
func Difference(s1, s2 Set) Set {
	s := New()
	for _, s1i := range s1.items {
		if !s2.Has(s1i) {
			s.Add(s1i)
		}
	}

	for _, s2i := range s1.items {
		if !s1.Has(s2i) {
			s.Add(s2i)
		}
	}

	return s
}

// Union returns the union of s1 and s2.
// This is items that exist in either.
func Union(s1, s2 Set) Set {
	s := New()
	for _, s1i := range s1.items {
		s.Add(s1i)
	}
	for _, s2i := range s2.items {
		s.Add(s2i)
	}
	return s
}
