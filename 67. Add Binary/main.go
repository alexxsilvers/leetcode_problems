package main

import (
	"log"
)

// Given two binary strings, return their sum (also a binary string).
//
// The input strings are both non-empty and contains only characters 1 or 0.
func main() {
	log.Println(addBinary("11", "1"))
	log.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	res := make([]byte, 0)
	i, k := len(a)-1, len(b)-1
	var carry uint8
	for i >= 0 || k >= 0 {
		if i >= 0 {
			carry += a[i] - '0'
			i--
		}

		if k >= 0 {
			carry += b[k] - '0'
			k--
		}

		res = append(res, carry%2+'0')
		carry /= 2
	}

	if carry > 0 {
		res = append(res, 1+'0')
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}
