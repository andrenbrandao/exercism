package strand

import "strings"

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	sb := strings.Builder{}

	for _, nucleotide := range dna {
		sb.WriteRune(dnaToRna[nucleotide])
	}

	return sb.String()
}
