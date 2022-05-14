package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	strippedNumbers := StripSpaces(id)

	if len(strippedNumbers) <= 1 {
		return false
	}

	if HasNonDigits(strippedNumbers) {
		return false
	}

	modifiedSecondDigits := DoubleSecondDigits(strippedNumbers)
	sumOfDigits := SumDigits(modifiedSecondDigits)
	return sumOfDigits%10 == 0
}

func StripSpaces(id string) string {
	strippedNumbers := strings.ReplaceAll(id, " ", "")
	return strippedNumbers
}

func HasNonDigits(id string) bool {
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func DoubleSecondDigits(id string) string {
	var res strings.Builder

	for i := len(id) - 1; i >= 0; i-- {
		isSecondDigit := (len(id)-1-i)%2 == 1
		digit, _ := strconv.Atoi(string(id[i]))

		if isSecondDigit {
			doubledDigit := digit * 2
			if doubledDigit > 9 {
				doubledDigit = doubledDigit - 9
			}
			res.WriteString(strconv.Itoa(doubledDigit))
		} else {
			res.WriteString(strconv.Itoa(digit))
		}
	}

	return res.String()
}

func SumDigits(id string) int {
	total := 0
	for _, digit := range id {
		value, _ := strconv.Atoi(string(digit))
		total += value
	}
	return total
}
