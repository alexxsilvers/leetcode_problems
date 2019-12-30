package main

import (
	"fmt"
)

// Given n and m which are the dimensions of a matrix initialized by zeros and given an array indices where
// indices[i] = [ri, ci]. For each pair of [ri, ci] you have to increment all cells in row ri and column ci by 1.
//
// Return the number of cells with odd values in the matrix after applying the increment to all indices.
func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}})) // 6
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}})) // 0
}

// Runtime O(n), space - O(n)
func oddCells(n int, m int, indices [][]int) int {
	cells := make([]int, n*m)
	for _, ind := range indices {
		row := ind[0]
		col := ind[1]

		// fill rows
		for i := row * m; i < row*m+m; i++ {
			cells[i]++
		}

		// fill cols
		for i := 0; i < n; i++ {
			cells[i*m+col]++
		}
	}

	odd := 0
	for _, n := range cells {
		if n&1 == 1 {
			odd++
		}
	}

	return odd
}
