package main

import (
	"log"
	"sort"
)

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers
// sum to target. You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
func main() {
	log.Println(combinationSum([]int{2, 3, 6, 7}, 7))  // [[2,2,3], [7]]
	log.Println(combinationSum([]int{2, 3, 5}, 8))     // [[2,2,2,2], [2,3,3], [3,5]]
	log.Println(combinationSum([]int{2}, 1))           // []
	log.Println(combinationSum([]int{1}, 1))           // [[1]]
	log.Println(combinationSum([]int{1}, 2))           // [[1,1]]
	log.Println(combinationSum([]int{8, 7, 4, 3}, 11)) // [[3,4,4],[3,8],[4,7]]
}

func combinationSum(candidates []int, target int) [][]int {
	var backtracking func(nums []int, start int, target int, combinations *[][]int)
	backtracking = func(nums []int, start int, target int, combinations *[][]int) {
		if target == 0 {
			*combinations = append(*combinations, append([]int{}, nums...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}

			backtracking(append(nums, candidates[i]), i, target-candidates[i], combinations)
		}
	}

	sort.Ints(candidates)
	var combinations [][]int
	backtracking([]int{}, 0, target, &combinations)

	return combinations
}
