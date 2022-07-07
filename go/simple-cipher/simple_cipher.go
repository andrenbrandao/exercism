package cipher

import (
	"strings"
	"unicode"
)

type shift int
type vigenere []int

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
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			lowercase := unicode.ToLower(char)
			shiftedChar := shiftChar(byte(lowercase), int(c))
			res.WriteByte(shiftedChar)
		}
	}

	return res.String()
}

func (c shift) Decode(input string) string {
	res := strings.Builder{}

	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			lowercase := unicode.ToLower(char)
			shiftedChar := shiftChar(byte(lowercase), -1*int(c))
			res.WriteByte(shiftedChar)
		}
	}

	return res.String()
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	aCount := 0
	res := make(vigenere, 0, len(key))

	for _, char := range key {
		if char < 'a' || char > 'z' {
			return nil
		}

		if char == 'a' {
			aCount++
		}

		res = append(res, int(char-'a'))
	}

	if aCount == len(key) {
		return nil
	}

	return res
}

func (v vigenere) Encode(input string) string {
	res := strings.Builder{}

	convertPos := 0
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			lowercase := unicode.ToLower(char)
			viginerePos := convertPos % len(v)
			shiftedChar := shiftChar(byte(lowercase), v[viginerePos])
			res.WriteByte(shiftedChar)
			convertPos++
		}
	}

	return res.String()
}

func (v vigenere) Decode(input string) string {
	res := strings.Builder{}

	convertPos := 0
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			lowercase := unicode.ToLower(char)
			viginerePos := convertPos % len(v)
			shiftedChar := shiftChar(byte(lowercase), -v[viginerePos])
			res.WriteByte(shiftedChar)
			convertPos++
		}
	}

	return res.String()
}

func shiftChar(char byte, val int) byte {
	charInt := int(char)
	shiftedChar := byte((charInt-int('a')+val%modulus+modulus)%modulus + int('a'))

	return shiftedChar
}
