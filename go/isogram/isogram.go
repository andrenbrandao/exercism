package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	letterSet := make(map[rune]bool)

	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		lowercaseLetter := unicode.ToLower(letter)

		if _, ok := letterSet[lowercaseLetter]; ok {
			return false
		}

		letterSet[lowercaseLetter] = true
	}

	return true
}
