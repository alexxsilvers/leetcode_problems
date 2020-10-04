package main

import (
	"log"
	"sort"
)

// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
//
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
//
// 0 <= i, j < nums.length
// i != j
// a <= b
// b - a == k
func main() {
	log.Println(findPairs([]int{3, 1, 4, 1, 5}, 2)) // 2 - [1, 3], [3, 5]
	log.Println(findPairs([]int{1, 2, 3, 4, 5}, 1)) // 4
	log.Println(findPairs([]int{1, 3, 1, 5, 4}, 0)) // 1 - [1, 1]
}

func findPairs(nums []int, k int) int {
	//return findPairsWithSpace(nums, k)
	return findPairsLinear(nums, k)
}

func findPairsLinear(nums []int, k int) int {
	sort.Ints(nums)
	ans, i, j := 0, 0, 1

	for i < len(nums) && j < len(nums) {
		if nums[j]-nums[i] == k && i != j {
			ans++

			for i+1 < len(nums) && nums[i+1] == nums[i] {
				i++
			}

			for j+1 < len(nums) && nums[j+1] == nums[j] {
				j++
			}
		}
		if nums[j]-nums[i] > k {
			i++
		} else {
			j++
		}
	}

	return ans
}

func findPairsWithSpace(nums []int, k int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	ans := 0
	for num := range m {
		if k > 0 {
			if _, ok := m[num+k]; ok {
				ans++
			}
			continue
		}

		if k == 0 {
			n, ok := m[num]
			if ok && n > 1 {
				ans++
			}
			continue
		}
	}

	return ans
}
