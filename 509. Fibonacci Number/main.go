package main

import "log"

// The Fibonacci numbers, commonly denoted F(n) form a sequence,
// called the Fibonacci sequence, such that each number is the sum of
// the two preceding ones, starting from 0 and 1. That is,
//
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), for N > 1.
// Given N, calculate F(N).
func main() {
	log.Println(fib(7))  // 13
	log.Println(fib(10)) // 55
	log.Println(fib(0))  // 0
	log.Println(fib(1))  // 1
	log.Println(fib(2))  // 1
}

func fib(N int) int {
	if N < 2 {
		return N
	}

	table := make([]int, 2)
	table[0] = 0
	table[1] = 1
	for i := 2; i <= N; i++ {
		table[i%2] = table[0] + table[1]
	}
	return table[N%2]
}
