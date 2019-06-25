package main

import (
	"log"
)

// A self-dividing number is a number that is divisible by every digit it contains.
//
// For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
//
// Also, a self-dividing number is not allowed to contain the digit zero.
//
// Given a lower and upper number bound, output a list of every possible self
// dividing number, including the bounds if possible.
func main() {
	log.Println(selfDividingNumbers(1, 22)) // [1,2,3,4,5,6,7,8,9,11,12,15,22]
}


func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for left <= right {
		if isSDN(left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func isSDN(n int) bool {
	digit := 0
	tmp := n
	for tmp != 0 {
		digit = tmp % 10
		tmp = tmp / 10
		if digit == 0 {
			return false
		}
		if n % digit != 0 {
			return false
		}
	}
	return true
}