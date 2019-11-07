package main

import "log"

// Given an array of integers arr, write a function that returns true if and only if
// the number of occurrences of each value in the array is unique.
func main() {
	log.Println(uniqueOccurrences([]int{1, 1, 1, 2, 3, 2}))                 // true
	log.Println(uniqueOccurrences([]int{1, 2}))                             // false
	log.Println(uniqueOccurrences([]int{10}))                               // true
	log.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, n := range arr {
		occurrences[n]++
	}

	existOccurrences := make(map[int]bool)
	for _, occurrence := range occurrences {
		if !existOccurrences[occurrence] {
			existOccurrences[occurrence] = true
		} else {
			return false
		}
	}

	return true
}
