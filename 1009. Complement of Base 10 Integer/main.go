package main

import (
	"log"
	"math/bits"
)

// Every non-negative integer N has a binary representation.  For example,
// 5 can be represented as "101" in binary, 11 as "1011" in binary, and so on.
// Note that except for N = 0, there are no leading zeroes in any binary representation.
//
// The complement of a binary representation is the number in binary you get when
// changing every 1 to a 0 and 0 to a 1.  For example, the complement of "101" in binary is "010" in binary.
//
// For a given number N in base-10, return the complement of it's binary
// representation as a base-10 integer.
func main() {
	log.Println(bitwiseComplement(5)) // 2
	log.Println(bitwiseComplement(7)) // 2
	log.Println(bitwiseComplement(10)) // 2
}

func bitwiseComplement(N int) int {
	return bitwiseComplementDivision(N)
}

func bitwiseComplementDivision(N int) int {
	n := 0

	for N > 0 {
		if N%2 > 0 {
			n++
		}
		N = N / 2
	}

	return n
}

func bitwiseComplementBits(N int) int {
	if N == 0 {
		return 1
	}

	shift := bits.Len(uint(N))
	return N ^ (1<<shift - 1)
}
