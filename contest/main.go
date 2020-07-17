package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	log.Println(canMakeArithmeticProgression([]int{3, 5, 1, 2}))
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
