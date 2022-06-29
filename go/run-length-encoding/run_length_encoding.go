package encode

import (
	"fmt"
	"strings"
)

/*

Algorithm:
- Read letters in pairs
- Count the number of occurrences
- When the pair of letters is different, store the count with the last letter seen

Time Complexity: O(n)
Space Complexity: O(n)

BenchmarkRunLengthEncode
BenchmarkRunLengthEncode-8       1382593               913.8 ns/op            72 B/op          7 allocs/op

*/

func RunLengthEncode(input string) string {
	count := 1
	sb := strings.Builder{}

	if len(input) == 0 {
		return input
	}

	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			count++
			continue
		}

		sb.WriteString(encodeSymbol(count, input[i-1]))
		count = 1
	}

	sb.WriteString(encodeSymbol(count, input[len(input)-1]))

	return sb.String()
}

func encodeSymbol(count int, symbol byte) string {
	if count == 1 {
		return string(symbol)
	}

	return fmt.Sprintf("%d%c", count, symbol)
}

/*

Algorithm:
- Read character by character
- If we find a number, we add it to our current count (start it with 0, then multiply by 10 and add the number)
- If the character is a letter
	- Add the letter count times to the string
	- But, if count == 0, we only add it once
	- Set count = 0

BenchmarkRunLengthDecode
BenchmarkRunLengthDecode-8       1281919              1009 ns/op             200 B/op         11 allocs/op

*/

func RunLengthDecode(input string) string {
	count := 0
	sb := strings.Builder{}

	for i := 0; i < len(input); i++ {
		symbol := input[i]

		if symbol >= '0' && symbol <= '9' {
			count = count*10 + int(symbol-'0')
			continue
		}

		if count == 0 {
			sb.WriteByte(input[i])
			continue
		}

		for j := 0; j < count; j++ {
			sb.WriteByte(input[i])
		}

		count = 0
	}

	return sb.String()
}
