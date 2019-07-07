package main

import (
	"log"
)

// Given an array of integers A sorted in non-decreasing order,
// return an array of the squares of each number, also in sorted non-decreasing order.
func main() {
	log.Println(sortedSquares([]int{-10, -7, 1, 0, 4, 5, 9})) // [0, 1, 16, 25, 49, 81, 100]
	log.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))       // [4, 9, 9, 49, 121]
}

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	left := 0
	right := len(A) - 1
	index := len(A) - 1
	for left <= right {
		l, r := A[left]*A[left], A[right]*A[right]
		if l < r {
			result[index] = r
			right--
		} else {
			result[index] = l
			left++
		}
		index--
	}
	return result
}
