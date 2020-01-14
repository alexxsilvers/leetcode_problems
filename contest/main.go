package main

import "log"

func main() {
	log.Println(getNoZeroIntegers(2))     // [1,1]
	log.Println(getNoZeroIntegers(11))    // [2,9]
	log.Println(getNoZeroIntegers(10000)) // [1,9999]
	log.Println(getNoZeroIntegers(69))    // [1,68]
	log.Println(getNoZeroIntegers(1010))  // [11, 999]
}

func getNoZeroIntegers(n int) []int {
	first, second := 1, n-1

	for containsZero(first) || containsZero(second) {
		first++
		second--
	}

	return []int{first, second}
}

func containsZero(n int) bool {
	if n == 0 {
		return true
	}

	for n > 0 {
		if n%10 == 0 {
			return true
		}

		n /= 10
	}

	return false
}
