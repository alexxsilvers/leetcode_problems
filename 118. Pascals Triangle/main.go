package main

import "log"

// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
func main() {
	log.Println(generate(1))
	/*
		1
	*/

	log.Println(generate(3))
	/*
		1
		11
		121
	*/

	log.Println(generate(5))
	/*
		1
		11
		121
		1331
		14541
	*/
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows <= 2 {
		return emptyPT(numRows)
	}

	pt := emptyPT(numRows)
	for row := 2; row < numRows; row++ {
		prevRow := pt[row-1]
		for i := 1; i < row; i++ {
			pt[row][i] = prevRow[i-1] + prevRow[i]
		}
	}

	return pt
}

func emptyPT(n int) [][]int {
	pt := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		pt[i] = row
	}
	return pt
}
