package main

import (
	"log"
)

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
func main() {
	log.Println(generateParenthesis(3))
	/*
		[
		  "((()))",
		  "(()())",
		  "(())()",
		  "()(())",
		  "()()()"
		]
	*/

	log.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backtracking(&ans, []byte{}, 0, 0, n)
	return ans
}

func backtracking(ans *[]string, cur []byte, open, closed, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, string(cur))
		return
	}

	if open < max {
		backtracking(ans, append(cur, '('), open+1, closed, max)
	}

	if closed < open {
		backtracking(ans, append(cur, ')'), open, closed+1, max)
	}
}
