package matrix

import (
	"errors"
	"strconv"
	"strings"
)

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

func (m *Matrix) Cols() [][]int {
	height := len(*m)
	width := len((*m)[0])
	columns := make([][]int, width)

	for j := 0; j < width; j++ {
		columns[j] = make([]int, height)
		for i := 0; i < height; i++ {
			columns[j][i] = (*m)[i][j]
		}
	}

	return columns
}

func (m *Matrix) Rows() [][]int {
	height := len(*m)
	width := len((*m)[0])
	rows := make([][]int, height)

	for i := 0; i < height; i++ {
		rows[i] = make([]int, width)
		copy(rows[i], (*m)[i])
	}

	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	height := len(*m)
	width := len((*m)[0])

	if row < 0 || row >= height || col < 0 || col >= width {
		return false
	}

	(*m)[row][col] = val

	return true
}
