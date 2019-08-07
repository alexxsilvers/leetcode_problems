package main

import (
	"log"
	"math"
)

// Given an integer array, find three numbers whose product is maximum and output the maximum product.
func main() {
	log.Println(maximumProduct([]int{1, 2, 3}))            // 6
	log.Println(maximumProduct([]int{-1, -2, -3}))         // 6
	log.Println(maximumProduct([]int{1, 2, 3, 4}))         // 24
	log.Println(maximumProduct([]int{-4, -3, -2, -1, 60})) // 720
}

// Runtime - O(n)
// Space - O(1)
func maximumProduct(nums []int) int {
	max1, max2, max3, min1, min2 := math.MinInt32, math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n >= max1 {
			max3, max2, max1 = max2, max1, n
		} else if n >= max2 {
			max3, max2 = max2, n
		} else if n >= max3 {
			max3 = n
		}

		if n <= min1 {
			min2, min1 = min1, n
		} else if n <= min2 {
			min2 = n
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}

// Runtime - O(nlogn)
// Space - O(logn)
//func maximumProduct(nums []int) int {
//	sort.Ints(nums)
//	l := len(nums)
//	return max(nums[l-1]*nums[l-2]*nums[l-3],
//		nums[0]*nums[1]*nums[l-1])
//}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
