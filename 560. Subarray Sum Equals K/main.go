package main

import "log"

// Given an array of integers and an integer k, you need to find the total number of
// continuous subarrays whose sum equals to k.
func main() {
	log.Println(subarraySum([]int{1, 1, 1}, 2))       // 2
	log.Println(subarraySum([]int{1, 1, 1, 1, 1}, 2)) // 4
}

func subarraySum(nums []int, k int) int {
	return subarraySumMap(nums, k)
}

// prefix-sum technique. Runtime O(n), space - O(n)
func subarraySumMap(nums []int, k int) int {
	memo := make(map[int]int)
	memo[0] = 1
	ans, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		key := sum - k

		cnt, ok := memo[key]
		if ok {
			ans += cnt
		}
		memo[sum] += 1
	}

	return ans
}

// Runtime O(n2), space O(1)
func subarraySumBruteForce(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}

	return ans
}
