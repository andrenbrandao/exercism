package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n <= 0 {
		return 0, errors.New("n must be a positive integer")
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}

	return count, nil
}
