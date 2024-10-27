package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix holds data for matrices.
type Matrix struct {
	rows [][]int
	cols [][]int
}

// New is a factory function for matrices.
func New(str string) (*Matrix, error) {
	var err error
	var m Matrix

	rowLines := strings.Split(str, "\n")

	// Build the structures needed to
	// parse matrix data.
	numRows := len(rowLines)
	numCols := len(strings.Split(rowLines[0], " "))

	rows := make([][]int, numRows)
	for rn := range rows {
		rows[rn] = make([]int, numCols)
	}
	cols := make([][]int, numCols)
	for cn := range cols {
		cols[cn] = make([]int, numRows)
	}

	for rn, rl := range rowLines {
		numStrs := strings.Split(strings.TrimSpace(rl), " ")

		// Check the number of columns and perform conversions.
		if len(numStrs) != numCols {
			return &m, errors.New("rows uneven")
		}

		nums := make([]int, len(numStrs))
		for cn, numStr := range numStrs {
			var num int
			if num, err = strconv.Atoi(numStr); nil != err {
				return &m, err
			}
			nums[cn] = num
		}

		// Fill one entire row.
		rows[rn] = nums

		// Fill one element in every column.
		for cn := 0; cn < numCols; cn++ {
			cols[cn][rn] = nums[cn]
		}
	}

	m.rows = rows
	m.cols = cols

	return &m, nil
}

// Set sets a matrix element after checking that the
// values are within the bounds of the matrix.
func (m *Matrix) Set(row int, col int, val int) bool {
	if row < 0 || row > (len(m.rows)-1) ||
		col < 0 || col > (len(m.cols)-1) {
		return false
	}

	m.rows[row][col] = val
	m.cols[col][row] = val

	return true
}

// Rows returns an independent copy of the
// row data for a matrix.
func (m *Matrix) Rows() [][]int {
	rowsCopy := make([][]int, len(m.rows))

	for rn, r := range m.rows {
		rowsCopy[rn] = make([]int, len(r))
		copy(rowsCopy[rn], r)
	}

	return rowsCopy
}

// Cols returns an independent copy of the
// column data for a matrix.
func (m *Matrix) Cols() [][]int {
	colsCopy := make([][]int, len(m.cols))

	for cn, c := range m.cols {
		colsCopy[cn] = make([]int, len(c))
		copy(colsCopy[cn], c)
	}

	return colsCopy
}
