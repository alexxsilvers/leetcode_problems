package main

import "log"

// Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
// elements appear twice and others appear once.
//
// Find all the elements of [1, n] inclusive that do not appear in this array.
//
// Could you do it without extra space and in O(n) runtime? You may assume
// the returned list does not count as extra space.
func main() {
	log.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [5, 6]
}

// Runtime - O(n)
// Space - O(n) - in bad case.
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)

	for i := 0; i < size; i++ {
		num := abs(nums[i]) - 1
		if nums[num] > 0 {
			nums[num] = nums[num] * -1
		}
	}

	ret := make([]int, 0)
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
