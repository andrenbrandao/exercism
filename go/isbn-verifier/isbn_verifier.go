package isbn

import (
	"errors"
	"unicode"
)

/*

- Iterate over each character
- If not a digit, ignore it
- Start with a multiplier of 10 and multiply each digit
decrementing the multiplier by 1.
- Sum each digit * multiplier and mod 11 at the end.
- If the result is 0, this is a valid ISBN

Edge cases:
- If the last character is X, we consider it as a digit 10
- More than 10 digits makes it invalid
- Invalid digit, such as Y should be invalid

Time Complexity: O(n)
Space Complexity: O(1).

*/

// IsValidISBN checks if a ISBN-10 string is valid.
// BenchmarkIsValidISBN-8           2082508               555.7 ns/op             0 B/op          0 allocs/op
func IsValidISBN(isbn string) bool {
	charPos := 1
	total := 0

	for _, c := range isbn {
		switch {
		case c == '-':
			continue
		case !unicode.IsDigit(c) && c != 'X':
			return false
		case c == 'X' && charPos != 10:
			return false
		}

		val, err := getDigitVal(c)

		if err != nil {
			return false
		}

		total += val * (11 - charPos)
		charPos++
	}

	return charPos == 11 && total%11 == 0
}

func getDigitVal(c rune) (int, error) {
	switch {
	case unicode.IsDigit(c):
		return int(c) - '0', nil
	case c == 'X':
		return 10, nil
	default:
		return 0, errors.New("character is not a digit")
	}
}
