package main

import "log"

// Given two arrays, write a function to compute their intersection.
func main() {
	log.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))       // [2]
	log.Println(intersection([]int{4, 5, 9}, []int{9, 4, 9, 8, 4})) // [4,9]
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	memo := make(map[int]bool)
	for _, n := range nums1 {
		memo[n] = true
	}

	var ret []int
	for _, n := range nums2 {
		if memo[n] {
			ret = append(ret, n)
			delete(memo, n)
		}
	}

	return ret
}
