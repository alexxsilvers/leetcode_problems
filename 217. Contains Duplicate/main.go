package main

import (
	"log"
	"sort"
)

// Given an array of integers, find if the array contains any duplicates.
//
// Your function should return true if any value appears at least twice in
// the array, and it should return false if every element is distinct.
func main() {
	log.Println(containsDuplicateWithSpace([]int{3, 1, 2, 3})) // true
	log.Println(containsDuplicateWithoutSpace([]int{3, 1, 2, 3})) // true
	log.Println(containsDuplicateWithSpace([]int{3, 1, 2, 4})) // false
	log.Println(containsDuplicateWithoutSpace([]int{3, 1, 2, 4})) // false
}

// Runtime O(n), Space O(n)
func containsDuplicateWithSpace(nums []int) bool {
	memo := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		memo[nums[i]]++
		if memo[nums[i]] > 1 {
			return true
		}
	}

	return false
}

// Runtime O(n logn), Space O(1)
func containsDuplicateWithoutSpace(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
