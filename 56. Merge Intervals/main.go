package main

import (
	"log"
	"sort"
)

// Given a collection of intervals, merge all overlapping intervals.
func main() {
	log.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	log.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
	log.Println(merge([][]int{{1, 4}, {-2, 10}, {4, 5}}))          // [[-2,10]]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	ind := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[ind][1] {
			merged = append(merged, intervals[i])
			ind++
		} else {
			merged[ind][1] = max(intervals[i][1], merged[ind][1])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
