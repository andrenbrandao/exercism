package protein

import "errors"

var ErrStop = errors.New("Stop codon found")
var ErrInvalidBase = errors.New("Invalid codon found")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 2; i < len(rna); i += 3 {
		currProtein, err := FromCodon(rna[i-2 : i+1])

		if err == ErrStop {
			break
		}

		if err != nil {
			return []string{}, err
		}

		proteins = append(proteins, currProtein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	codonToProtein := map[string]string{
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

	if _, ok := codonToProtein[codon]; !ok {
		return "", ErrInvalidBase
	}

	if codonToProtein[codon] == "STOP" {
		return "", ErrStop
	}

	return codonToProtein[codon], nil
}
