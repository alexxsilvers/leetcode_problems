package main

import "log"

func main() {
	log.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	log.Println(pivotIndex([]int{1, 2, 3}))
}

func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	leftSum := 0
	for i := range nums {
		if sum == leftSum * 2 + nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}