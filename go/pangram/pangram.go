package pangram

import "unicode"

/*

Algorithm:
- Create an array of bool to represent the 26 characters
- Iterate over the string reading character by character
- If it is a character between a-z (transform to lowercase), we
mark it as true in an array
- In the end, check if all values in array are true

Time Complexity: O(n)
Space Complexity: O(1) because we only use an array of size 26
as extra memory

Notice that I coded up another solution below but it was slower than
the solution with the boolean array.
*/

// BenchmarkPangram
// BenchmarkPangram-8       1000000              1063 ns/op               0 B/op          0 allocs/op

func IsPangram(input string) bool {
	seen := make([]bool, 26)

	for _, char := range input {
		lowercase := unicode.ToLower(char)

		if lowercase >= 'a' && lowercase <= 'z' {
			seen[lowercase-'a'] = true
		}
	}

	for _, val := range seen {
		if !val {
			return false
		}
	}

	return true
}

// BenchmarkPangram
// BenchmarkPangram-8        503050              2012 ns/op              96 B/op          2 allocs/op
// var alphabet = "abcdefghijklmnopqrstuvwxyz"

// func IsPangram(input string) bool {
// 	s := strings.ToLower(input)

// 	for _, char := range alphabet {
// 		if !strings.ContainsRune(s, char) {
// 			return false
// 		}
// 	}

// 	return true
// }
