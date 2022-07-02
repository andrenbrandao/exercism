package brackets

import (
	"strings"
)

/*

Cases:

[
)

()

(]

]()

Algorithm:
- Opening bracket: add to stack
- Closing bracket:
	- If stack is empty, return false
	- Otherwise, check if last element of stack matches,
	if yes, pop that last one
	else, return false
- In the end, if the stack is empty return true

Time Complexity: O(n)
Space Complexity: O(n)

Edge cases:
- Ignore any other type of characters
*/

var matchingBrackets = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func Bracket(input string) bool {
	brackets := "([{}])"
	stack := make([]rune, 0)

	for _, char := range input {
		if !strings.ContainsRune(brackets, char) {
			continue
		}

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
