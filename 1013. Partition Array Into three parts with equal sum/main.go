package main

import "log"

// Given an array A of integers, return true if and only if we can partition the array
// into three non-empty parts with equal sums.
//
// Formally, we can partition the array if we can find indexes i+1 < j with
// (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])
func main() {
	log.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})) // true
	// 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

	log.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1})) // false

	log.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4})) // true
	// 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

	log.Println(canThreePartsEqualSum([]int{12, -4, 16, -5, 9, -3, 3, 8, 0})) // true
	// 12 = -4 + 16 = -5 + 9 - 3 + 3 + 8 + 0
}

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, num := range A {
		sum += num
	}
	if sum%3 != 0 {
		return false
	}
	part := sum / 3
	partSum, partCnt := 0, 0

	for _, num := range A {
		partSum += num
		if partSum == part {
			partCnt++
			partSum = 0
		}
	}

	return partCnt == 3 && partSum == 0
}
