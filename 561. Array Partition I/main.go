package main

import (
	"log"
	"sort"
)

// Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ...,
// (an, bn) such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.
func main() {
	log.Println(arrayPairSum([]int{4, 1, 3, 2}))       // 4 => min(2, 1) + min(3, 4) = 1 + 3
	log.Println(arrayPairSum([]int{6, 1, 2, 2, 5, 6})) // 9 => min(1, 2) + min(2, 5) + min(6, 6) = 1 + 2 + 6
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0

	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}

	return res
}
