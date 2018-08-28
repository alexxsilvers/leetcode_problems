package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	var count int
	for _, s := range S {
		for _, j := range J {
			if j == s {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}