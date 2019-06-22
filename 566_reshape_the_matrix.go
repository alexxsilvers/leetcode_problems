package main

import "log"

// In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one
// with different size but keep its original data.
//
// You're given a matrix represented by a two-dimensional array, and two positive integers r and c
// representing the row number and column number of the wanted reshaped matrix, respectively.
//
// The reshaped matrix need to be filled with all the elements of the original matrix in the same
// row-traversing order as they were.
//
// If the 'reshape' operation with given parameters is possible and legal, output the new reshaped
// matrix; Otherwise, output the original matrix.
func main() {
	ex1 := [][]int{
		{1,2},
		{3,4},
	}
	log.Println(matrixReshape(ex1, 1, 4))
	/*
	[1,2,3,4]
	 */

	ex2 := [][]int{
		{1,2},
		{3,4},
	}
	log.Println(matrixReshape(ex2, 2,4))
	/*
	[1,2]
	[3,4]
	 */

	ex3 := [][]int{
		{1,2},
		{3,4},
		{5,6},
	}
	log.Println(matrixReshape(ex3, 2, 3))
	/*
	[1,2,3]
	[4,5,6]
	 */

	ex4 := [][]int{
		{1},
		{1},
		{1},
	}
	log.Println(matrixReshape(ex4, 1, 3))
	/*
	[1,1,1]
	 */
}

// Runtime O(m * n), space O(r * c) == O(m * n)
//func matrixReshape(nums [][]int, r int, c int) [][]int {
//	if len(nums) == 1 && len(nums[0]) <= 1 {
//		return nums
//	}
//
//	if r*c != len(nums) * len(nums[0]) {
//		return nums
//	}
//
//	result := make([][]int, r)
//	for i := range result {
//		result[i] = make([]int, c)
//	}
//
//	row := 0
//	col := 0
//	for _, numsRow := range nums {
//		for _, n := range numsRow {
//			if col == c {
//				row++
//				col = 0
//			}
//			result[row][col] = n
//			col++
//		}
//	}
//	return result
//}

// The idea behind this approach is as follows. Do you know how a 2-D array is stored in the main
// memory(which is 1-D in nature)? It is internally represented as a 1-D array only.
// The element nums[i][j] of nums array is represented in the form of a one
// dimensional array by using the index in the form: nums[n*i + j], where n is
// the number of columns in the given matrix. Looking at the same in the reverse order,
// while putting the elements in the elements in the resultant matrix, we can make use of a
// count variable which gets incremented for every element traversed as if we are putting
// the elements in a 1-D resultant array. But, to convert the count back into 2-D matrix
// indices with a column count of c, we can obtain the indices as
// res[count/c][count\%c] where [count/c] is the row number
// and count%c is the column number. Thus, we can save the extra checking required at each step.
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 1 && len(nums[0]) <= 1 {
		return nums
	}
	if r * c != len(nums) * len(nums[0]) {
		return nums
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	count := 0
	for _, row := range nums {
		for _, num := range row {
			result[count/c][count%c] = num
			count++
		}
	}
	return result
}