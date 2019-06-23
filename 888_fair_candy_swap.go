package main

import "log"

// Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy
// that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.
//
// Since they are friends, they would like to exchange one candy bar each so that after the exchange,
// they both have the same total amount of candy. (The total amount of candy a person has is
// the sum of the sizes of candy bars they have.)
//
// Return an integer array ans where ans[0] is the size of the candy bar that Alice
// must exchange, and ans[1] is the size of the candy bar that Bob must exchange.
//
// If there are multiple answers, you may return any one of them. It is guaranteed an answer exists.
func main() {
	log.Println(fairCandySwap([]int{1,1}, []int{2,2})) // [1, 2]
	log.Println(fairCandySwap([]int{1}, []int{3,2})) // [1, 3]
}

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	indexB := make(map[int]bool)
	for _, num := range A {
		sumA += num
	}
	for _, num := range B {
		sumB += num
		indexB[num] = true
	}

	halfDiff := (sumB - sumA) / 2
	for _, numA := range A {
		numB := numA + halfDiff
		if indexB[numB] {
			return []int{numA, numB}
		}
	}

	return nil
}