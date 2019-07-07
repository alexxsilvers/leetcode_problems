package main

import "log"

// Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.
//
// Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.
//
// You may return any answer array that satisfies this condition.
func main() {
	log.Println(sortArrayByParityII([]int{2, 8, 10, 7, 3, 15, 22, 17}))
	log.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	evenI, oddI := 0, 1
	for evenI <= len(A)-2 && oddI <= len(A)-1 {
		if A[evenI]&1 == 1 && A[oddI]&1 == 0 {
			A[evenI], A[oddI] = A[oddI], A[evenI]
		}
		if A[evenI]&1 == 0 {
			evenI += 2
		}
		if A[oddI]&1 == 1 {
			oddI += 2
		}
	}
	return A
}
