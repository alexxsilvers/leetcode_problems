package main

import "log"

// Given a matrix A, return the transpose of A.
//
// The transpose of a matrix is the matrix flipped over it's main diagonal,
// switching the row and column indices of the matrix.
func main() {
	// [1, 2, 3]
	// [4, 5, 6]
	// [7, 8, 9]
	log.Println(transpose([][]int{{1,2,3}, {4,5,6}, {7,8,9}}))
	// [1, 4, 7]
	// [2, 5, 8]
	// [3, 6, 9]

	// [1, 2, 3]
	// [4, 5, 6]
	log.Println(transpose([][]int{{1,2,3}, {4,5,6}}))
	// [1, 4]
	// [2, 5]
	// [3, 6]

	// [1, 2, 3, 4]
	// [5, 6, 7, 8]
	// [9, 10, 11, 12]
	// [13, 14, 15, 16]
	log.Println(transpose([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}}))
	// [1, 5, 9,  13]
	// [2, 6, 10, 14]
	// [3, 7, 11, 15]
	// [4, 8, 12, 16]

	// [1, 2, 3, 4]
	// [5, 6, 7, 8]
	// [9, 10, 11, 12]
	log.Println(transpose([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
	// [1, 5, 9]
	// [2, 6, 10]
	// [3, 7, 11]
	// [4, 8, 12]
}

func transpose(A [][]int) [][]int {
	rowSize := len(A)
	colSize := len(A[0])
	result := make([][]int, colSize)

	for i := 0; i < colSize; i++ {
		result[i] = make([]int, rowSize)
	}

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			result[j][i] = A[i][j]
		}
	}

	return result
}