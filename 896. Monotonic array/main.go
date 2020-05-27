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
	if len(A) <= 1 {
		return true
	}

	if A[0] < A[len(A)-1] { // monotonic increasing
		for i := 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
	} else { // monotonic decreasing
		for i := 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
	}

	return true
}

func isMonotonicSimple(A []int) bool {
	monotonicDecreasing := true
	monotonicIncreasing := true
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			monotonicIncreasing = false
		}
		if A[i] > A[i+1] {
			monotonicDecreasing = false
		}
	}

	return monotonicDecreasing || monotonicIncreasing
}
