package main

import "log"

// Given a non-empty array of non-negative integers nums, the degree of this
// array is defined as the maximum frequency of any one of its elements.
//
// Your task is to find the smallest possible length of a (contiguous)
// subarray of nums, that has the same degree as nums.
func main() {
	log.Println(findShortestSubArray([]int{1, 3, 2, 2, 3, 1}))    // 2 [2,2]
	log.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})) // 6 [2,2,3,1,4,2]
}

// Runtime - O(n) - one pass
// Space - O(M) - M - size of numbers set
func findShortestSubArray(nums []int) int {
	counter := make(map[int]int)
	left := make(map[int]int)
	min, d := 0, 0
	for i, n := range nums {
		counter[n]++
		if _, ok := left[n]; !ok {
			left[n] = i
		}

		if counter[n] > d {
			d = counter[n]
			min = i - left[n] + 1
		} else if counter[n] == d {
			if i-left[n]+1 < min {
				min = i - left[n] + 1
			}
		}
	}

	return min
}

// Runtime - O(n) - two pass
// Space - O(n)
//func findShortestSubArray(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	left := make(map[int]int)
//	right := make(map[int]int)
//	count := make(map[int]int)
//
//	for i, n := range nums {
//		if _, ok := count[n]; !ok {
//			count[n]++
//			left[n] = i
//			right[n] = i
//			continue
//		}
//		count[n]++
//		right[n] = i
//	}
//
//	minLen, degree := len(nums), 0
//	for n, d := range count {
//		l := right[n] - left[n] + 1
//		if d > degree {
//			degree = d
//			minLen = l
//		} else if d == degree && l < minLen {
//			minLen = l
//		}
//	}
//
//	return minLen
//}
