package main

import "log"

// Given an array, rotate the array to the right by k steps, where k is non-negative.
func main() {
	ex1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(ex1, 3)
	log.Println(ex1) // [5,6,7,1,2,3,4]

	ex2 := []int{-1, -100, 3, 99}
	rotate(ex2, 2)
	log.Println(ex2) // [3, 99, -1, -100]
}

// Time complexity - O(n), space complexity - O(1)
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, s int, e int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}

// Time complexity - O(n), space complexity - O(n)
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	offset := k % len(nums)
	for i := 0; i < len(nums); i++ {
		n := i + offset
		if n >= len(nums) {
			n -= len(nums)
		}
		newNums[n] = nums[i]
	}

	for i, n := range newNums {
		nums[i] = n
	}
}
