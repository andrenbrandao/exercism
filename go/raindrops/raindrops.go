package raindrops

import "strconv"

func Convert(n int) (result string) {
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		result = strconv.Itoa(n)
	}

	if n%3 == 0 {
		result += "Pling"
	}

	if n%5 == 0 {
		result += "Plang"
	}

	if n%7 == 0 {
		result += "Plong"
	}

	return
}
