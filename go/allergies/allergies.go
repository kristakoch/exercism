package allergies

// substances is a slice of substance names
// indexed by substance value.
var substances = []string{
	"eggs",         // 1
	"peanuts",      // 2
	"shellfish",    // 4
	"strawberries", // 8
	"tomatoes",     // 16
	"chocolate",    // 32
	"pollen",       // 64
	"cats",         // 128
}

// Allergies returns a list of allergies based on
// a given allergy score.
func Allergies(score uint) []string {
	var allergies []string

	pow := 0
	for pow <= 7 {
		allergic := AllergicTo(score, substances[pow])
		if allergic {
			allergies = append(allergies, substances[pow])
		}
		pow++
	}

	return allergies
}

// AllergicTo returns whether a given score
// signifies an allergy to a given substance.
func AllergicTo(score uint, substance string) bool {
	for idx := range substances {
		if substances[idx] == substance {
			substanceValue := uint(1 << idx)
			return score&substanceValue > 0
		}
	}
	return false
}
