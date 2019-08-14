package main

import "log"

// Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
// Note that the row index starts from 0.
func main() {
	log.Println(getRow(0))
	log.Println(getRow(1))
	log.Println(getRow(2))
	log.Println(getRow(3))
	log.Println(getRow(4))
	log.Println(getRow(5))
	log.Println(getRow(6))
}

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	res[0] = 1
	prev := 1
	for i := 0; i < len(res)-1; i++ {
		prev = prev * (rowIndex - i) / (i + 1)
		res[i+1] = prev
	}

	return res
}
