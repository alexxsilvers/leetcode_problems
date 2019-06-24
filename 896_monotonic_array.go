package main

import "log"

// An array is monotonic if it is either monotone increasing or monotone decreasing.
//
// An array A is monotone increasing if for all i <= j, A[i] <= A[j].
// An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
//
// Return true if and only if the given array A is monotonic.
func main() {
	log.Println(isMonotonic([]int{1, 2, 2, 3}))    // true
	log.Println(isMonotonic([]int{6, 5, 4, 4}))    // true
	log.Println(isMonotonic([]int{1, 3, 2}))       // false
	log.Println(isMonotonic([]int{1, 2, 4, 5}))    // true
	log.Println(isMonotonic([]int{1, 1, 1}))       // true
	log.Println(isMonotonic([]int{10, 10, 10, 9})) // true
	log.Println(isMonotonic([]int{1}))             // true
}

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	// check first and last elements to determine monotone direction
	if A[0] >= A[len(A)-1] { // it is monotone decreasing
		for i := 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
	} else { // it is monotone increasing
		for i := 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
	}
	return true
}
