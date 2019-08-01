package main

import "log"

// Given an array of integers that is already sorted in ascending order, find two
// numbers such that they add up to a specific target number.
//
// he function twoSum should return indices of the two numbers such that they add up
// to the target, where index1 must be less than index2.
//
// Note:
// Your returned answers (both index1 and index2) are not zero-based.
// You may assume that each input would have exactly one solution and you may not use the same element twice.
func main() {
	log.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1,2]
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		v := numbers[left] + numbers[right]
		if v == target {
			return []int{left + 1, right + 1}
		}

		if v > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
