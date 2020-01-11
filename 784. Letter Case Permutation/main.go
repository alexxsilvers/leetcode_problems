package main

import "log"

// Given a string S, we can transform every letter individually to be lowercase
// or uppercase to create another string. Return a list of all possible strings we could create.
func main() {
	log.Println(letterCasePermutation("a1b2"))  // ["a1b2", "A1b2", "A1B2", "a1B2"]
	log.Println(letterCasePermutation("3z4"))   // ["3z4", "3Z4"]
	log.Println(letterCasePermutation("12345")) // ["12345"]
	log.Println(letterCasePermutation("C"))     // ["C", "c"]
	log.Println(letterCasePermutation("CC"))    // ["CC", "cC", "Cc", "cc"]
	log.Println(letterCasePermutation("0"))     // ["0"]
}

func letterCasePermutation(S string) []string {
	ans := make([]string, 0)
	backtracking(&ans, 0, []byte(S))
	return ans
}

func backtracking(ans *[]string, pos int, chars []byte) {
	if pos == len(chars) {
		*ans = append(*ans, string(chars))
		return
	}

	if isNum(chars[pos]) {
		backtracking(ans, pos+1, chars)
	} else {
		chars[pos] = toUpper(chars[pos])
		backtracking(ans, pos+1, chars)

		chars[pos] = toLower(chars[pos])
		backtracking(ans, pos+1, chars)
	}
}

func toUpper(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char
	}

	return char - 32
}

func toLower(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char
	}

	return char + 32
}

func isNum(char byte) bool {
	return char >= '0' && char <= '9'
}
