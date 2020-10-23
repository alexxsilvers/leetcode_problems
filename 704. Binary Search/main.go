package main

import "log"

// Given a sorted (in ascending order) integer array nums of n elements and a
// target value, write a function to search target in nums. If target exists,
// then return its index, otherwise return -1.
func main() {
	log.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	log.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
}

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		mid := l + (h-l)/2
		num := nums[mid]

		if num < target {
			l = mid + 1
		} else if num > target {
			h = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
