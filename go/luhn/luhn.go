package luhn

// Performant Approach
// (Based on this solution: https://exercism.org/tracks/go/exercises/luhn/solutions/bobahop)

// Check the other one below for more readability
// BenchmarkValid-8        39514232                28.90 ns/op            0 B/op          0 allocs/op
func Valid(id string) bool {
	reversePos := 0
	sumOfDigits := 0

	for i := len(id) - 1; i >= 0; i-- {
		char := id[i]

		if char == ' ' {
			continue
		}

		if char < '0' || char > '9' {
			return false
		}

		digit := int(char - '0')

		if reversePos%2 == 0 {
			sumOfDigits += digit
		} else {
			doubledDigit := digit << 1

			if doubledDigit > 9 {
				sumOfDigits += doubledDigit - 9
			} else {
				sumOfDigits += doubledDigit
			}
		}

		reversePos += 1
	}

	if reversePos <= 1 {
		return false
	}

	return sumOfDigits%10 == 0
}

// Helper Functions Approach
// BenchmarkValid-8         2907538               385.2 ns/op            40 B/op          3 allocs/op

// func Valid(id string) bool {
// 	strippedNumbers := StripSpaces(id)

// 	if len(strippedNumbers) <= 1 {
// 		return false
// 	}

// 	if HasNonDigits(strippedNumbers) {
// 		return false
// 	}

// 	modifiedSecondDigits := DoubleSecondDigits(strippedNumbers)
// 	sumOfDigits := SumDigits(modifiedSecondDigits)
// 	return sumOfDigits%10 == 0
// }

// func StripSpaces(id string) string {
// 	strippedNumbers := strings.ReplaceAll(id, " ", "")
// 	return strippedNumbers
// }

// func HasNonDigits(id string) bool {
// 	for _, char := range id {
// 		if !unicode.IsDigit(char) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func DoubleSecondDigits(id string) string {
// 	var res strings.Builder

// 	for i := len(id) - 1; i >= 0; i-- {
// 		isSecondDigit := (len(id)-1-i)%2 == 1
// 		digit := int(id[i] - '0')

// 		if isSecondDigit {
// 			doubledDigit := digit * 2
// 			if doubledDigit > 9 {
// 				doubledDigit = doubledDigit - 9
// 			}
// 			res.WriteString(strconv.Itoa(doubledDigit))
// 		} else {
// 			res.WriteString(strconv.Itoa(digit))
// 		}
// 	}

// 	return res.String()
// }

// func SumDigits(id string) int {
// 	total := 0
// 	for _, digit := range id {
// 		total += int(digit - '0')
// 	}
// 	return total
// }
