package main

import "log"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sums to target.
// The same repeated number may be chosen from candidates unlimited number of times.
// Note:
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
func main() {
	log.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[7], [2, 2, 3]]
	log.Println(combinationSum([]int{2, 3, 5}, 8)) // [[2,2,2,2], [3, 5], [2, 3, 3]]
}

func combinationSum(candidates []int, target int) [][]int {

}
