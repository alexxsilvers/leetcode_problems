package main

import "log"

// Given an array of integers, find if the array contains any duplicates.
//
// Your function should return true if any value appears at least twice in
// the array, and it should return false if every element is distinct.
func main() {
	log.Println(containsDuplicate([]int{3, 1, 2, 3})) // true
	log.Println(containsDuplicate([]int{3, 1, 2, 4})) // false
}

// Runtime - O(n)
// Space - O(n)
func containsDuplicate(nums []int) bool {
	existMap := make(map[int]struct{})
	for _, n := range nums {
		_, ok := existMap[n]
		if ok {
			return true
		}
		existMap[n] = struct{}{}
	}
	return false
}
