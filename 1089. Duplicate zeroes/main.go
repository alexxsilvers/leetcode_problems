package main

import "log"

func main() {
	ex1 := []int{1, 0, 2, 3, 0, 4, 5, 0} // 10023004
	duplicateZeros(ex1)
	log.Println(ex1)

	ex2 := []int{1, 2, 3}
	duplicateZeros(ex2) // 123
	log.Println(ex2)
}

func duplicateZeros(arr []int) {
	cp := make([]int, len(arr))
	copy(cp, arr)
	p := 0
	for i := 0; i < len(cp) && p < len(cp); i++ {
		arr[p] = cp[i]
		p++
		if cp[i] == 0 && p < len(cp) {
			arr[p] = 0
			p++
		}
	}
}
