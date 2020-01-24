package main

import "log"

// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it is able to trap after raining.
func main() {
	log.Println(trapBruteForce([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	log.Println(trapDP([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))         // 6
}
