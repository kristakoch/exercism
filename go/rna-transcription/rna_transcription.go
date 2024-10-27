package strand

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns a dna strand's rna complement.
func ToRNA(dna string) string {
	rnaNucleotides := make([]rune, len(dna))

	for i, dn := range []rune(dna) {
		rnaNucleotides[i] = complements[dn]
	}

	return string(rnaNucleotides)
}
