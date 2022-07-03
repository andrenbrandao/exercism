// Package etl converts scrablle scores in a legacy format to a new format
package etl

import "strings"

/*

We are going to receive a hashmap containing the scores and the following
letters per score.

{
	1: ['A', 'B' ...],
	2: ['D', 'G', ...]
}

We want to convert it to a hashmap from letters (lowercase) to scores.

{
	'a': 1,
	'b': 1,
	'd': 2,
	'g': 2,
}

To do this, we will need to iterate over all the keys in the original hashmap
and the letters.
Then, lowercase the letters and add them as a key in the new hashmap with the
corresponding score.

Time Complexity: O(n) being n the letters available in the input.
Space Complexity: O(n) if we consider the new hashmap created.

*/

// Transform takes a legacy format of scrabble scores and transforms it
// into the new format
func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}

	for score, letters := range in {
		for _, letter := range letters {
			res[strings.ToLower(letter)] = score
		}
	}

	return res
}
