package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(countArrangement(1)) // 1
	log.Println(countArrangement(2)) // 2
	log.Println(countArrangement(3)) // 3 => [1, 2, 3], [3, 2, 1], [2, 1, 3]
}

func countArrangement(n int) int {
	ans := 1
	result := make([]int, n+1)
	for i := 1; i < n; i++ {
		result[i] = i
	}

	for i := 1; i < n; i++ {
		for k := i; k < n; k++ {

		}
	}

}

func isBeautiful(arr []int) bool {

}
