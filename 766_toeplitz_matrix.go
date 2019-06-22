package main

import "log"

// A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.
//
// Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
func main() {
	ex2 := [][]int{
		{1,2},
		{3,1},
		{4,3},
		{5,4},
	}
	log.Println(isToeplitzMatrix(ex2)) // true

	ex3 := [][]int{
		{1,2},
		{1,1},
	}
	log.Println(isToeplitzMatrix(ex3)) // true

	ex4 := [][]int{
		{1},
		{1},
		{1},
	}
	log.Println(isToeplitzMatrix(ex4)) // true

	ex1 := [][]int{
		{1,2,3},
		{4,1,2},
		{6,4,1},
	}
	log.Println(isToeplitzMatrix(ex1)) // true

	ex5 := [][]int{
		{1,2,3},
		{1,2,2},
	}
	log.Println(isToeplitzMatrix(ex5)) // false
}

// solution with O(m * n) runtime and O(m + n) space
// they will be good, if we need to store early results in some storage
func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		return true
	}

	store := make(map[int]int)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if _, ok := store[r-c]; !ok {
				store[r-c] = matrix[r][c]
			}
			if matrix[r][c] != store[r-c] {
				return false
			}
		}
	}

	return true
}

// solution with O(m * n) runtime and O(1) space
//func isToeplitzMatrix(matrix [][]int) bool {
//	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
//		return true
//	}
//
//	for r := 1; r < len(matrix); r++ {
//		for c := 1; c < len(matrix[0]); c++ {
//			if matrix[r][c] != matrix[r-1][c-1] {
//				return false
//			}
//		}
//	}
//
//	return true
//}
