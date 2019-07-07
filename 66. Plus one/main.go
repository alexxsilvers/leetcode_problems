package main

import "log"

func main() {
	log.Println(plusOne([]int{1, 2, 3}))    // [1, 2, 4]
	log.Println(plusOne([]int{4, 3, 2, 1})) // [4, 3, 2, 2]
	log.Println(plusOne([]int{4, 5, 9}))    // [4, 6, 0]
	log.Println(plusOne([]int{9, 9, 9}))    // [1, 0, 0, 0]
}

func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += plus
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			plus = 0
			break
		}
	}

	if plus == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
