package protein

import (
	"github.com/pkg/errors"
)

// ErrStop is an error returned when a terminating codon
// has been reached during protein translation.
var ErrStop = errors.New("stop")

// ErrInvalidBase is an error returned when an unknown
// codon has been reached during protein translation.
var ErrInvalidBase = errors.New("invalid base")

// FromCodon translates an input codon to a protein.
func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return protein, nil
}

// FromRNA translates an input RNA string made up of
// codons into a protein sequence.
func FromRNA(rnaString string) ([]string, error) {
	var err error
	var proteinSequence []string

	for i := 0; i < len(rnaString); i += 3 {
		var protein string

		codon := rnaString[i : i+3]
		if protein, err = FromCodon(codon); nil != err {
			switch {
			case errors.Is(err, ErrInvalidBase):
				return proteinSequence, ErrInvalidBase
			case errors.Is(err, ErrStop):
				return proteinSequence, nil
			}
		}

		proteinSequence = append(proteinSequence, protein)
	}

	return proteinSequence, nil
}
