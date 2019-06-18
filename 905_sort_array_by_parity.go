package main

import "log"

// Given an array A of non-negative integers, return an array
// consisting of all the even elements of A, followed by all the odd elements of A.
//
// You may return any answer array that satisfies this condition.
func main() {
	log.Println(sortArrayByParity([]int{1,3,8,15,20})) // [8, 20, 1, 3, 15]
	log.Println(sortArrayByParity([]int{3,1,2,4})) // [2, 4, 3, 1]
}

func sortArrayByParity(A []int) []int {
	odd := make([]int, 0)
	even := make([]int, 0)
	for _, num := range A {
		if num & 1 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return append(even, odd...)
}
