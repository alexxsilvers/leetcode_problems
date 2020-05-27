package main

import (
	"log"
)

// Given two binary strings, return their sum (also a binary string).
// The input strings are both non-empty and contains only characters 1 or 0.
func main() {
	log.Println(addBinary("11", "1"))      // "100"
	log.Println(addBinary("1010", "1011")) // "10101"
	log.Println(addBinary("1010", "1"))    // "1011"
	log.Println(addBinary("1111", "1"))    // "10000"
	log.Println(addBinary("0", "0"))       // "0"
}

func addBinary(a string, b string) string {
	res := make([]byte, max(len(a), len(b))+1)
	i, j, k := len(res)-1, len(a)-1, len(b)-1
	var carry byte
	for j >= 0 || k >= 0 {
		sum := carry
		if j >= 0 {
			sum += a[j] - '0'
			j--
		}
		if k >= 0 {
			sum += b[k] - '0'
			k--
		}

		carry = sum / 2
		res[i] = sum%2 + '0'
		i--
	}

	if carry > 0 {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
