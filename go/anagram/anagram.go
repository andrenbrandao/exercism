package anagram

import (
	"sort"
	"strings"
	"unicode"
)

/*

We can count the letters of the subject.
For every candidate, count the letters
If it matches the count of the subject, we have
an anagram.

Time Complexity: O(n*s)
Longest string length as "s"
Number of candidates as "n"

Space Complexity: O(n) if
we create a new hashmap for every candidate or
we can decrease it to O(1) if we always use the same hashmap.


-- Can we do another way without using the hashmaps?

We can sort the candidate letters, sort the subject
and compare the two strings.

Time Complexity: O(n*slogs) to sort the strings
Space Complexity: O(1) because we are not using any
extra memory besides sorting.

*/

// ** Hashmap Method **
// BenchmarkDetectAnagrams-8          84477             12887 ns/op            1040 B/op         47 allocs/op
func Detect(subject string, candidates []string) []string {
	lowercaseSubject := strings.ToLower(subject)
	letterCount := map[rune]int{}
	result := []string{}

	for _, letter := range subject {
		lowercaseLetter := unicode.ToLower(letter)
		letterCount[lowercaseLetter] += 1
	}

	for _, candidate := range candidates {
		lowercaseCandidate := strings.ToLower(candidate)
		candidateLetterCount := map[rune]int{}

		if len(candidate) != len(subject) || lowercaseSubject == lowercaseCandidate {
			continue
		}

		for _, letter := range candidate {
			lowercaseLetter := unicode.ToLower(letter)
			candidateLetterCount[lowercaseLetter] += 1
		}

		if hasSameLetterCount(letterCount, candidateLetterCount) {
			result = append(result, candidate)
		}
	}

	return result
}

func hasSameLetterCount(letterCount map[rune]int, candidateLetterCount map[rune]int) bool {
	for key, val := range candidateLetterCount {
		if letterCount[key] != val {
			return false
		}
	}
	return true
}

// ** Sorting Method **
// BenchmarkDetectAnagrams-8          39752             30205 ns/op            6432 B/op           109 allocs/op
func DetectBySort(subject string, candidates []string) []string {
	lowercaseSubject := strings.ToLower(subject)
	sortedSubject := strings.Split(lowercaseSubject, "")
	sort.Strings(sortedSubject)

	result := []string{}

	for _, candidate := range candidates {
		lowercaseCandidate := strings.ToLower(candidate)
		sortedCandidate := strings.Split(lowercaseCandidate, "")
		sort.Strings(sortedCandidate)

		if len(sortedSubject) != len(sortedCandidate) || lowercaseSubject == lowercaseCandidate {
			continue
		}

		if isSameStringSlice(sortedSubject, sortedCandidate) {
			result = append(result, candidate)
		}

	}

	return result
}

func isSameStringSlice(sortedSubject []string, sortedCandidate []string) bool {
	for i := range sortedSubject {
		if sortedSubject[i] != sortedCandidate[i] {
			return false
		}
	}

	return true
}
