package romannumerals

import (
	"errors"
	"math"
	"strings"
)

/*

SOLUTION 1

We want to convert from normal numbers to romans.

1 -> I
2 -> II
3 -> III
4 -> IV
5 -> V
6 -> VI
...
10 -> X
50 -> L
100 -> C
500 -> D
1000 -> M

1990 -> MCMXC
90 = XC
900 = CM
1000 = M

Notice that we need to take the digits from right to left and form the roman
numerals.

1999 -> MCMXCIX

Algorithm:
- Take rightmost digit
- If it is not zero, convert it to roman numeral
- Iterate to the left digit and keep converting them
- Reverse the string at the end

To do this we will need a hashmap to convert the decimals to romans.

If digit is [2,3] we repeat the roman numeral that represents 1,
but if it is 4, we have to prefix the roman at position 5 with the roman at 1.

The digits 6, 7 and 8 also add a suffix of digit 1 to a roman numeral 5.
A digit 9 adds a prefix of digit 1 to a roman numeral of 10.

*/

var intToRoman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

// BenchmarkRomanNumerals-8          549159              3447 ns/op             384 B/op         48 allocs/op
func ToRomanNumeralSol1(input int) (string, error) {
	digitPos := 0
	sb := strings.Builder{}

	if input <= 0 || input > 3000 {
		return "", errors.New("can only convert integers from 1 to 3000")
	}

	for input != 0 {
		lastDigit := input % 10

		switch lastDigit {
		case 1, 2, 3:
			for i := 0; i < lastDigit; i++ {
				sb.WriteString(digitToRoman(1, digitPos))
			}
		case 4:
			sb.WriteString(digitToRoman(5, digitPos))
			sb.WriteString(digitToRoman(1, digitPos))
		case 5:
			sb.WriteString(digitToRoman(5, digitPos))
		case 6, 7, 8:
			for i := 0; i < lastDigit-5; i++ {
				sb.WriteString(digitToRoman(1, digitPos))
			}
			sb.WriteString(digitToRoman(5, digitPos))
		case 9:
			sb.WriteString(digitToRoman(10, digitPos))
			sb.WriteString(digitToRoman(1, digitPos))
		}

		input /= 10
		digitPos++
	}

	romanNumerals := reverseString(sb.String())
	return romanNumerals, nil
}

func digitToRoman(digit, digitPos int) string {
	return intToRoman[digit*int(math.Pow10(digitPos))]
}

func reverseString(s string) string {
	sb := strings.Builder{}
	n := len(s)
	sb.Grow(n)

	for i := n - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	return sb.String()
}

/*

SOLUTION 2

Use a greedy approach and take the maximum value we can from the input.
Keep repeating until we reach zero.

*/

type intRoman struct {
	val   int
	roman string
}

var romanNumerals = []intRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("can only convert integers from 1 to 3000")
	}

	sb := strings.Builder{}

	for input != 0 {
		for _, c := range romanNumerals {
			if input >= c.val {
				sb.WriteString(c.roman)
				input -= c.val
				break
			}
		}
	}

	return sb.String(), nil
}
