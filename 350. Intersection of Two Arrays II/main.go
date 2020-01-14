package main

import "log"

// Given two arrays, write a function to compute their intersection.
func main() {
	log.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       // [2, 2]
	log.Println(intersect([]int{4, 5, 9}, []int{9, 4, 9, 8, 4})) // [4, 9]
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	memo := make(map[int]int)
	for _, n := range nums1 {
		memo[n]++
	}

	var ret []int
	for _, n := range nums2 {
		cnt := memo[n]
		if cnt > 0 {
			ret = append(ret, n)
			memo[n]--
		}
	}

	return ret
}
