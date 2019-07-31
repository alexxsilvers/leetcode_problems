package main

import "log"

// Given an array A of non-negative integers, return the maximum sum of elements in two
// non-overlapping (contiguous) subarrays, which have lengths L and M.
// (For clarification, the L-length subarray could occur before or after the M-length subarray.)
//
// Formally, return the largest V for which V = (A[i] + A[i+1] + ... + A[i+L-1]) + (A[j] + A[j+1] + ... + A[j+M-1]) and either:
//
// 0 <= i < i + L - 1 < j < j + M - 1 < A.length, or
// 0 <= j < j + M - 1 < i < i + L - 1 < A.length.

// EXPLANATION
// I finded only brute force solution
// We need find all subarray sum with len M and L and store it
// to the two maps map[last index of subarray]value
// and after intersect it
func main() {
	log.Println(maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2)) // 20
	// One choice of sub arrays is [9] with length 1, and [6,5] with length 2.

	log.Println(maxSumTwoNoOverlap([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2)) // 29
	// One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
}

// Brute force
// Runtime O(n)
// Space O(n)
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	l := len(A)
	if M+L == l {
		return sumArray(A)
	}

	lDict := combine(A, L)
	mDict := combine(A, M)
	max := 0

	for lEndIndex, lSum := range lDict {
		for mEndIndex, mSum := range mDict {
			t := lSum + mSum
			if !intersect(mEndIndex-M, mEndIndex-1, lEndIndex-L, lEndIndex-1) && t > max {
				max = t
			}
		}
	}
	return max
}

func intersect(mSTartIndex, lStartIndex, mEndIndex, lEndIndex int) bool {
	if lStartIndex > mSTartIndex && lEndIndex < mEndIndex {
		return true
	}
	return false
}

func combine(A []int, subLen int) map[int]int {
	l := len(A)
	ret := make(map[int]int)
	for i := 0; i < l-subLen; i++ {
		ret[i+subLen] = sumArray(A[i : i+subLen])
	}
	return ret
}

func sumArray(A []int) int {
	s := 0
	for _, n := range A {
		s += n
	}
	return s
}
