package main

import "log"

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
func main() {
	ex1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	log.Println(spiralOrder(ex1)) // [1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]

	ex2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	log.Println(spiralOrder(ex2)) // [1,2,3,6,9,8,7,4,5]

	ex3 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	log.Println(spiralOrder(ex3)) // [1,2,3,4,8,12,11,10,9,5,6,7]

	ex4 := [][]int{
		{1},
	}
	log.Println(spiralOrder(ex4)) // [1]
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0, len(matrix)*len(matrix[0]))
	rs, re, cs, ce := 0, len(matrix)-1, 0, len(matrix[0])-1
	for rs <= re && cs <= ce {
		for c := cs; c <= ce; c++ {
			res = append(res, matrix[rs][c])
		}
		for r := rs + 1; r <= re; r++ {
			res = append(res, matrix[r][ce])
		}
		if rs < re && cs < ce {
			for c := ce - 1; c > cs; c-- {
				res = append(res, matrix[re][c])
			}
			for r := re; r > rs; r-- {
				res = append(res, matrix[r][cs])
			}
		}

		rs++
		cs++
		re--
		ce--
	}
	return res
}
