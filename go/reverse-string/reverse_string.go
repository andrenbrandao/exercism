package reverse

import "strings"

func Reverse(input string) string {
	var strBuilder strings.Builder
	runes := []rune(input)
	n := len(runes)

	for i := n - 1; i >= 0; i-- {
		strBuilder.WriteRune(runes[i])
	}

	return strBuilder.String()
}
