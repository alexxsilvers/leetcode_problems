package _00__Longest_increasing_subsequence

import (
	"log"
)

// Given an unsorted array of integers, find the length of longest increasing subsequence.
func main() {
	log.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))    // [2, 3, 7, 101] = 4
	log.Println(lengthOfLIS([]int{103, 4, 40, 4, 8, 22, 45, 120})) // [4, 8, 22, 45, 120] = 5
}

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	d, size := make([]int, l), 0
	for n := 0; n < l; n++ {
		i, j := 0, size
		for i != j {
			mid := (i + j) / 2
			if d[mid] < nums[n] {
				i = mid + 1
			} else {
				j = mid
			}
		}
		d[i] = nums[n]
		if i == size {
			size++
		}
	}
	return size
}
