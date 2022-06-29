package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*

** Brute Force **

- Create a map[string]bool to store the usedNames
- When we create or reset a robot, we keep trying to
generate names and check if they were already used

The problem with this approach is that we can have many collisions
and the for loop will keep iterating for a long time.

** How can we do better? **

- Precompute all possible robot names
- Shuffle them
- Take one from the list and remove it
- If the list is empty we know we have used all of them

*/

type Robot struct {
	name string
}

var availableNames []string = generateAllNames()

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(availableNames) == 0 {
			return "", errors.New("All names have already been used")
		}

		r.Reset()
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	newName := availableNames[len(availableNames)-1]
	availableNames = availableNames[:len(availableNames)-1]
	r.name = newName
}

/*

First generated name: AA000
Last generated name: ZZ999

Since we can have 26 * 26 * 10 * 10 * 10 = 676000 names,
and they take 5 bytes each, we would use a total of aprox. 3.22 MB of memory.

To generate all possible names:
- Generate strings AA to ZZ
- Generate numbers 000 to 999
- Concatenate the two

*/

func generateAllNames() []string {
	rand.Seed(time.Now().UnixNano())
	maxNames := 65 * 65 * 10 * 10 * 10
	result := make([]string, 0, maxNames)

	letterPairs := make([]string, 0, 676)
	numberTriplets := make([]string, 0, 1000)

	// Generate pairs of letters from AA to ZZ
	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			letterPairs = append(letterPairs, fmt.Sprintf("%c%c", i, j))
		}
	}

	// Generate triplets of numbers from 000 to 999
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				numberTriplets = append(numberTriplets, fmt.Sprintf("%c%c%c", i, j, k))
			}
		}
	}

	// Join them
	for _, pair := range letterPairs {
		for _, triplet := range numberTriplets {
			result = append(result, pair+triplet)
		}
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

// ** BRUTE FORCE **

// type Robot struct {
// 	name string
// }

// var usedNames = map[string]bool{}

// func (r *Robot) Name() (string, error) {
// 	if r.name == "" {
// 		r.Reset()
// 	}

// 	return r.name, nil
// }

// func (r *Robot) Reset() {
// 	newName := r.generateName()

// 	// Keep retrying won't scale
// 	for usedNames[newName] {
// 		newName = r.generateName()
// 	}

// 	usedNames[newName] = true
// 	r.name = newName
// }

// func (r *Robot) generateName() string {
// 	rand.Seed(time.Now().UnixNano())
// 	firstLetterInt := rand.Int31n(26) + 65
// 	secondLetterInt := rand.Int31n(26) + 65
// 	thirdDigit := rand.Int31n(10)
// 	fourthDigit := rand.Int31n(10)
// 	fifthDigit := rand.Int31n(10)

// 	return fmt.Sprintf("%c%c%d%d%d",
// 		firstLetterInt,
// 		secondLetterInt,
// 		thirdDigit,
// 		fourthDigit,
// 		fifthDigit)
// }
