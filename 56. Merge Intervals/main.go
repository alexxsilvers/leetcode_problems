package main

import (
	"log"
	"sort"
)

// Given a collection of intervals, merge all overlapping intervals.
func main() {
	log.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	log.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
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
