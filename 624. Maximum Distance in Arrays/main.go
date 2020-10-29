package main

import (
	"log"
)

// Given m arrays, and each array is sorted in ascending order. Now you can pick up two
// integers from two different arrays (each array picks one) and calculate the distance.
// We define the distance between two integers a and b to be their absolute difference
// |a-b|. Your task is to find the maximum distance.
func main() {
	log.Println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}})) // 4
	log.Println(maxDistance([][]int{{1, 4}, {0, 5}}))               // 4
	log.Println(maxDistance([][]int{{1, 5}, {3, 4}}))               // 3
}

func maxDistance(arrays [][]int) int {
	minVal, maxVal, res := arrays[0][0], arrays[0][len(arrays[0])-1], 0
	for i := 1; i < len(arrays); i++ {
		res = max(res, max(abs(arrays[i][len(arrays[i])-1]-minVal), abs(maxVal-arrays[i][0])))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][len(arrays[i])-1])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
