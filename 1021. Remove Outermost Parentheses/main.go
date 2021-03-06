package main

import (
	"log"
	"strings"
)

// A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid
// parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))"
// are all valid parentheses strings.
//
// A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it
// into S = A+B, with A and B nonempty valid parentheses strings.
//
// Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k,
// where P_i are primitive valid parentheses strings.
//
// Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.
func main() {
	log.Println(removeOuterParentheses("(()())(())"))         // ()()()
	log.Println(removeOuterParentheses("(()())(())(()(()))")) // ()()()()(())
	log.Println(removeOuterParentheses("()()"))               // ""
}

func removeOuterParentheses(S string) string {
	opened := 0
	sb := strings.Builder{}
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if opened > 0 {
				sb.WriteByte('(')
			}
			opened++
		} else {
			if opened > 1 {
				sb.WriteByte(')')
			}
			opened--
		}
	}

	return sb.String()
}
