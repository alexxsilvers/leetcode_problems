package main

import "log"

// On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
//
// Once you pay the cost, you can either climb one or two steps. You need to find
// minimum cost to reach the top of the floor, and you can either start from the
// step with index 0, or the step with index 1.
func main() {
	log.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // 15
	log.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
}

func minCostClimbingStairs(cost []int) int {
	beforeOne := cost[1]
	beforeTwo := cost[0]
	for i := 2; i < len(cost); i++ {
		curr := min(beforeOne, beforeTwo) + cost[i]
		beforeTwo = beforeOne
		beforeOne = curr
	}

	return min(beforeOne, beforeTwo)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
