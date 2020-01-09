package main

import "log"

// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate
// (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container, such that the container contains
// the most water.
//
// Note: You may not slant the container and n is at least 2.
func main() {
	log.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}

// Runtime - O(n), Space - O(1)
func maxArea(height []int) int {
	max, s, e := 0, 0, len(height)-1

	for s < e {
		h := min(height[s], height[e]) * (e - s)
		if h > max {
			max = h
		}

		if height[s] < height[e] {
			s++
		} else {
			e--
		}
	}

	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
