package main

import "log"

// Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
func main() {
	log.Println(missingNumber([]int{3, 0, 1}))                   // 2
	log.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // 8
}

// Runtime - O(n)
// Space - O(1)
// Use Gauss formula sum = (1 + size(last number)) * size / 2
func missingNumber(nums []int) int {
	s := (len(nums) + 1) * len(nums) / 2
	for _, n := range nums {
		s -= n
	}
	return s
}
