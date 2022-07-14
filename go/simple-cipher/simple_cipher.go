package cipher

import (
	"strings"
	"unicode"
)

type shift int32
type vigenere []int32

const (
	defaultShift = 3
	modulus      = 26
)

func NewCaesar() Cipher {
	return shift(defaultShift)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	res := strings.Builder{}

	for _, char := range input {
		if unicode.IsLetter(char) {
			lowercase := unicode.ToLower(char)
			shiftedChar := shiftChar(lowercase, int32(c))
			res.WriteRune(shiftedChar)
		}
	}

	return res.String()
}

func (c shift) Decode(input string) string {
	c = -c
	return c.Encode(input)
}

func NewVigenere(key string) Cipher {
	if !isValidViginere(key) {
		return nil
	}

	res := make(vigenere, 0, len(key))

	for _, char := range key {
		res = append(res, char-'a')
	}

	return res
}

func isValidViginere(key string) bool {
	return len(key) != 0 && !hasOnlyAs(key) && !hasInvalidChars(key)
}

func hasOnlyAs(key string) bool {
	for _, char := range key {
		if char != 'a' {
			return false
		}
	}

	return true
}

func hasInvalidChars(key string) bool {
	for _, char := range key {
		if char < 'a' || char > 'z' {
			return true
		}
	}

	return false
}

func (v vigenere) Encode(input string) string {
	res := strings.Builder{}

	convertPos := 0
	for _, char := range input {
		if unicode.IsLetter(char) {
			lowercase := unicode.ToLower(char)
			viginerePos := convertPos % len(v)
			shiftedChar := shiftChar(lowercase, v[viginerePos])
			res.WriteRune(shiftedChar)
			convertPos++
		}
	}

	return res.String()
}

func (v vigenere) Decode(input string) string {
	res := strings.Builder{}

	convertPos := 0
	for _, char := range input {
		if unicode.IsLetter(char) {
			lowercase := unicode.ToLower(char)
			viginerePos := convertPos % len(v)
			shiftedChar := shiftChar(lowercase, -v[viginerePos])
			res.WriteRune(shiftedChar)
			convertPos++
		}
	}

	return res.String()
}

func shiftChar(char rune, val int32) rune {
	shiftedChar := (char-'a'+val%modulus+modulus)%modulus + 'a'

	return shiftedChar
}
