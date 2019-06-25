package main

import "log"

// Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.
//
// Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:
//
// If S[i] == "I", then A[i] < A[i+1]
// If S[i] == "D", then A[i] > A[i+1]
func main() {
	log.Println(diStringMatch("IDID")) // [0,4,2,3,1]
	log.Println(diStringMatch("III"))  // [0,1,2,3]
	log.Println(diStringMatch("IIID")) // [0,1,2,4,3]
	log.Println(diStringMatch("DDDI")) // [4,3,2,0,1]
}

func diStringMatch(S string) []int {
	low, high := 0, len(S)
	res := make([]int, len(S)+1)
	for i := 0; i < len(S); i++ {
		if S[i] == 'D' {
			res[i] = high
			high--
		} else {
			res[i] = low
			low++
		}
	}
	res[len(S)] = low
	return res
}
