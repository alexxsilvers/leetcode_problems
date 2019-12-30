package main

import "fmt"

// Balanced strings are those who have equal quantity of 'L' and 'R' characters.
//
// Given a balanced string s split it in the maximum amount of balanced strings.
//
// Return the maximum amount of splitted balanced strings.
func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4 [RL RRLL RL RL]
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // 3 [RL LLLRRR LR]
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // 1 [LLLLRRRR]
}

func balancedStringSplit(s string) int {
	lCnt, cnt, start := 0, 0, 0
	for start < len(s) {
		if s[start] == 'L' {
			lCnt++
		} else {
			lCnt--
		}

		if lCnt == 0 {
			cnt++
		}

		start++
	}

	return cnt
}
