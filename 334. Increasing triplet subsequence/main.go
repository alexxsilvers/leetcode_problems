package main

import (
	"log"
	"math"
)

// Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.
//
// Formally the function should:
//
// Return true if there exists i, j, k
// such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
// Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.
func main() {
	log.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))      // true
	log.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))      // false
	log.Println(increasingTriplet([]int{20, 15, 17, 13, 19})) // true
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	smallest, secondSmallest := nums[0], math.MaxInt64

	for i := 1; i < len(nums); i++ {
		switch num := nums[i]; {
		case num < smallest:
			smallest = num
		case num > smallest && num < secondSmallest:
			secondSmallest = num
		case num > secondSmallest:
			return true
		}
	}

	return false
}
