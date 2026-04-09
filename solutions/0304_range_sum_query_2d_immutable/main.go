package main

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	prefix := make([][]int, rows+1)
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}

	for r := 1; r <= rows; r++ {
		for c := 1; c <= cols; c++ {
			prefix[r][c] = matrix[r-1][c-1] +
				prefix[r-1][c] +
				prefix[r][c-1] -
				prefix[r-1][c-1]
		}
	}

	return NumMatrix{prefix}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return nm.prefix[row2+1][col2+1] -
		nm.prefix[row1][col2+1] -
		nm.prefix[row2+1][col1] +
		nm.prefix[row1][col1]
}
