// Package wordcount provides methods to count words
// in phrases.
package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

/*

"That's the password: 'PASSWORD 123'!", cried the Special Agent.\nSo I fled.

Algorithm:
- Iterate through the string
- If we find a letter or number, start adding it to an array of runes
- When we find a delimiter, we store this word in our count
- If we find an apostrophe, we look ahead and if it is a valid character,
add the two characters to the current string
- In the end, if we have a string left in our array, we add it to the count

Time Complexity: O(n)
Space Complexity: O(n)

*/

type Frequency map[string]int

// WordCountRegex counts the words in a phrase
// It returns a map of strings with their frequencies.
// BenchmarkWordCount
// BenchmarkWordCount-8       15409             82745 ns/op           26629 B/op        344 allocs/op.
func WordCountRegex(phrase string) Frequency {
	freq := Frequency{}
	re := regexp.MustCompile(`\b[\w']+\b`)

	for _, str := range re.FindAllString(phrase, -1) {
		freq[strings.ToLower(str)]++
	}

	return freq
}

// WordCount counts the words in a phrase
// It returns a map of strings with their frequencies.
// BenchmarkWordCount
// BenchmarkWordCount-8      144775              9465 ns/op            3504 B/op         76 allocs/op.
func WordCount(phrase string) Frequency {
	freq := Frequency{}
	sb := strings.Builder{}

	for i, char := range phrase {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || isValidApostrophe(i, char, phrase, sb) {
			sb.WriteRune(unicode.ToLower(char))
		} else if sb.Len() > 0 {
			freq[sb.String()]++
			sb.Reset()
		}
	}

	if sb.Len() > 0 {
		freq[sb.String()]++
	}

	return freq
}

func isValidApostrophe(pos int, char rune, phrase string, sb strings.Builder) bool {
	return char == '\'' && pos+1 < len(phrase) && unicode.IsLetter(rune(phrase[pos+1])) && sb.Len() > 0
}
