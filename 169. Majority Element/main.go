package main

import "log"

// Given an array of size n, find the majority element. The majority
// element is the element that appears more than [n/2] times.
//
// You may assume that the array is non-empty and the majority element
// always exist in the array.
func main() {
	log.Println(majorityElement([]int{3, 2, 3}))       // 3
	log.Println(majorityElement([]int{1}))             // 1
	log.Println(majorityElement([]int{2, 2, 1, 1, 2})) // 2
}

// Runtime - O(n)
// Space - O(1)
func majorityElement(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return nums[0]
	}

	m, counter := nums[0], 1
	for i := 1; i < size; i++ {
		if counter == 0 {
			counter = 1
			m = nums[i]
		} else if m == nums[i] {
			counter++
		} else {
			counter--
		}
	}

	return m
}

// Runtime - O(n)
// Space - O(n)
//func majorityElement(nums []int) int {
//	size := len(nums)
//	if size <= 2 {
//		return nums[0]
//	}
//
//	half := size / 2
//	counterMap := make(map[int]int)
//	for _, n := range nums {
//		counterMap[n]++
//		if counterMap[n] > half {
//			return n
//		}
//	}
//	return 0
//}
