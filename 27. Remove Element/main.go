package main

import "log"

// Given an array nums and a value val, remove all instances of that value in-place and return the new length.
//
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
// The order of elements can be changed. It doesn't matter what you leave beyond the new length.
func main() {
	nums1 := []int{3, 2, 2, 4}
	log.Println(removeElement(nums1, 2), len(nums1)) // len=2, [3,4]
}

func removeElement(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	nums = nums[:n]
	return n
}
