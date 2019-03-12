package main

import (
	"log"
)

func main() {
	log.Println(FindMissing([]int{5, 6, 8, 12}, []int{6,8, 12}))
}

func FindMissingXOR(full []int, partial []int) int {
	var xor_sum byte
	for _, num := range full {
		xor_sum ^= byte(num)
	}
	for _, num := range partial {
		xor_sum ^= byte(num)
	}

	return int(xor_sum)
}

func FindMissing(full []int, partial []int) int {
	var num_int int
	for _, num := range full {
		num_int += num
	}
	for _, num := range partial {
		num_int -= num
	}

	return num_int
}