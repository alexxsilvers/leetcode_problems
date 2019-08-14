package main

import "log"

// Given an integer array nums, find the contiguous subarray within an array (containing at least one number)
// which has the largest product.
func main() {
	log.Println(maxProduct([]int{2, 3, -2, 4})) // 6
	log.Println(maxProduct([]int{-2, 0, -1}))   // 0
	log.Println(maxProduct([]int{-2, -5, 1}))   // 10
}

// Runtime - O(n)
// Space - O(1)
func maxProduct(nums []int) int {
	imax, imin, totalMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := nums[i] * imax
		tempMin := nums[i] * imin

		imax = max(max(tempMax, tempMin), nums[i])
		imin = min(min(tempMax, tempMin), nums[i])

		totalMax = max(totalMax, imax)
	}
	return totalMax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
