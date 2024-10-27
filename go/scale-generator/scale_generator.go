package scale

import "strings"

var chromaticSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromaticFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

// Scale generates the musical scale starting with the
// tonic and following the specified interval pattern.
func Scale(tonic string, interval string) []string {
	var scale []string

	var chromatic []string
	switch tonic {
	case "G", "D", "A", "E", "B", "F#","e", "b", "f#", "c#", "g#", "d#":
		chromatic = chromaticSharps
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chromatic = chromaticFlats
	default:
		chromatic = chromaticSharps
	}

	// Treat empty intervals as chromatics, with all half-steps.
	if len(interval) == 0 {
		interval = "mmmmmmmmmmmm"
	}

	idx := startingIndex(tonic, chromatic)

	// Step along the scale, adding a note before each step.
	for _, iv := range interval {
		scale = append(scale, chromatic[idx])
		switch string(iv) {
		case "m":
			idx = (idx + 1) % 12
		case "M":
			idx = (idx + 2) % 12
		case "A":
			idx = (idx + 3) % 12
		}
	}

	return scale
}

// startingIndex returns the starting index for a given 
// tonic in a given chromatic scale.
func startingIndex(tonic string, chromatic []string) int {
	for idx, note := range chromatic {
		if strings.EqualFold(tonic, note) {
			return idx
		}
	}
	return 0
}