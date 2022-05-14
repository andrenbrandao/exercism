package grains

import "errors"

// OPTIMIZING FOR SPEED

// BenchmarkSquare-8       83434033                12.66 ns/op            0 B/op          0 allocs/op
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Square number should be between 1 and 64")
	}

	return 1 << (number - 1), nil
}

// BenchmarkTotal-8        1000000000               0.3803 ns/op          0 B/op          0 allocs/op
func Total() uint64 {
	return (1 << 64) - 1
}

// First Solution Below
// Done for Readability

// BenchmarkSquare-8        2557620               426.9 ns/op            48 B/op          3 allocs/op
// func Square(number int) (uint64, error) {
// 	if number < 1 || number > 64 {
// 		return 0, errors.New("Square number should be between 1 and 64")
// 	}

// 	return uint64(math.Pow(2.0, float64(number)-1)), nil
// }

// BenchmarkTotal-8          652273              1811 ns/op               0 B/op          0 allocs/op
// func Total() uint64 {
// 	var total uint64

// 	for i := 1; i <= 64; i++ {
// 		amount, err := Square(i)

// 		if err != nil {
// 			panic(err)
// 		}

// 		total += amount
// 	}

// 	return total
// }
