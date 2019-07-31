package main

import (
	"log"
	"sort"
)

// Given two arrays arr1 and arr2, the elements of arr2 are distinct,
// and all elements in arr2 are also in arr1.
//
// Sort the elements of arr1 such that the relative ordering of items
// in arr1 are the same as in arr2.  Elements that don't appear in arr2
// should be placed at the end of arr1 in ascending order.
func main() {
	log.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	// [2,2,2,1,4,3,3,9,6,7,19]

	log.Println(relativeSortArray([]int{2, 3, 1, 7, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	// [2,2,2,1,4,3,3,9,6,7,19]
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arrMap1 := make(map[int]int, len(arr1))
	for _, n := range arr1 {
		arrMap1[n]++
	}

	ret := make([]int, 0, len(arr1))
	for _, n := range arr2 {
		for i := 0; i < arrMap1[n]; i++ {
			ret = append(ret, n)
		}
		delete(arrMap1, n)
	}

	tail := make([]int, 0, len(arrMap1))
	for n := range arrMap1 {
		tail = append(tail, n)
	}

	sort.Ints(tail)

	for _, n := range tail {
		for i := 0; i < arrMap1[n]; i++ {
			ret = append(ret, n)
		}
	}

	return ret
}
