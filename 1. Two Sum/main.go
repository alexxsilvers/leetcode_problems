package main

import "log"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
func main() {
	log.Println(twoSumMap([]int{2, 7, 11, 15}, 18)) // [1, 2]
	log.Println(twoSumMap([]int{3, 2, 4}, 6))       // [1, 2]
}

// Runtime O(n), Space O(n)
func twoSumMap(nums []int, target int) []int {
	memo := make(map[int]int)
	for i, n := range nums {
		ind, ok := memo[target - n]
		if ok {
			return []int{i, ind}
		} else {
			memo[n] = i
		}
	}

	return []int{}
}
