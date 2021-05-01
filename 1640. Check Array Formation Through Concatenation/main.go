package main

import "log"

func main() {
	log.Println(canFormArray([]int{15, 88}, [][]int{[]int{15}, []int{88}}))         // true
	log.Println(canFormArray([]int{88, 15}, [][]int{[]int{15}, []int{88}}))         // true
	log.Println(canFormArray([]int{49, 18, 16}, [][]int{[]int{16, 18, 49}}))        // false
	log.Println(canFormArray([]int{49, 18, 16}, [][]int{[]int{18, 16}, []int{49}})) // true
	log.Println(canFormArray([]int{1, 2, 3, 4}, [][]int{[]int{3, 4, 5, 6}}))        // false
}

func canFormArray(arr []int, pieces [][]int) bool {
	indMap := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indMap[arr[i]] = i
	}

	for i := 0; i < len(pieces); i++ {
		for k := 0; k < len(pieces[i]); k++ {
			_, exist := indMap[pieces[i][k]]
			if !exist {
				return false
			}

			if k > 0 {
				if indMap[pieces[i][k-1]]+1 != indMap[pieces[i][k]] {
					return false
				}
			}
		}
	}

	return true
}
