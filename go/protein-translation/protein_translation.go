package protein

import "errors"

var ErrStop = errors.New("stop codon found")
var ErrInvalidBase = errors.New("invalid codon found")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 0; i < len(rna)-2; i += 3 {
		currProtein, err := FromCodon(rna[i : i+3])

		if errors.Is(err, ErrStop) {
			break
		}

		if err != nil {
			return []string{}, err
		}

		proteins = append(proteins, currProtein)
	}

	return proteins, nil
}

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) (string, error) {
	if _, ok := codonToProtein[codon]; !ok {
		return "", ErrInvalidBase
	}

	if codonToProtein[codon] == "STOP" {
		return "", ErrStop
	}

	return codonToProtein[codon], nil
}
