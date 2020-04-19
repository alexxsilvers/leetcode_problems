package main

import (
	"fmt"
)

// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return
// the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1)
// extra memory.
// Clarification:
// Confused why the returned value is an integer but your answer is an array?
// Note that the input array is passed in by reference, which means modification to the input array will be known
// to the caller as well.
func main() {
	ex1 := []int{0, 0, 0, 1, 1, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(ex1), ex1) // l = 5 [0,1,2,3,4]

	ex2 := []int{1, 2, 3}
	fmt.Println(removeDuplicates(ex2), ex2) // l = 3 [1,2,3]

	ex3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 4, 4, 4}
	fmt.Println(removeDuplicates(ex3), ex3) // l = 5 [0,1,2,3,4]
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
