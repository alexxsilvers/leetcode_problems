package main

import (
	"log"
)

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
func main() {
	log.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	// [4,-1,2,1]

	log.Println(maxSubArray([]int{-7, -5, -5, 5, -2})) // 5

	log.Println(maxSubArray([]int{-1})) // -1

	log.Println(maxSubArray([]int{-2, -1})) // -1
}

// Runtime O(n)
// Space O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	m, mEnd := nums[0], nums[0]
	for _, n := range nums[1:] {
		mEnd = max(mEnd+n, n)
		m = max(m, mEnd)
	}
	return m
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
