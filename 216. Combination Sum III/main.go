package main

import "log"

// Find all valid combinations of k numbers that sum up to n such that the following
// conditions are true:
//
// Only numbers 1 through 9 are used.
// Each number is used at most once.
// Return a list of all possible valid combinations. The list must not contain the same
// combination twice, and the combinations may be returned in any order.
func main() {
	log.Println(combinationSum3(3, 7))  // [[1, 2, 4]]
	log.Println(combinationSum3(9, 45)) // [[1, 2, 3, 4, 5, 6, 7, 8, 9]]
	log.Println(combinationSum3(4, 12)) // [[1, 2, 3, 6], [1, 2, 4, 5]]
	log.Println(combinationSum3(8, 36)) // [[1,2,3,4,5,6,7,8]]
}

func combinationSum3(k int, n int) [][]int {
	var backtracking func(nums []int, start int, target int, combinations *[][]int)
	backtracking = func(nums []int, start int, target int, combinations *[][]int) {
		if len(nums) == k {
			if target == 0 {
				*combinations = append(*combinations, append([]int{}, nums...))
			}
			return
		}

		for i := start; i < 10; i++ {
			if i > target {
				break
			}

			backtracking(append(nums, i), i+1, target-i, combinations)
		}
	}

	var combinations [][]int
	backtracking([]int{}, 1, n, &combinations)
	return combinations
}
