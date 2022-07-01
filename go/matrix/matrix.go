package matrix

import (
	"errors"
	"strconv"
	"strings"
)

/*

Tips from andrerfcsantos:

About the methods on Matrix - you can have value receivers(m Matrix) instead of pointer receivers (m Matrix).
New returns a *Matrix, but when the tests use the . notation to call methods on the result, if Go sees that there
are no pointer receivers for the type, it will see if there are value receivers and if so, it will dereference
the pointer automatically and call the value receiver.

Usually we don't use value receivers and use pointer receivers instead. This is because if we have a value
receiver, Go creates a copy of the instance we are calling the method on, while if we use pointer receivers
the copy is not created.

However, slices, maps and channels are sort of a special case. Internally in Go, they are represented with
a pointer to the data so even when you have a value receiver (or pass a slice as argument to a function),
Go actually passes a copy of the pointer, and not of the data like other types. So you can see slices, mapsa
nd channel values as "being already pointers", even if we don't use the pointer notation on them.

So, using a pointer receiver is not necessary, and you can avoid the weird dereferences like (*m) and just use m.

*/

// ** Benchmarks **
// BenchmarkNew
// BenchmarkNew-8            723206              2862 ns/op            1056 B/op         25 allocs/op
// BenchmarkRows
// BenchmarkRows-8          3867987               283.9 ns/op           192 B/op          5 allocs/op
// BenchmarkCols
// BenchmarkCols-8          2939164               404.5 ns/op           288 B/op          6 allocs/op

type Matrix [][]int

func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")
	matrix := Matrix{}

	for i, line := range lines {
		line := strings.TrimSpace(line)
		numbers := strings.Split(line, " ")

		matrix = append(matrix, []int{})
		matrix[i] = make([]int, 0)

		for _, number := range numbers {
			numberInteger, err := strconv.Atoi(number)

			if err != nil {
				return nil, err
			}

			matrix[i] = append(matrix[i], numberInteger)
		}

		// Check for rows of different sizes
		if i > 0 && len(matrix[i]) != len(matrix[i-1]) {
			return nil, errors.New("Matrix must have the same length for all rows")
		}
	}

	return &matrix, nil
}

func (m Matrix) Cols() [][]int {
	height := len(m)
	width := len(m[0])
	columns := make([][]int, width)

	for j := 0; j < width; j++ {
		columns[j] = make([]int, height)
		for i := 0; i < height; i++ {
			columns[j][i] = m[i][j]
		}
	}

	return columns
}

func (m Matrix) Rows() [][]int {
	height := len(m)
	width := len(m[0])
	rows := make([][]int, height)

	for i := 0; i < height; i++ {
		rows[i] = make([]int, width)
		copy(rows[i], m[i])
	}

	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	height := len(m)
	width := len(m[0])

	if row < 0 || row >= height || col < 0 || col >= width {
		return false
	}

	m[row][col] = val

	return true
}
