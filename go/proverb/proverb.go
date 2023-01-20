// Package proverb has methods to generate a proverb
package proverb

import (
	"fmt"
)

// Proverb creates a proverb based on strings given to it
func Proverb(rhyme []string) []string {
	result := []string{}

	if len(rhyme) == 0 {
		return result
	}

	for i := 1; i < len(rhyme); i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}

	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return result
}
