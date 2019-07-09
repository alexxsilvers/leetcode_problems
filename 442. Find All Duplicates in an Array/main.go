package main

import "log"

// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
// Find all the elements that appear twice in this array.
//
// Could you do it without extra space and in O(n) runtime?

// INTUITION
// Most important that this is an array of integers 1 ≤ a[i] ≤ n (n = size of array)
// It means that we can make negative each of a[a[i]-1] when traversal
// If a[a[i]-1] already negative - this is duplicate
func main() {
	log.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [2,3]
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// Runtime O(n)
// Space O(1)
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, n := range nums {
		tmpI := abs(n) - 1
		if nums[tmpI] > 0 {
			nums[tmpI] *= -1
		} else {
			res = append(res, abs(n))
		}
	}
	return res
}

// Runtime O(n)
// Space O(n)
//func findDuplicates(nums []int) []int {
//	counter := make(map[int]int)
//	for _, n := range nums {
//		counter[n]++
//	}
//	ret := make([]int, 0)
//	for n, cnt := range counter {
//		if cnt > 1 {
//			ret = append(ret, n)
//		}
//	}
//	return ret
//}
