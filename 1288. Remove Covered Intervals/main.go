package main

import (
	"log"
	"math"
	"sort"
)

// Given a list of intervals, remove all intervals that are covered by another interval in the list.
//
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
//
// After doing so, return the number of remaining intervals.
func main() {
	log.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))    // 2 [[1, 4],[2, 8]]
	log.Println(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))            // 1
	log.Println(removeCoveredIntervals([][]int{{0, 10}, {5, 12}}))          // 2
	log.Println(removeCoveredIntervals([][]int{{3, 10}, {4, 10}, {5, 11}})) // 2
	log.Println(removeCoveredIntervals([][]int{{1, 2}, {1, 4}, {3, 4}})) // 1
}

func removeCoveredIntervals(intervals [][]int) int {
	return removeCoveredIntervalsLinear(intervals)
}

func removeCoveredIntervalsLinear(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][0]
		}
	})

	ans, end := len(intervals), math.MinInt64
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] <= end {
			ans--
		} else {
			end = intervals[i][1]
		}
	}

	return ans
}

func removeCoveredIntervalsGreedy(intervals [][]int) int {
	ans := len(intervals)

	var covered func(intervals [][]int, ind int, int []int) bool
	covered = func(intervals [][]int, ind int, int []int) bool {
		for i, interval := range intervals {
			if i == ind {
				continue
			}

			if int[0] >= interval[0] && int[1] <= interval[1] {
				return true
			}
		}

		return false
	}

	for i, interval := range intervals {
		if covered(intervals, i, interval) {
			ans--
		}
	}

	return ans
}
