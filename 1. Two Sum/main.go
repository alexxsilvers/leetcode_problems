package main

import (
	"log"
)

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
func main() {
	log.Println(twoSum([]int{2, 7, 11, 15}, 18)) // [1, 2]
	log.Println(twoSum([]int{3, 2, 4}, 6))       // [1, 2]
}

func twoSum(nums []int, target int) []int {
	start := 0
	memo := make(map[int]int, len(nums))

	for start < len(nums) {
		ind, exist := memo[target-nums[start]]
		if exist {
			return []int{start, ind}
		} else {
			memo[nums[start]] = start
		}

		start++
	}

	return []int{-1, -1}
}
