package main

import "log"

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
func main() {
	log.Println(singleNumber([]int{2, 2, 1}))       // 1
	log.Println(singleNumber([]int{4, 1, 2, 1, 2})) // 4
}

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}

	return ret
}
