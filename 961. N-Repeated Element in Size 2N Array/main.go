package main

import "fmt"

// In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.
//
// Return the element repeated N times.
func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))             // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))       // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}

func repeatedNTimes(A []int) int {
	memory := make(map[int]int)
	for _, n := range A {
		memory[n]++
		if memory[n] > 1 {
			return n
		}
	}

	return 0
}
