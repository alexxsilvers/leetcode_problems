package main

import (
	"log"
	"math"
)

// Given an array of n positive integers and a positive integer s, find the minimal
// length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
func main() {
	log.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))    // 2 [4,3]
	log.Println(minSubArrayLen(7, []int{2, 3, 7}))             // 1 [7]
	log.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3, 7})) // 1 [7]
	log.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 7, 3}))    // 1 [7]
	log.Println(minSubArrayLen(6, []int{10, 2, 3}))            // 1 [10]
	log.Println(minSubArrayLen(10, []int{11}))                 // 1 [11]
	log.Println(minSubArrayLen(3, []int{1, 1}))                // 0
	log.Println(minSubArrayLen(3, []int{1}))                   // 0
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end, sum, max := 0, 0, nums[0], math.MaxInt64

	for end < len(nums) {
		if sum >= s {
			l := end - start + 1
			if l < max {
				max = l
			}

			sum -= nums[start]
			start++
			continue
		}

		end++
		if end < len(nums) {
			sum += nums[end]
		}
	}

	if max == math.MaxInt64 {
		return 0
	}

	return max
}
