package sublist

import (
	"reflect"
)

const (
	relationEqual     = "equal"
	relationSublist   = "sublist"
	relationSuperlist = "superlist"
	relationUnequal   = "unequal"
)

// Relation defines the relation between two lists.
type Relation string

// Sublist determines the relation between two lists.
func Sublist(a []int, b []int) Relation {
	if len(a) == len(b) {
		if reflect.DeepEqual(a, b) {
			return relationEqual
		}
		return relationUnequal
	}

	// Empty lists are sublists of nonempty lists.
	if len(a) == 0 {
		return relationSublist
	}

	if len(b) == 0 {
		return relationSuperlist
	}

	// Keep track of the shorter list to know when to
	// return sub or super list in the case of a match.
	shorter, longer := a, b
	listAShorter := true

	if !(len(a) < len(b)) {
		listAShorter = false
		shorter, longer = b, a
	}

	for i := 0; i <= (len(longer) - len(shorter)); i++ {
		sublist := longer[i : i+len(shorter)]

		if reflect.DeepEqual(shorter, sublist) {
			if listAShorter {
				return relationSublist
			}
			return relationSuperlist
		}
	}

	return relationUnequal
}
