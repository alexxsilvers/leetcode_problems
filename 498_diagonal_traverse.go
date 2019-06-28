package main

import "log"

func main() {
	ex1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	log.Println(findDiagonalOrder(ex1)) // [1,2,4,7,5,3,6,8,9]

	ex2 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	log.Println(findDiagonalOrder(ex2)) // [1,2,3,5,4,6]

	ex3 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	log.Println(findDiagonalOrder(ex3)) // [1,2,5,6,3,4,7,8]
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	row, col, m, n := 0, 0, len(matrix), len(matrix[0])
	size := m * n
	res := make([]int, size)
	for pos := 0; pos < size; pos++ {
		res[pos] = matrix[row][col]

		if (row+col)&1 == 0 { // go up
			if col == n-1 {
				row++
			} else if row == 0 {
				col++
			} else {
				row--
				col++
			}
		} else { // go down
			if row == m-1 {
				col++
			} else if col == 0 {
				row++
			} else {
				row++
				col--
			}
		}
	}
	return res
}
