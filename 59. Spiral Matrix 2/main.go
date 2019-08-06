package main

import "log"

// Given a positive integer n, generate a square matrix filled with elements from 1 to n in spiral order.
func main() {
	log.Println(generateMatrix(3))
	/*
		[1,2,3]
		[8,9,4]
		[7,6,5]
	*/

	log.Println(generateMatrix(1))
	/*
		[1]
	*/

	log.Println(generateMatrix(2))
	/*
		[1,2]
		[4,3]
	*/

	log.Println(generateMatrix(5))
	/*
		[1 ,2 ,3 ,4 ,5]
		[16,17,18,19,6]
		[15,24,25,20,7]
		[14,23,22,21,8]
		[13,12,11,10,9]
	*/
}

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	m := createEmptyMatrix(n)
	p, rs, re, cs, ce := 1, 0, n-1, 0, n-1
	for rs <= re && cs <= ce {
		for c := cs; c <= ce; c++ {
			m[rs][c] = p
			p++
		}
		for r := rs + 1; r <= re; r++ {
			m[r][ce] = p
			p++
		}

		if rs < re && cs < ce {
			for c := ce - 1; c > cs; c-- {
				m[re][c] = p
				p++
			}
			for r := re; r > rs; r-- {
				m[r][cs] = p
				p++
			}
		}

		rs++
		re--
		cs++
		ce--
	}
	return m
}

func createEmptyMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	return m
}
